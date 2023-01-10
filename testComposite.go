package main

import (
	"design-pattern-go/composite"
	"fmt"
)

func main() {
	f1 := composite.File{Name: "file1"}
	f2 := composite.File{Name: "file2"}
	f3 := composite.File{Name: "file3"}
	f4 := composite.File{Name: "file4"}
	f5 := composite.File{Name: "file5"}
	d1 := composite.Folder{
		Name:      "folder1",
		Component: []composite.Component{&f1, &f2, &f3},
	}
	d2 := composite.Folder{
		Name:      "folder2",
		Component: []composite.Component{&f4, &f5, &d1},
	}
	d3 := composite.Folder{
		Name:      "folder3",
		Component: []composite.Component{&d1, &d2, &f1, &f3, &f5},
	}
	d3.Search("file2")
	d3.Add(&f2)
	d3.Add(&f4)
	fmt.Println("----------------")
	d3.Search("file2")
}
