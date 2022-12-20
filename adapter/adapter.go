package adapter

import "fmt"

type WinAdapter struct {
	Old Winxp
	New Win10
}

func (r *WinAdapter) OldWork() {
	fmt.Printf("win10 %s is running on winxp: ", r.New.Name)
	r.New.NewWork()
}

func (r *WinAdapter) NewWork() {
	fmt.Printf("winxp %s is running on win10: ", r.Old.Name)
	r.Old.OldWork()
}
