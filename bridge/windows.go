package bridge

import "fmt"

type Windows struct {
	printer Printer
}

func (r *Windows) SetPrinter(p Printer) {
	r.printer = p
}

func (r *Windows) Print() {
	fmt.Printf("Using windows: ")
	r.printer.PrintFile()
}
