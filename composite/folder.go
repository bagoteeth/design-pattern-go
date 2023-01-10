package composite

import "fmt"

type Folder struct {
	Name      string
	Component []Component
}

func (r *Folder) Search(s string) {
	fmt.Printf("Searching for keyword %s in folder %s: %t\n", s, r.Name, s == r.Name)
	for _, v := range r.Component {
		v.Search(s)
	}
}

func (r *Folder) Add(c Component) {
	r.Component = append(r.Component, c)
}
