package main

import (
	"design-pattern-go/adapter"
	"fmt"
)

func main() {
	oldC := adapter.OldClient{}
	newC := adapter.NewClient{}
	win10 := adapter.Win10{Name: "newteeth"}
	winxp := adapter.Winxp{Name: "oldbago"}

	winAdapter := adapter.WinAdapter{
		Old: winxp,
		New: win10,
	}
	fmt.Println("*****old client*****")
	oldC.DoOldWork(&winxp)
	oldC.DoOldWork(&winAdapter)
	fmt.Println("*****new client*****")
	newC.DoNewWork(&win10)
	//newC.DoNewWork(&winxp) illegal
	newC.DoNewWork(&winAdapter)
}
