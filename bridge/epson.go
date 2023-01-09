package bridge

import "fmt"

type Epson struct {
	file string
}

func NewEpson(s string) Epson {
	return Epson{file: s}
}

func (r *Epson) PrintFile() {
	fmt.Printf("Epson print: %s\n", r.file)
}
