package memos

// TODO: Define the missing types
type Memo struct {
	DisplayTime string           `json:"displayTime"`
	Visibility  string           `json:"visibility"`
	RowStatus   string           `json:"rowStatus"`
	Creator     string           `json:"creator"`
	CreateTime  string           `json:"createTime"`
	UpdateTime  string           `json:"updateTime"`
	Uid         string           `json:"uid"`
	Content     string           `json:"content"`
	Name        string           `json:"name"`
	Tags        []string         `json:"tags"`
	Nodes       []map[string]any `json:"nodes"`
	Resources   []map[string]any `json:"resources"`
	Relations   []map[string]any `json:"relations"`
	Reactions   []map[string]any `json:"reactions"`
	Pinned      bool             `json:"pinned"`
}
