package main

import (
	"design-pattern-go/prototype"
	"fmt"
)

func main() {
	f1 := &prototype.File{Name: "file1"}
	f2 := &prototype.File{Name: "file2"}
	f3 := &prototype.File{Name: "file3"}
	folder1 := &prototype.Folder{
		Children: []prototype.PrototypeIntf{f1, f2},
		Name:     "folder1",
	}
	folder2 := &prototype.Folder{
		Children: []prototype.PrototypeIntf{folder1, f2, f3},
		Name:     "folder2",
	}
	folder2.Print("	")
	fmt.Print("**********\n")
	folder3 := folder2.Clone()
	folder3.Print("	")
	fmt.Print("**********\n")
	folder4 := folder3.Clone()
	folder4.Print("----")
}
