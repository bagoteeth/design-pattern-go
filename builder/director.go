package builder

type Director struct {
	Builder BuilderIntf
}

func NewDirector(b BuilderIntf) *Director {
	return &Director{Builder: b}
}

func (d *Director) SetDirector(b BuilderIntf) {
	d.Builder = b
}

func (d *Director) BuildProduct() Product {
	d.Builder.Step1()
	d.Builder.Step2()
	d.Builder.Step3()
	return d.Builder.GetProduct()
}
