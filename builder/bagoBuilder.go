package builder

type BagoBuilder struct {
	Product
}

func NewBagoBuilder() BuilderIntf {
	return &BagoBuilder{}
}

func (r *BagoBuilder) Step1() {
	r.Height = 100
}

func (r *BagoBuilder) Step2() {
	r.Weight = 100
}

func (r *BagoBuilder) Step3() {
	r.Material = "gold"
}

func (r *BagoBuilder) GetProduct() Product {
	return Product{
		Height:   r.Height,
		Weight:   r.Weight,
		Material: r.Material,
	}
}
