package kernel

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-plugin"
	"github.com/zj-sh/meta/v1/suite"
	"github.com/zohu/zlog"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"path"
	"reflect"
	"time"
)

/**
 * @description: no use database
 */

var m = new(manager)

// Init
// @Description: 初始化套件管理器，只能全局调用，不能多次调用
// @param suitePath 套件运行路径
// @param max 最大运行数量, 默认为0不限制
func Init(dir string, max int) {
	if m.path != "" {
		return
	}
	m.path = dir
	if !path.IsAbs(m.path) {
		zlog.Fatal("suite path must be absolute path")
		return
	}
	m.max = max
	m.suites = make(map[string]*Entry)
	unix := path.Join(m.path, ".grpc")
	_ = os.MkdirAll(unix, os.ModePerm)
	_ = os.Setenv(plugin.EnvUnixSocketDir, unix)
	go m.onStatus()
}

// LoadSuiteClientWithoutConf
// @Description: 	运行套件（未注入配置，需要自行调用InjectConfig）
// @param ctx
// @param suiteCode	套件编码
// @param version 	套件版本, 默认最新版本
// @return *Entry 	套件运行实例
// @return error 	错误信息
func LoadSuiteClientWithoutConf(suiteCode, version, filename, cmd string) (*Entry, error) {
	if m.path == "" {
		return nil, fmt.Errorf("suite path is not initialized, please call cube.Init or kernel.Init globally first")
	}
	// 套件唯一启动码
	uc := UniqueSuiteFileName(suiteCode, version)
	m.Lock()
	defer m.Unlock()
	// 如果已经运行, 则直接返回
	if e, ok := m.load(uc); ok {
		return e, nil
	}
	// 判断是否达到最大数量
	if m.isMaxRunning() {
		zlog.Warnf("当前运行中套件已达到最大数量，启动失败: 限制%d 运行中%d", m.max, len(m.suites))
		return nil, fmt.Errorf("the running suite has reached its maximum number and the startup fails")
	}
	// 否则启动套件
	zlog.Infof("start suite: %s[%s]", suiteCode, version)

	clt, grpc, err := suite.GRPCSuiteClient(m.path, cmd, filename)
	if err != nil {
		return nil, fmt.Errorf("start suite client error: %s", err.Error())
	}
	e := &Entry{
		pid:     grpc.ID(),
		code:    suiteCode,
		version: version,
		start:   time.Now(),
		last:    time.Now(),
		delay:   time.Minute * 5,
		grpc:    grpc,
		client:  clt,
		file:    filename,
	}
	m.run(uc, e)
	return e, nil
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
	clt, grpc, err := suite.GRPCSuiteClient(suitePath, suiteCmd, filename)
	if err != nil {
		return nil, fmt.Errorf("start suite client error: %s", err.Error())
	}
	defer grpc.Kill()

	info, err := clt.Inject(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("load suite inject error: %s", err.Error())
	}

	if info == nil {
		zlog.Warnf("load inject suite error, and clt type: %s", reflect.TypeOf(clt).String())
		return nil, fmt.Errorf("load inject suite error: suite info is empty")
	}
	return info, nil
}

// GetRunPath
// @Description: 获取套件运行器的运行路径
// @return string
func GetRunPath() string {
	return m.path
}

// ShutdownSuiteClient
// @Description: 手动关闭套件客户端
// @param uc 启动码
func ShutdownSuiteClient(uc string) {
	m.shutdown(uc)
}
func ShutdownSuiteClientByPid(pid string) {
	m.shutdownPid(pid)
}
func ShutdownSuiteClientAll() {
	m.shutdownAll()
}
