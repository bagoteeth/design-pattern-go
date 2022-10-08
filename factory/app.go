package factory

type App struct {
	AssetNode
}

func (c *App) GetCnnfEvent() map[string]interface{} {
	return map[string]interface{}{"app1": "cnnf1"}
}

func (c *App) GetRuntimeEvent() map[string]interface{} {
	return map[string]interface{}{"app1": "runtime1"}
}

func (c *App) GetHipsEvent() map[string]interface{} {
	return map[string]interface{}{"app1": "hips1"}
}
