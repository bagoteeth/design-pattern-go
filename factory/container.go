package factory

type Container struct {
	AssetNode
}

func (c *Container) GetCnnfEvent() map[string]interface{} {
	return map[string]interface{}{"container1": "cnnf1"}
}

func (c *Container) GetRuntimeEvent() map[string]interface{} {
	return map[string]interface{}{"container1": "runtime1"}
}

func (c *Container) GetHipsEvent() map[string]interface{} {
	return map[string]interface{}{"container1": "hips1"}
}
