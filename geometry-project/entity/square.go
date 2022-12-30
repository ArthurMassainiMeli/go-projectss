package entity

type Square struct {
	Side float64
}

func (s Square) GetArea() float64 {
	return s.Side * s.Side
}
