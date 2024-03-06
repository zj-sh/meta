package pipe

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

// Verify
// @Description: 变量验证
// @receiver s
// @return error
func (v *Var) Verify() error {
	if v.GetVariable() == "" {
		return fmt.Errorf("变量名不能为空")
	}
	if v.GetGroupId() == "" && !v.GetIsRunning() && v.GetValue() == "" {
		return fmt.Errorf("非运行时变量值不能为空")
	}
	return nil
}

func (p *Pipeline) Verify() error {
	if p.PipelineName == "" {
		return errors.New("流水线名称不能为空")
	}
	if p.PipelineType == PipelineType_SERVICES && p.ServiceId == "" {
		return errors.New("服务类流水线必须选择所属服务")
	}
	if p.PipelineType == PipelineType_SERVICES && p.EnvId == "" {
		return errors.New("服务类流水线必须选择所属环境")
	}
	if len(p.PipelineJobs) <= 2 {
		return errors.New("流水线任务不能为空")
	}
	if p.PipelineHook.CronEnable && (p.PipelineHook.CronMode == "" || p.PipelineHook.CronExpression == "") {
		return errors.New("定时模式和定时表达式不能为空")
	}
	if p.PipelineHook.WebhookEnable && p.PipelineHook.WebhookUri == "" {
		return errors.New("webhook地址不能为空")
	}
	for _, v := range p.PipelineVariables {
		if v.Variable == "" {
			return errors.New("变量名不能为空")
		}
	}
	for _, v := range p.PipelineCache.CacheBuild {
		if v.Path == "" {
			return errors.New("构建缓存路径不能为空")
		}
	}
	for _, v := range p.PipelineCache.CacheDependent {
		if v.Path == "" {
			return errors.New("依赖缓存路径不能为空")
		}
	}
	for _, v := range p.PipelineCache.CacheProduct {
		if v.Path == "" {
			return errors.New("产物缓存路径不能为空")
		}
	}
	if !p.AllowAllCorp {
		for _, v := range p.PipelineRole.Roles {
			if v.RoleType == "" || v.RoleCode == "" {
				return errors.New("自定义权限时，授权类型和主体不能为空")
			}
		}
	}
	return nil
}
func (p *Pipeline) Yaml() string {
	y, _ := yaml.Marshal(p)
	return string(y)
}
func (p *Pipeline) UnmarshalYaml(y []byte) error {
	return yaml.Unmarshal(y, p)
}
func (p *Pipeline) Copy() (*Pipeline, error) {
	y := p.Yaml()
	p2 := &Pipeline{}
	err := p2.UnmarshalYaml([]byte(y))
	return p2, err
}
