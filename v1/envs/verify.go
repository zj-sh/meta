package envs

import "fmt"

// Verify
// @Description: 环境验证
// @receiver env
// @return error
func (env *Env) Verify() error {
	if env.GetEnvName() == "" {
		return fmt.Errorf("环境名称不能为空")
	}
	if env.GetEnvCode() == "" {
		return fmt.Errorf("环境编码不能为空")
	}
	if env.GetClusterId() == "" {
		return fmt.Errorf("集群不能为空")
	}
	if env.GetProjectId() == "" {
		return fmt.Errorf("项目不能为空")
	}
	if env.GetNamespace() == "" {
		return fmt.Errorf("命名空间不能为空")
	}
	for _, v := range env.GetEnvVars() {
		if err := v.Verify(); err != nil {
			return err
		}
	}
	return nil
}
