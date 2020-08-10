package main

type shape interface {
	area() float32
}

type square struct {
	side float32
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return 3.1415 * c.radius * c.radius
}

func (s square) area() float32 {
	return s.side * s.side
}
