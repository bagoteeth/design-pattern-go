package builder

type TeethBuilder struct {
	Product
}

func NewTeethBuilder() BuilderIntf {
	return &TeethBuilder{}
}

func (r *TeethBuilder) Step1() {
	r.Height = 1000
}

func (r *TeethBuilder) Step2() {
	r.Weight = 1000
}

func (r *TeethBuilder) Step3() {
	r.Material = "wood"
}

func (r *TeethBuilder) GetProduct() Product {
	return Product{
		Height:   r.Height,
		Weight:   r.Weight,
		Material: r.Material,
	}
}
