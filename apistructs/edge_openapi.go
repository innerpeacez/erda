package apistructs

import "fmt"

const (
	EdgeOperationChangePage  = "changePage"
	EdgeOperationChangeRadio = "changeViewType"
	EdgeOperationClick       = "click"
	EdgeOperationSubmit      = "submit"
	EdgeOperationDelete      = "delete"
	EdgeOperationAdd         = "add"
	EdgeOperationUpdate      = "update"
	EdgeOperationViewDetail  = "viewDetail"
	EdgeOperationOffline     = "offline"
	EdgeOperationFilter      = "filter"
	EdgeOperationAddSite     = "addSite"
	EdgeOperationAddCluster  = "addCluster"
	EdgeOperationAddApp      = "addApp"
	EdgeDefaultPageNo        = 1
	EdgeDefaultPageSize      = 20
	EdgeListValueTypeID      = "id"
	EdgeListValueTypeName    = "name"
)

const (
	EdgeDefaultMatchPattern   = "^[a-z0-9-]*$"
	EdgeDefaultRegexpError    = "可输入小写字母、数字或中划线"
	EdgeDefaultNameMaxLength  = 50
	EdgeDefaultValueMaxLength = 100
)

var (
	EdgeDefaultRegexp = fmt.Sprintf("/%v/", EdgeDefaultMatchPattern)
)

type EdgeTableProps struct {
	PageSizeOptions []string      `json:"pageSizeOptions"`
	RowKey          string        `json:"rowKey"`
	Columns         []EdgeColumns `json:"columns"`
}

type EdgeColumns struct {
	Title     string `json:"title"`
	DataIndex string `json:"dataIndex"`
	Width     int    `json:"width"`
}

type EdgeDrawerProps struct {
	Size  string `json:"size"`
	Title string `json:"title"`
}

type EdgeOperations = map[string]interface{}

type EdgeOperation struct {
	Key      string                 `json:"key"`
	Reload   bool                   `json:"reload"`
	Command  EdgeJumpCommand        `json:"command,omitempty"`
	FillMeta string                 `json:"fillMeta,omitempty"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
}

type EdgeItemOperations struct {
	RenderType string                       `json:"renderType"`
	Value      string                       `json:"value,omitempty"`
	Operations map[string]EdgeItemOperation `json:"operations,omitempty"`
	Status     string                       `json:"status,omitempty"`
}

type EdgeTextBadge struct {
	RenderType string `json:"renderType"`
	Value      string `json:"value"`
	Status     string `json:"status,omitempty"`
}

type EdgeItemOperation struct {
	ShowIndex   int                    `json:"showIndex,omitempty"`
	Key         string                 `json:"key"`
	Text        string                 `json:"text"`
	Reload      bool                   `json:"reload"`
	Disabled    bool                   `json:"disabled,omitempty"`
	DisabledTip string                 `json:"disabledTip,omitempty"`
	Confirm     string                 `json:"confirm,omitempty"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
	Command     EdgeJumpCommand        `json:"command,omitempty"`
}

type EdgePageState struct {
	Total    int `json:"total"`
	PageSize int `json:"pageSize"`
	PageNo   int `json:"pageNo"`
}

type EdgeFormModalProps struct {
	Title  string               `json:"title,omitempty"`
	Name   string               `json:"name,omitempty"`
	Fields []EdgeFormModalField `json:"fields"`
}

type EdgeFormModalPointProps struct {
	Title  string                `json:"title,omitempty"`
	Name   string                `json:"name,omitempty"`
	Fields []*EdgeFormModalField `json:"fields"`
}

type EdgeFormModalField struct {
	Key            string                     `json:"key"`
	Label          string                     `json:"label"`
	LabelTip       string                     `json:"labelTip,omitempty"`
	Component      string                     `json:"component"`
	Required       bool                       `json:"required"`
	Rules          []EdgeFormModalFieldRule   `json:"rules,omitempty"`
	Group          string                     `json:"group,omitempty"`
	Disabled       bool                       `json:"disabled"`
	InitialValue   string                     `json:"initialValue,omitempty"`
	DefaultValue   string                     `json:"defaultValue,omitempty"`
	RemoveWhen     [][]map[string]interface{} `json:"removeWhen,omitempty"`
	ComponentProps map[string]interface{}     `json:"componentProps,omitempty"`
}

type EdgeFormModalFieldRule struct {
	Pattern string `json:"pattern"`
	Message string `json:"msg"`
}

type EdgeButtonProps struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type EdgeRadioProps struct {
	RadioType   string             `json:"radioType"`
	ButtonStyle string             `json:"buttonStyle"`
	Size        string             `json:"size"`
	Options     []EdgeButtonOption `json:"options"`
}

type EdgeButtonOption struct {
	Text   string `json:"text"`
	Status string `json:"status"`
	Key    string `json:"key"`
}

type EdgeConditions struct {
	Fixed       bool                     `json:"fixed"`
	EmptyText   string                   `json:"emptyText,omitempty"`
	Key         string                   `json:"key"`
	Label       string                   `json:"label"`
	Type        string                   `json:"type"`
	Placeholder string                   `json:"placeholder"`
	ShowIndex   int                      `json:"showIndex,omitempty"`
	Options     []map[string]interface{} `json:"options,omitempty"`
}

type EdgeEventMeta struct {
	Meta map[string]int64 `json:"meta"`
}

type EdgeSiteState struct {
	SiteID    int64 `json:"siteID"`
	PageNo    int   `json:"pageNo,omitempty"`
	PageSize  int   `json:"pageSize,omitempty"`
	Visible   bool  `json:"visible,omitempty"`
	FormClear bool  `json:"formClear,omitempty"`
}

type EdgeCfgSetState struct {
	ConfigSetItemID int64 `json:"configSetItemID"`
	PageNo          int   `json:"pageNo,omitempty"`
	PageSize        int   `json:"pageSize,omitempty"`
	Visible         bool  `json:"visible,omitempty"`
	FormClear       bool  `json:"formClear,omitempty"`
}

type EdgeAppState struct {
	AppID         uint64 `json:"appID"`
	PageNo        int    `json:"pageNo,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
	Visible       bool   `json:"visible,omitempty"`
	FormClear     bool   `json:"formClear,omitempty"`
	OperationType string `json:"operationType,omitempty"`
}

type EdgeRenderingID struct {
	ID int64 `json:"id"`
}

type EdgeViewGroupSelectState struct {
	ViewGroupSelected string `json:"viewGroupSelected"`
	PageNo            int    `json:"pageNo"`
	PageSize          int    `json:"pageSize"`
}

type EdgeSearchState struct {
	SearchCondition string `json:"searchCondition"`
	PageNo          int    `json:"pageNo"`
	PageSize        int    `json:"pageSize"`
}

type EdgeJumpCommand struct {
	Key     string               `json:"key,omitempty"`
	Target  string               `json:"target,omitempty"`
	JumpOut bool                 `json:"jumpOut"`
	State   EdgeJumpCommandState `json:"state,omitempty"`
}

type EdgeJumpCommandState struct {
	Params   map[string]interface{} `json:"params,omitempty"`
	Query    map[string]interface{} `json:"query,omitempty"`
	Visible  bool                   `json:"visible,omitempty"`
	FormData map[string]interface{} `json:"formData,omitempty"`
	ReadOnly bool                   `json:"readOnly,omitempty"`
}
