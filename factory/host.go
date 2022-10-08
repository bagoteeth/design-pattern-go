package factory

type Host struct {
	AssetNode
}

func (c *Host) GetCnnfEvent() map[string]interface{} {
	return map[string]interface{}{"host1": "cnnf1"}
}

func (c *Host) GetRuntimeEvent() map[string]interface{} {
	return map[string]interface{}{"host1": "runtime1"}
}

func (c *Host) GetHipsEvent() map[string]interface{} {
	return map[string]interface{}{"host1": "hips1"}
}
