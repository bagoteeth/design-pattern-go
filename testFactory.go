package main

import (
	"design-pattern-go/factory"
	"fmt"
)

func main() {
	host := factory.ClassifyNode(factory.AssetNode{
		NodeType:  factory.TYPE_HOST,
		Cluster:   "hostCluster1",
		HostGroup: "hostHG1",
		Host:      "host1",
	})
	app := factory.ClassifyNode(factory.AssetNode{
		NodeType:  factory.TYPE_APP,
		Cluster:   "appCluster1",
		Namespace: "appNS1",
		App:       "app1",
	})
	container := factory.ClassifyNode(factory.AssetNode{
		NodeType:  factory.TYPE_CONTAINER,
		Cluster:   "cCluster1",
		HostGroup: "cHG1",
		Host:      "cHost1",
		Container: "ccc",
	})
	fmt.Printf("host: \n%+v\n%+v\n%+v\n", host.GetRuntimeEvent(), host.GetCnnfEvent(), host.GetHipsEvent())
	fmt.Printf("app: \n%+v\n%+v\n%+v\n", app.GetRuntimeEvent(), app.GetCnnfEvent(), app.GetHipsEvent())
	fmt.Printf("container: \n%+v\n%+v\n%+v\n", container.GetRuntimeEvent(), container.GetCnnfEvent(), container.GetHipsEvent())
}
