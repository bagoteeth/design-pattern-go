package factory

type AssetNodeIntf interface {
	GetCnnfEvent() map[string]interface{}
	GetRuntimeEvent() map[string]interface{}
	GetHipsEvent() map[string]interface{}
}

func GetAssetNodeData(node AssetNodeIntf) (cnnf, runtime, hips map[string]interface{}) {
	cnnf = node.GetCnnfEvent()
	runtime = node.GetRuntimeEvent()
	hips = node.GetHipsEvent()
	return
}
