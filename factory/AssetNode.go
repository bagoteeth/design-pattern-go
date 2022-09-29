package factory

const (
	TYPE_HOST      = "host"
	TYPE_APP       = "app"
	TYPE_CONTAINER = "container"
)

type AssetNode struct {
	NodeType  string `json:"nodeType"`
	Cluster   string `json:"cluster"`
	HostGroup string `json:"hostGroup"`
	Namespace string `json:"namespace"`
	Host      string `json:"host"`
	Container string `json:"container"`
	App       string `json:"app"`
}
