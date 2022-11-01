package builder

type Product struct {
	Height   int
	Weight   int
	Material string
}

type BuilderIntf interface {
	Step1()
	Step2()
	Step3()
	GetProduct() Product
}
