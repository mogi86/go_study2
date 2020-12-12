package construction

type Vertex1 struct {
	X, Y int
}

func (v *Vertex1) SetX(a int) {
	v.X = a
}

func (v *Vertex1) Calc() int {
	return v.X * v.Y
}
