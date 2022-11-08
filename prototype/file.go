package prototype

import "fmt"

type File struct {
	Name string
}

func (r *File) Print(s string) {
	fmt.Println(s + r.Name)
}

func (r *File) Clone() PrototypeIntf {
	return &File{Name: r.Name + "_clone"}
}
