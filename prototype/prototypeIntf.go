package prototype

type PrototypeIntf interface {
	Clone() PrototypeIntf
	Print(s string)
}
