package factory

type Host struct {
	AssetNode
}

func (c *Host) GetCnnfEvent() map[string]interface{} {
	return nil
}

func (c *Host) GetRuntimeEvent() map[string]interface{} {
	return nil
}

func (c *Host) GetHipsEvent() map[string]interface{} {
	return nil
}
