package main

import (
	"design-pattern-go/builder"
	"fmt"
)

func main() {
	director := builder.NewDirector(builder.NewBagoBuilder())
	fmt.Println(director.BuildProduct())
	director.SetDirector(builder.NewTeethBuilder())
	fmt.Println(director.BuildProduct())
}
