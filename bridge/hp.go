package bridge

import "fmt"

type Hp struct {
	file string
}

func NewHp(s string) Hp {
	return Hp{file: s}
}

func (r *Hp) PrintFile() {
	fmt.Printf("Hp print: %s\n", r.file)
}
