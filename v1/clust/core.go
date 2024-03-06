package clust

func ClusterTypeOptions() []*ClusterTypeOption {
	return []*ClusterTypeOption{
		{Label: "阿里云CCE集群", Value: ClusterType_CLUSTER_CCE},
		{Label: "华为云ACK集群", Value: ClusterType_CLUSTER_ACK},
		{Label: "腾讯云TKE集群", Value: ClusterType_CLUSTER_TKE},
		{Label: "自建集群", Value: ClusterType_CLUSTER_SELF},
	}
}

func LinkTypeOptions() []*LinkTypeOption {
	return []*LinkTypeOption{
		{Label: "KubeConfig", Value: LinkType_LINK_KUBE_CONFIG},
	}
}
