package pipe

import "slices"

type JobAttribOption struct {
	Label string    `json:"label"` // 任务类型
	Value JobAttrib `json:"value"` // 任务类型
}

func JobAttribOptions() []JobAttribOption {
	return []JobAttribOption{
		{Label: "平台", Value: JobAttrib_ATTRIB_PLATFORM},
		{Label: "流水线", Value: JobAttrib_ATTRIB_PIPELINE},
	}
}

type JobEffectOption struct {
	Label string    `json:"label"` // 任务类型
	Value JobEffect `json:"value"` // 任务类型
	Func  string    `json:"func"`  // 函数
	Color string    `json:"color"` // 颜色
	Desc  string    `json:"desc"`  // 描述
}

func JobEffectAllOptions() []JobEffectOption {
	return slices.Concat(
		JobEffectOptions(JobAttrib_ATTRIB_PLATFORM),
		JobEffectOptions(JobAttrib_ATTRIB_PIPELINE),
	)
}

// JobEffectOptions
// @Description: 指定套件属性查询套件类型下拉选项
// @param typ
// @return []PipelineJobEffectOption
func JobEffectOptions(attrib JobAttrib) []JobEffectOption {
	if attrib == JobAttrib_ATTRIB_PLATFORM {
		return []JobEffectOption{
			{Label: "仓库授权", Value: JobEffect_EFFECT_SOURCE_AUTH, Color: "arcoblue", Desc: "代码仓库授权访问"},
			{Label: "用户同步", Value: JobEffect_EFFECT_ORG, Color: "red", Desc: "组织架构同步"},
		}
	}
	if attrib == JobAttrib_ATTRIB_PIPELINE {
		return []JobEffectOption{
			{Label: "代码拉取", Value: JobEffect_EFFECT_SOURCE_PULL, Func: "SourceClone", Color: "arcoblue", Desc: "从源码仓库拉取代码"},
			{Label: "代码校验", Value: JobEffect_EFFECT_SOURCE_SCAN, Func: "SourceScan", Color: "blue", Desc: "代码扫描等"},
			{Label: "配置管理", Value: JobEffect_EFFECT_CONFIG, Func: "Config", Color: "cyan", Desc: "配置文件管理"},
			{Label: "测试任务", Value: JobEffect_EFFECT_TEST, Func: "TestRun", Color: "purple", Desc: "单元测试、接口测试等"},
			{Label: "构建任务", Value: JobEffect_EFFECT_BUILD, Func: "Build", Color: "pinkpurple", Desc: "构建、打包等"},
			{Label: "镜像管理", Value: JobEffect_EFFECT_IMAGE, Func: "ImageBuild", Color: "magenta", Desc: "镜像生成、镜像上传等"},
			{Label: "部署发布", Value: JobEffect_EFFECT_DEPLOY, Func: "Deploy", Color: "green", Desc: "主机发布、集群发布等"},
			{Label: "数据管理", Value: JobEffect_EFFECT_DATA, Func: "", Color: "gold", Desc: "库表同步、数据版本、数据加工等"},
			{Label: "版本采集", Value: JobEffect_EFFECT_VERSION, Func: "Version", Color: "lime", Desc: "版本管理、迭代说明等"},
			{Label: "文件存储", Value: JobEffect_EFFECT_STORAGE, Func: "StorageUploadFile", Color: "red", Desc: "上传、保存等"},
			{Label: "上传制品", Value: JobEffect_EFFECT_PRODUCT, Func: "ProductUploadFile", Color: "arcoblue", Desc: "上传制品到制品库"},
			{Label: "执行脚本", Value: JobEffect_EFFECT_SHELL, Func: "ShellExec", Color: "blue", Desc: "执行SHELL脚本"},
			{Label: "任务协同", Value: JobEffect_EFFECT_ISSUES, Func: "Issues", Color: "cyan", Desc: "项目任务协同"},
			{Label: "审批流程", Value: JobEffect_EFFECT_APPROVAL, Func: "ApprovalSend", Color: "orangered", Desc: "审批流程"},
			{Label: "通知提醒", Value: JobEffect_EFFECT_NOTIFY, Func: "Notify", Color: "arcoblue", Desc: "短信、微信、钉钉、飞书等通知类"},
			{Label: "自定义", Value: JobEffect_EFFECT_CUSTOM, Func: "Custom", Color: "purple", Desc: "多功能自定义套件"},
			{Label: "流水线", Value: JobEffect_EFFECT_PIPE, Func: "PipeRun", Color: "red", Desc: "调度、临时修改流水线等"},
		}
	}
	return []JobEffectOption{}
}
