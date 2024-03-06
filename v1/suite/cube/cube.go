package cube

import (
	"context"
	"fmt"
	"github.com/zj-sh/meta/v1/base"
	"github.com/zj-sh/meta/v1/suite"
	"github.com/zj-sh/meta/v1/suite/kernel"
	"github.com/zohu/zdb"
	"github.com/zohu/zlog"
	"mored/zrepo"
	"os"
	"path"
	"regexp"
	"strings"
)

/**
 * @description: must be use mored 's database
 */

var fn = func() {}

// LoadSuiteClient
// @Description: 运行套件 [已注入配置]
// @param ctx
// @param suiteCode
// @param version
// @param cid
// @return *Entry
// @return string
// @return error
func LoadSuiteClient(ctx context.Context, suiteCode, version, cid string) (suite.SuiteClient, func(), error) {
	// 确定套件版本
	if version == "" || version == "latest" {
		version = DefaultLatestVersion(suiteCode)
	} else {
		reg, _ := regexp.Compile("^[v|V]?\\d{1,4}(\\.\\d{1,4}){0,4}$")
		if !reg.MatchString(version) {
			return nil, fn, fmt.Errorf("版本号格式错误, 请满足：[v]x(.x)(.x)(.x), x区间0-9999")
		}
		version = strings.ToLower(version)
	}
	// 确定套件是否安装启用
	if err := zdb.DB(ctx).Where("e=? and suite_code=? and suite_version=? and is_enable=true", cid, suiteCode, version).First(&zrepo.SuiteCorp{}).Error; err != nil {
		return nil, fn, fmt.Errorf("套件未安装或未启用")
	}
	// 查询套件详情
	var st zrepo.Suite
	if err := zdb.DB(ctx).Where("suite_code=? and suite_version=?", suiteCode, version).First(&st).Error; err != nil {
		return nil, fn, fmt.Errorf("套件版本不存在")
	}
	// 确定套件文件是否存在
	filename, err := SuiteFilePathFromDatabase(ctx, GetRunPath(), suiteCode, version, "")
	if err != nil {
		return nil, fn, err
	}
	// 确定套件是否已经运行
	e, err := LoadSuiteClientWithoutConf(suiteCode, version, filename, st.SuiteCmd)
	if err != nil {
		return nil, fn, err
	}
	conf := GetConfigStr(ctx, suiteCode, version, cid)
	if conf != "" {
		if _, err = e.Client().InjectConfig(ctx, &base.JsonString{Value: conf}); err != nil {
			return nil, fn, err
		}
		e.Release()
	}
	return e.Client(), e.Release, nil
}

// GetConfigStr
// @Description: 获取套件配置
// @param ctx
// @param suiteCode
// @param cid
// @return string
// @return error
func GetConfigStr(ctx context.Context, suiteCode, version, cid string) string {
	var st zrepo.SuiteCorp
	db := zdb.DB(ctx).Table("suite_corp sc").
		Select("sc.*").
		Where("sc.e=? and sc.suite_code=?", cid, suiteCode)
	if version == "" || version == "latest" {
		db = db.Joins("left join suite s on s.suite_code=sc.suite_code and s.suite_version=sc.suite_version").Order("s.suite_version_idx desc")
	} else {
		db = db.Where("sc.suite_version=?", version)
	}
	if err := db.Limit(1).Scan(&st).Error; err != nil {
		return ""
	}
	if st.SuiteConfig == "" {
		return ""
	}
	return st.SuiteConfig
}

// DefaultLatestVersion
// @Description: 获取套件最新版本
// @param suiteCode
// @return string
func DefaultLatestVersion(suiteCode string) string {
	var st zrepo.Suite
	zdb.DB(context.TODO()).Where("suite_code=?", suiteCode).Order("suite_version_idx desc").First(&st)
	return st.SuiteVersion
}

func SuiteFilePathFromDatabase(ctx context.Context, dir, suiteCode, version, suffix string) (string, error) {
	filename := path.Join(dir, kernel.UniqueSuiteFileName(suiteCode, version))
	if suffix != "" {
		filename += suffix
	}
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			zlog.Warnf("suite file %s not found, try to create it", filename)
			var sf zrepo.SuiteFile
			if err := zdb.DB(ctx).Where("suite_code=? and suite_version=?", suiteCode, version).First(&sf).Error; err != nil {
				return "", fmt.Errorf("suite file %s[%s] not found in suite file table", suiteCode, version)
			}
			if err := os.WriteFile(filename, sf.SuiteFile, 0755); err != nil {
				return "", fmt.Errorf("suite actuator creation failure: %s", err.Error())
			}
		} else {
			return "", fmt.Errorf("suite file %s[%s] is err", suiteCode, version)
		}
	}
	return filename, nil
}

// Init
// @Description: 初始化套件管理器，只能全局调用，不能多次调用
// @param suitePath 套件运行路径
// @param max 最大运行数量, 默认为0不限制
func Init(dir string, max int) {
	kernel.Init(dir, max)
}

// LoadSuiteClientWithoutConf
// @Description: 	运行套件（未注入配置，需要自行调用InjectConfig）
// @param ctx
// @param suiteCode	套件编码
// @param version 	套件版本, 默认最新版本
// @return *Entry 	套件运行实例
// @return error 	错误信息
func LoadSuiteClientWithoutConf(suiteCode, version, filename, cmd string) (*kernel.Entry, error) {
	return kernel.LoadSuiteClientWithoutConf(suiteCode, version, filename, cmd)
}

// LoadSuiteInjectInfos
// @Description: 获取套件自身的基础信息
// @param ctx
// @param suiteCode
// @param suitePath
// @param suiteCmd
// @param suiteType
// @return *Suite
// @return error
func LoadSuiteInjectInfos(ctx context.Context, suitePath, suiteCmd, filename string) (*suite.SuiteInfo, error) {
	return kernel.LoadSuiteInjectInfos(ctx, suitePath, suiteCmd, filename)
}

// GetRunPath
// @Description: 获取套件运行器的运行路径
// @return string
func GetRunPath() string {
	return kernel.GetRunPath()
}

// ShutdownSuiteClient
// @Description: 手动关闭套件客户端
// @param uc 启动码
func ShutdownSuiteClient(uc string) {
	kernel.ShutdownSuiteClient(uc)
}
func ShutdownSuiteClientByPid(pid string) {
	kernel.ShutdownSuiteClientByPid(pid)
}
func ShutdownSuiteClientAll() {
	kernel.ShutdownSuiteClientAll()
}
