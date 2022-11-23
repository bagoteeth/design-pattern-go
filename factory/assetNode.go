package factory

const (
	TYPE_HOST      = "host"
	TYPE_APP       = "app"
	TYPE_CONTAINER = "container"
)

type AssetNodeIntf interface {
	GetCnnfEvent() map[string]interface{}
	GetRuntimeEvent() map[string]interface{}
	GetHipsEvent() map[string]interface{}
}

type AssetNode struct {
	NodeType  string `json:"nodeType"`
	Cluster   string `json:"cluster"`
	HostGroup string `json:"hostGroup"`
	Namespace string `json:"namespace"`
	Host      string `json:"host"`
	Container string `json:"container"`
	App       string `json:"app"`
}

func (node *AssetNode) Transform() AssetNodeIntf {
	switch node.NodeType {
	case TYPE_HOST:
		return &Host{AssetNode: *node}
	case TYPE_APP:
		return &App{AssetNode: *node}
	case TYPE_CONTAINER:
		return &Container{AssetNode: *node}
	}
	return nil
}
