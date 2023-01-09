package bridge

import "fmt"

type Mac struct {
	printer Printer
}

func (r *Mac) SetPrinter(p Printer) {
	r.printer = p
}

func (r *Mac) Print() {
	fmt.Printf("Using Mac: ")
	r.printer.PrintFile()
}
