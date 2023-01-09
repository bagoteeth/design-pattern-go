package main

import "design-pattern-go/bridge"

func main() {
	mac := bridge.Mac{}
	windows := bridge.Windows{}
	epson := bridge.NewEpson("this is epson")
	hp := bridge.NewHp("this is hp")
	mac.SetPrinter(&epson)
	mac.Print()
	mac.SetPrinter(&hp)
	mac.Print()

	windows.SetPrinter(&epson)
	windows.Print()
	windows.SetPrinter(&hp)
	windows.Print()
}
