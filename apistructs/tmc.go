package apistructs

// RuntimeServiceRequest 部署runtime之后，orchestrator需要将服务域名信息通过此接口提交给hepa
type RuntimeServiceRequest struct {
	// OrgID 企业ID
	OrgID string `json:"orgId"`
	// ProjectID 项目ID
	ProjectID string `json:"projectId"`
	// Workspace 所属环境
	Workspace string `json:"env"`
	// CluserName 集群名称
	ClusterName string `json:"clusterName"`
	// RuntimeID runtimeID
	RuntimeID string `json:"runtimeId"`
	// RuntimeName runtime名称
	RuntimeName string `json:"runtimeName"`
	// AppID 应用ID
	AppID string `json:"appId"`
	// AppName 应用名称
	AppName string `json:"appName"`
	// Services 服务组成的列表
	Services []ServiceItem `json:"services"`
	// UseApigw 是否通过addon依赖了api网关
	UseApigw bool `json:"useApigw"`
	// ReleaseId
	ReleaseID string `json:"releaseId"`
	// ServiceGroupNamespace
	ServiceGroupNamespace string `json:"serviceGroupNamespace"`
	// ServiceGroupName
	ServiceGroupName string `json:"serviceGroupName"`
}

// ServiceItem service信息
type ServiceItem struct {
	// ServiceName 服务名称
	ServiceName string `json:"serviceName"`
	// InnerAddress 服务内部地址
	InnerAddress string `json:"innerAddress"`
}

// EndpointDomainsItem 对外暴露地址信息
type EndpointDomainsItem struct {
	// Domain 域名
	Domain string `json:"domain"`
	// Type 域名类型,CUSTOM or DEFAULT
	Type string `json:"type"`
}

// TenantGroupDetailsResponse .
type TenantGroupDetailsResponse struct {
	Header
	Data TenantGroupDetails `json:"data"`
}

// TenantGroupDetails .
type TenantGroupDetails struct {
	ProjectID string `json:"projectID"`
}

// MonitorStatusMetricDetailsResponse .
type MonitorStatusMetricDetailsResponse struct {
	Header
	Data MonitorStatusMetricDetails `json:"data"`
}

// MonitorStatusMetricDetails .
type MonitorStatusMetricDetails struct {
	ProjectID int64  `json:"projectId"`
	Name      string `json:"name"`
}
