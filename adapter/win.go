package adapter

import "fmt"

type Winxp struct {
	Name string
}

func (r *Winxp) OldWork() {
	fmt.Printf("winxp %s is doing old work\n", r.Name)
}

type Win10 struct {
	Name string
}

func (r *Win10) NewWork() {
	fmt.Printf("win10 %s is doing new work\n", r.Name)
}
