package composite

import "fmt"

type File struct {
	Name string
}

func (r *File) Search(s string) {
	fmt.Printf("Searching for keyword %s in file %s: %t\n", s, r.Name, s == r.Name)
}
