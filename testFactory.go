package main

import (
	"design-pattern-go/factory"
	"fmt"
)

func main() {
	node1 := &factory.AssetNode{
		NodeType:  factory.TYPE_HOST,
		Cluster:   "hostCluster1",
		HostGroup: "hostHG1",
		Host:      "host1",
	}
	node2 := &factory.AssetNode{
		NodeType:  factory.TYPE_APP,
		Cluster:   "appCluster1",
		Namespace: "appNS1",
		App:       "app1",
	}
	node3 := factory.AssetNode{
		NodeType:  factory.TYPE_CONTAINER,
		Cluster:   "cCluster1",
		HostGroup: "cHG1",
		Host:      "cHost1",
		Container: "ccc",
	}
	fmt.Printf("host: \n%+v\n%+v\n%+v\n", node1.Transform().GetRuntimeEvent(), node1.Transform().GetCnnfEvent(), node1.Transform().GetHipsEvent())
	fmt.Printf("app: \n%+v\n%+v\n%+v\n", node2.Transform().GetRuntimeEvent(), node2.Transform().GetCnnfEvent(), node2.Transform().GetHipsEvent())
	fmt.Printf("container: \n%+v\n%+v\n%+v\n", node3.Transform().GetRuntimeEvent(), node3.Transform().GetCnnfEvent(), node3.Transform().GetHipsEvent())
}
