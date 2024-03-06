package svrs

// ServiceTypeOptions
// @Description: 服务类型选项
// @return []*ServiceTypeOption
func ServiceTypeOptions() []*ServiceTypeOption {
	return []*ServiceTypeOption{
		{Label: "Kubernetes服务", Value: ServiceType_KUBERNETES},
		{Label: "Helm Chart服务", Value: ServiceType_HELM},
		{Label: "Docker服务", Value: ServiceType_DOCKER},
		{Label: "主机服务", Value: ServiceType_MACHINE},
		{Label: "其他服务", Value: ServiceType_OTHER},
	}
}
