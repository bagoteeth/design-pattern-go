package prototype

import "fmt"

type Folder struct {
	Children []PrototypeIntf
	Name     string
}

func (r *Folder) Print(s string) {
	fmt.Println(s + r.Name)
	for _, v := range r.Children {
		v.Print(s + s)
	}
}

func (r *Folder) Clone() PrototypeIntf {
	clone := &Folder{
		Name: r.Name + "_clone",
	}
	tmpChildren := make([]PrototypeIntf, 0)
	for _, v := range r.Children {
		copy1 := v.Clone()
		tmpChildren = append(tmpChildren, copy1)
	}
	clone.Children = tmpChildren
	return clone
}
