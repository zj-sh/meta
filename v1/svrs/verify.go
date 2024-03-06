package svrs

// Verify
// @Description: 服务验证
// @receiver s
// @return error
func (s *Service) Verify() error {
	for _, v := range s.GetServiceVariables() {
		if err := v.Verify(); err != nil {
			return err
		}
	}
	for _, env := range s.GetServiceEnvs() {
		if err := env.Verify(); err != nil {
			return err
		}
	}
	return nil
}
