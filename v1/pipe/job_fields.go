package pipe

import (
	"fmt"
	"github.com/zj-sh/meta/v1/base"
	"strconv"
	"strings"
)

func JobEffectEnumToString(arr []JobEffect) string {
	var b []string
	for _, v := range arr {
		b = append(b, fmt.Sprintf("%d", v))
	}
	return strings.Join(b, ",")
}
func JobEffectStringToEnum(str string) []JobEffect {
	var b []JobEffect
	for _, v := range strings.Split(str, ",") {
		if d, err := strconv.Atoi(v); err == nil {
			b = append(b, JobEffect(d))
		}
	}
	return b
}

// JobFieldsQuery
// @Description: 获取任务必须的Fields
// @param typ
// @return *JobFields
func JobFieldsQuery(typ *JobEffect) *JobFields {
	switch *typ {
	case JobEffect_EFFECT_SOURCE_PULL:
		return fieldSourcePull()
	case JobEffect_EFFECT_SOURCE_AUTH:
		return fieldSourceAuth()
	case JobEffect_EFFECT_SOURCE_SCAN:
		return fieldSourceScan()
	case JobEffect_EFFECT_CONFIG:
		return fieldConfig()
	case JobEffect_EFFECT_TEST:
		return fieldTest()
	case JobEffect_EFFECT_BUILD:
		return fieldBuild()
	case JobEffect_EFFECT_IMAGE:
		return fieldImage()
	case JobEffect_EFFECT_DEPLOY:
		return fieldDeploy()
	case JobEffect_EFFECT_DATA:
		return fieldData()
	case JobEffect_EFFECT_VERSION:
		return fieldVersion()
	case JobEffect_EFFECT_NOTIFY:
		return fieldNotify()
	case JobEffect_EFFECT_APPROVAL:
		return fieldApproval()
	case JobEffect_EFFECT_STORAGE:
		return fieldStorage()
	case JobEffect_EFFECT_PRODUCT:
		return fieldProduct()
	case JobEffect_EFFECT_SHELL:
		return fieldShell()
	case JobEffect_EFFECT_ISSUES:
		return fieldIssues()
	case JobEffect_EFFECT_ORG:
		return fieldOrg()
	case JobEffect_EFFECT_CUSTOM:
		return fieldCustom()
	case JobEffect_EFFECT_PIPE:
		return fieldPipe()
	default:
		return nil
	}
}

func fieldSourcePull() *JobFields {
	return &JobFields{
		Params: []*base.Field{
			{
				Label:       "授权用户",
				Field:       "src_auth",
				Span:        12,
				Type:        base.FormFieldType_FIELD_SRC_AUTH_USER,
				Required:    true,
				Multivalued: []string{"src_auth_type", "cre_private"},
				Rules: []*base.FieldRule{
					{Required: true, Message: "授权用户不可为空"},
				},
			},
			{
				Label: "代码仓库",
				Field: "src_repo",
				Span:  12,
				Type:  base.FormFieldType_FIELD_SRC_REPO,
				Depend: []*base.FieldDepend{
					{Field: "src_auth", After: true},
					{Field: "src_auth_type", NoEqual: "4"},
				},
				Params: []string{"src_auth"},
				Rules: []*base.FieldRule{
					{Required: true, Message: "代码仓库不可为空"},
				},
				Multivalued:   []string{"ssh_url"},
				AllowVariable: true,
			},
			{
				Label: "代码仓库",
				Field: "ssh_url",
				Span:  12,
				Type:  base.FormFieldType_FIELD_INPUT,
				Depend: []*base.FieldDepend{
					{Field: "src_auth", After: true},
					{Field: "src_auth_type", Equal: "4"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "代码仓库不可为空"},
					{Match: `^git@.+\.git$`, Message: "源码仓库格式不正确"},
				},
				AllowVariable: true,
			},
			{
				Label:    "拉取方式",
				Field:    "src_branch_type",
				Span:     12,
				Type:     base.FormFieldType_FIELD_SELECT,
				Required: true,
				Options: []*base.Option{
					{Label: "Branch", Value: "branch"},
					{Label: "Tag", Value: "tag"},
					{Label: "Commit", Value: "commit"},
				},
				Depend: []*base.FieldDepend{
					{Field: "src_repo", After: true},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "拉取方式不可为空"},
				},
			},
			{
				Label:    "拉取分支",
				Field:    "src_branch",
				Span:     12,
				Type:     base.FormFieldType_FIELD_SRC_BRANCH,
				Required: true,
				Depend: []*base.FieldDepend{
					{Field: "src_repo", After: true},
					{Field: "src_branch_type", Equal: "branch"},
					{Field: "src_auth_type", NoEqual: "4"},
				},
				Params: []string{"src_auth", "src_repo"},
				Rules: []*base.FieldRule{
					{Required: true, Message: "拉取分支不可为空"},
				},
				AllowVariable: true,
			},
			{
				Label: "拉取分支",
				Field: "src_branch",
				Span:  12,
				Type:  base.FormFieldType_FIELD_INPUT,
				Depend: []*base.FieldDepend{
					{Field: "src_repo", After: true},
					{Field: "src_branch_type", Equal: "branch"},
					{Field: "src_auth_type", Equal: "4"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "拉取分支不可为空"},
				},
				AllowVariable: true,
			},
			{
				Label:    "拉取标签",
				Field:    "src_tag",
				Span:     12,
				Type:     base.FormFieldType_FIELD_SRC_TAG,
				Required: true,
				Depend: []*base.FieldDepend{
					{Field: "src_repo", After: true},
					{Field: "src_branch_type", Equal: "tag"},
					{Field: "src_auth_type", NoEqual: "4"},
				},
				Params: []string{"src_auth", "src_repo"},
				Rules: []*base.FieldRule{
					{Required: true, Message: "拉取标签不可为空"},
				},
			},
			{
				Label: "拉取标签",
				Field: "src_tag",
				Span:  12,
				Type:  base.FormFieldType_FIELD_INPUT,
				Depend: []*base.FieldDepend{
					{Field: "src_repo", After: true},
					{Field: "src_branch_type", Equal: "tag"},
					{Field: "src_auth_type", Equal: "4"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "拉取标签不可为空"},
				},
			},
			{
				Label: "拉取commit",
				Field: "src_commit",
				Span:  12,
				Type:  base.FormFieldType_FIELD_INPUT,
				Depend: []*base.FieldDepend{
					{Field: "src_repo", After: true},
					{Field: "src_branch_type", Equal: "commit"},
					{Field: "src_auth_type", Equal: "4"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "拉取commit不可为空"},
				},
			},
		},
		Outs: []*base.LabelField{
			{Field: "SOURCE_PATH", Label: "源码路径"},
		},
	}
}
func fieldSourceAuth() *JobFields {
	return &JobFields{
		Params: []*base.Field{
			{
				Label:    "授权方式",
				Field:    "auth_type",
				DefValue: "1",
				Span:     24,
				Type:     base.FormFieldType_FIELD_SRC_AUTH_TYPE,
				Params:   []string{"suite_code"},
				Depend: []*base.FieldDepend{
					{Field: "suite_code", After: true},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "授权方式不可为空"},
				},
			},
			{
				Label: "用户名",
				Field: "auth_user",
				Span:  24,
				Type:  base.FormFieldType_FIELD_INPUT,
				Depend: []*base.FieldDepend{
					{Field: "suite_code", After: true},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "用户名不能为空"},
				},
			},
			{
				Label: "密码",
				Field: "auth_password",
				Span:  24,
				Type:  base.FormFieldType_FIELD_PWD,
				Depend: []*base.FieldDepend{
					{Field: "auth_type", Equal: "2"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "密码不能为空"},
				},
			},
			{
				Label: "令牌",
				Field: "auth_token",
				Span:  24,
				Type:  base.FormFieldType_FIELD_TEXT,
				Depend: []*base.FieldDepend{
					{Field: "auth_type", Equal: "3"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "令牌不可为空"},
				},
			},
			{
				Label:    "warning",
				DefValue: "若选择SSH私钥模式,则不支持查询仓库、分支、commit等信息,相关选项需自行填写",
				Type:     base.FormFieldType_FIELD_ALERT,
				Depend: []*base.FieldDepend{
					{Field: "auth_type", Equal: "4"},
				},
			},
			{
				Label: "SSH私钥",
				Field: "auth_ssh",
				Span:  24,
				Type:  base.FormFieldType_FIELD_TEXT,
				Depend: []*base.FieldDepend{
					{Field: "auth_type", Equal: "4"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "SSH私钥不能为空"},
				},
			},
			{
				Label:    "授权时长",
				Field:    "auth_expire_mode",
				Span:     24,
				DefValue: "long",
				Type:     base.FormFieldType_FIELD_SELECT,
				Options: []*base.Option{
					{Label: "长期", Value: "long"},
					{Label: "临时", Value: "short"},
				},
				Depend: []*base.FieldDepend{
					{Field: "suite_code", After: true},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "授权时长不能为空"},
				},
			},
			{
				Label:    "到期时间",
				Field:    "auth_expire_time",
				Span:     24,
				DefValue: "end_day",
				Type:     base.FormFieldType_FIELD_TIME,
				Depend: []*base.FieldDepend{
					{Field: "auth_expire_mode", Equal: "short"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "授权到期时间不能为空"},
				},
			},
		},
		Outs: []*base.LabelField{},
	}
}
func fieldSourceScan() *JobFields {
	return &JobFields{
		Params: []*base.Field{
			{
				Label:    "代码路径",
				Field:    "source",
				Span:     24,
				Type:     base.FormFieldType_FIELD_INPUT,
				Required: true,
				DefValue: "./",
				Rules: []*base.FieldRule{
					{Required: true, Message: "代码路径不能为空"},
				},
			},
			{
				Label:   "排除路径",
				Field:   "exclude",
				Span:    24,
				Type:    base.FormFieldType_FIELD_TEXT,
				Tooltip: "每行一个",
			},
		},
		Outs: []*base.LabelField{
			{Field: "SOURCE_SCAN_REPORT_PATH", Label: "报告文件"},
			{Field: "SOURCE_SCAN_REPORT_URL", Label: "报告地址"},
		},
	}
}
func fieldConfig() *JobFields {
	return &JobFields{
		Params: []*base.Field{},
		Outs:   []*base.LabelField{},
	}
}
func fieldTest() *JobFields {
	return &JobFields{
		Params: []*base.Field{
			{
				Label:    "测试脚本",
				Field:    "cmd",
				Span:     24,
				Type:     base.FormFieldType_FIELD_CODE,
				Required: true,
				Rules: []*base.FieldRule{
					{Required: true, Message: "测试脚本不能为空"},
				},
			},
			{
				Label:    "超时时间",
				Field:    "timeout",
				Span:     24,
				Type:     base.FormFieldType_FIELD_NUMBER,
				DefValue: "0",
				Rules: []*base.FieldRule{
					{Min: 0, Message: "超时时间不能小于0"},
				},
			},
			{
				Label:    "测试报告目录",
				Field:    "report_dir",
				Span:     24,
				Type:     base.FormFieldType_FIELD_INPUT,
				Required: true,
				DefValue: "reports/test",
				Rules: []*base.FieldRule{
					{Required: true, Message: "测试报告目录不能为空"},
				},
			},
		},
		Outs: []*base.LabelField{
			{Field: "TEST_REPORT_DIR", Label: "测试报告路径"},
		},
	}
}
func fieldBuild() *JobFields {
	return &JobFields{
		Params: []*base.Field{
			{
				Label:    "构建方式",
				Field:    "build_type",
				Span:     24,
				Type:     base.FormFieldType_FIELD_SELECT,
				DefValue: "BUILD_TYPE_IN_IMAGE",
				Options: []*base.Option{
					{Label: "常规构建", Value: "BUILD_TYPE_IN_IMAGE"},
					{Label: "主机构建", Value: "BUILD_TYPE_IN_MACHINE"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "构建方式不能为空"},
				},
			},
			{
				Label:    "构建镜像",
				Field:    "build_in",
				Span:     24,
				Type:     base.FormFieldType_FIELD_BUILD_IMAGE,
				Required: true,
				Depend: []*base.FieldDepend{
					{Field: "build_type", Equal: "BUILD_TYPE_IN_IMAGE"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "构建镜像不能为空"},
				},
			},
			{
				Label:    "构建主机",
				Field:    "build_in",
				Span:     24,
				Type:     base.FormFieldType_FIELD_BUILD_MACHINE,
				Required: true,
				Depend: []*base.FieldDepend{
					{Field: "build_type", Equal: "BUILD_TYPE_IN_MACHINE"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "构建主机不能为空"},
				},
			},
			{
				Label:    "构建物路径",
				Field:    "build_out_path",
				Span:     24,
				Type:     base.FormFieldType_FIELD_INPUT,
				Required: true,
				Rules: []*base.FieldRule{
					{Required: true, Message: "构建物路径不能为空"},
				},
			},
		},
		Outs: []*base.LabelField{
			{Field: "BUILD_OUT_PATH", Label: "构建物路径"},
		},
	}
}
func fieldImage() *JobFields {
	return &JobFields{
		Params: []*base.Field{
			{
				Label:       "镜像仓库",
				Field:       "repo",
				Span:        24,
				Type:        base.FormFieldType_FIELD_BUILD_IMAGE_REPO,
				Required:    true,
				Multivalued: []string{"registry_auth"},
				Rules: []*base.FieldRule{
					{Required: true, Message: "镜像仓库不能为空"},
				},
			},
			{
				Label:   "镜像名称",
				Field:   "name",
				Span:    24,
				Type:    base.FormFieldType_FIELD_INPUT,
				Tooltip: "不含仓库名，如：nginx",
				Rules: []*base.FieldRule{
					{Required: true, Message: "镜像名称不能为空"},
				},
			},
			{
				Label:    "镜像标签",
				Field:    "tag",
				Span:     24,
				Type:     base.FormFieldType_FIELD_INPUT,
				DefValue: "latest",
				Tooltip:  "可使用变量，如：${BUILD_NUMBER}",
				Rules: []*base.FieldRule{
					{Required: true, Message: "镜像标签不能为空"},
				},
			},
			{
				Label:    "Dockerfile路径",
				Field:    "dockerfile",
				Span:     24,
				Type:     base.FormFieldType_FIELD_INPUT,
				DefValue: "./Dockerfile",
				Tooltip:  "可使用变量",
				Rules: []*base.FieldRule{
					{Required: true, Message: "Dockerfile路径不能为空"},
				},
			},
			{
				Label:   "构建上下文路径",
				Field:   "context_path",
				Span:    24,
				Type:    base.FormFieldType_FIELD_INPUT,
				Tooltip: "若为空则和dockerfile路径相同",
			},
			{
				Label:    "构建架构",
				Field:    "arch",
				Span:     24,
				Type:     base.FormFieldType_FIELD_SELECT,
				Multiple: true,
				Tooltip:  "为空则默认同构建环境架构，若构建环境架构为amd64，则默认构建架构为amd64",
				Options: []*base.Option{
					{Label: "amd64", Value: "BUILD_ARCH_AMD64"},
					{Label: "arm64", Value: "BUILD_ARCH_ARM64"},
				},
			},
			{
				Label:    "自动推送",
				Field:    "auto_push",
				Span:     24,
				Type:     base.FormFieldType_FIELD_SWITCH,
				DefValue: "true",
			},
		},
		Outs: []*base.LabelField{},
	}
}
func fieldDeploy() *JobFields {
	return &JobFields{
		Params: []*base.Field{},
		Outs:   []*base.LabelField{},
	}
}
func fieldData() *JobFields {
	return &JobFields{
		Params: []*base.Field{},
		Outs:   []*base.LabelField{},
	}
}
func fieldVersion() *JobFields {
	return &JobFields{
		Params: []*base.Field{},
		Outs:   []*base.LabelField{},
	}
}
func fieldNotify() *JobFields {
	return &JobFields{
		Params: []*base.Field{
			{
				Label:         "通知标题",
				Field:         "title",
				Span:          24,
				Type:          base.FormFieldType_FIELD_INPUT,
				Required:      true,
				AllowVariable: true,
				Rules: []*base.FieldRule{
					{Required: true, Message: "通知标题不能为空"},
				},
			},
			{
				Label:         "通知内容",
				Field:         "content",
				Span:          24,
				Type:          base.FormFieldType_FIELD_TEXT,
				Required:      true,
				AllowVariable: true,
				Rules: []*base.FieldRule{
					{Required: true, Message: "通知内容不能为空"},
				},
			},
			{
				Label:    "通知级别",
				Field:    "level",
				Span:     24,
				Type:     base.FormFieldType_FIELD_SELECT,
				DefValue: "INFO",
				Options: []*base.Option{
					{Label: "普通", Value: "INFO"},
					{Label: "重要", Value: "WARN"},
					{Label: "紧急", Value: "ERROR"},
					{Label: "成功", Value: "SUCCESS"},
				},
				Required:      true,
				AllowVariable: true,
				Rules: []*base.FieldRule{
					{Required: true, Message: "通知级别不能为空"},
				},
			},
			{
				Label:         "通知跳转",
				Field:         "link",
				Span:          24,
				AllowVariable: true,
				Type:          base.FormFieldType_FIELD_INPUT,
			},
			{
				Label: "接收人类型",
				Field: "to_type",
				Span:  24,
				Type:  base.FormFieldType_FIELD_SELECT,
				Options: []*base.Option{
					{Label: "指定人员", Value: "NOTIFY_TO_TYPE_USER"},
					{Label: "指定机构", Value: "NOTIFY_TO_TYPE_ORG"},
					{Label: "指定角色", Value: "NOTIFY_TO_TYPE_ROLE"},
					{Label: "自定义", Value: "NOTIFY_TO_TYPE_CUSTOM"},
				},
			},
			{
				Label: "接收人",
				Field: "to",
				Span:  24,
				Type:  base.FormFieldType_FIELD_CORE_USER,
				Depend: []*base.FieldDepend{
					{Field: "to_type", Equal: "NOTIFY_TO_TYPE_USER"},
				},
				Multiple: true,
				Required: true,
				Tooltip:  "第一个为主要接收人，第二个起为抄送",
				Rules: []*base.FieldRule{
					{Required: true, Message: "接收人不能为空"},
				},
			},
			{
				Label: "接收组织",
				Field: "to",
				Span:  24,
				Type:  base.FormFieldType_FIELD_CORE_ORG,
				Depend: []*base.FieldDepend{
					{Field: "to_type", Equal: "NOTIFY_TO_TYPE_ORG"},
				},
				Multiple: true,
				Required: true,
				Rules: []*base.FieldRule{
					{Required: true, Message: "接收组织不能为空"},
				},
			},
			{
				Label: "接收角色",
				Field: "to",
				Span:  24,
				Type:  base.FormFieldType_FIELD_CORE_ROLE,
				Depend: []*base.FieldDepend{
					{Field: "to_type", Equal: "NOTIFY_TO_TYPE_ROLE"},
				},
				Multiple: true,
				Required: true,
				Rules: []*base.FieldRule{
					{Required: true, Message: "接收角色不能为空"},
				},
			},
			{
				Label: "接收人",
				Field: "to",
				Span:  24,
				Type:  base.FormFieldType_FIELD_INPUT,
				Depend: []*base.FieldDepend{
					{Field: "to_type", Equal: "NOTIFY_TO_TYPE_CUSTOM"},
				},
				Tooltip:  "多个接收人用逗号分隔, 值内容类型取决于套件",
				Required: true,
				Rules: []*base.FieldRule{
					{Required: true, Message: "接收人不能为空"},
				},
			},
		},
		Outs: []*base.LabelField{},
	}
}
func fieldApproval() *JobFields {
	return &JobFields{
		Params: []*base.Field{},
		Outs:   []*base.LabelField{},
	}
}
func fieldStorage() *JobFields {
	return &JobFields{
		Params: []*base.Field{},
		Outs:   []*base.LabelField{},
	}
}
func fieldProduct() *JobFields {
	return &JobFields{
		Params: []*base.Field{
			{
				Label: "构建物路径",
				Field: "filepath",
				Span:  24,
				Type:  base.FormFieldType_FIELD_INPUT,
				Rules: []*base.FieldRule{
					{Required: true, Message: "构建物路径不能为空"},
				},
			},
			{
				Label: "产物名称",
				Field: "pkg_name",
				Span:  24,
				Type:  base.FormFieldType_FIELD_INPUT,
				Rules: []*base.FieldRule{
					{Required: true, Message: "产物名称不能为空"},
				},
			},
			{
				Label: "产物版本",
				Field: "pkg_version",
				Span:  24,
				Type:  base.FormFieldType_FIELD_INPUT,
				Rules: []*base.FieldRule{
					{Required: true, Message: "产物版本不能为空"},
				},
			},
		},
		Outs: []*base.LabelField{
			{Field: "PRODUCT_URL", Label: "产物地址"},
		},
	}
}
func fieldShell() *JobFields {
	return &JobFields{
		Params: []*base.Field{
			{
				Label:    "脚本类型",
				Field:    "command",
				Span:     24,
				Type:     base.FormFieldType_FIELD_SELECT,
				DefValue: "shell",
				Required: true,
				Options: []*base.Option{
					{Label: "Shell", Value: "shell"},
					{Label: "Bash", Value: "bash"},
				},
				Rules: []*base.FieldRule{
					{Required: true, Message: "脚本类型不能为空"},
				},
			},
			{
				Label:    "脚本执行目录",
				Field:    "cwd",
				Span:     24,
				Type:     base.FormFieldType_FIELD_INPUT,
				Required: true,
				Rules: []*base.FieldRule{
					{Required: true, Message: "脚本执行目录不能为空"},
				},
			},
			{
				Label:    "脚本内容",
				Field:    "stdin",
				Span:     24,
				Type:     base.FormFieldType_FIELD_CODE,
				Required: true,
				Rules: []*base.FieldRule{
					{Required: true, Message: "脚本内容不能为空"},
				},
			},
			{
				Label:    "超时时间",
				Field:    "timeout",
				Span:     24,
				Type:     base.FormFieldType_FIELD_NUMBER,
				DefValue: "0",
				Rules: []*base.FieldRule{
					{Min: 0, Message: "超时时间不能小于0"},
				},
			},
		},
		Outs: []*base.LabelField{},
	}
}
func fieldIssues() *JobFields {
	return &JobFields{
		Params: []*base.Field{},
		Outs:   []*base.LabelField{},
	}
}
func fieldOrg() *JobFields {
	return &JobFields{
		Params: []*base.Field{},
		Outs:   []*base.LabelField{},
	}
}
func fieldCustom() *JobFields {
	return &JobFields{
		Params: []*base.Field{},
		Outs:   []*base.LabelField{},
	}
}
func fieldPipe() *JobFields {
	return &JobFields{
		Params: []*base.Field{},
		Outs:   []*base.LabelField{},
	}
}
