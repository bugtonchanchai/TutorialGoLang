package pg

import (
	"fmt"
	"math"
)

type Vertex04 struct {
	X, Y float64
}

type MyFloat04 float64

func Pg04(isActive bool) {
	if !isActive {
		return
	}
	fmt.Println("Welcome to playground 04 (methods)")

	// methods
	v := Vertex04{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat04(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Pointer receivers
	v.Scale2(10)
	fmt.Println(v.Abs())

	// Method vs Function
	v.Scale2(2)
	ScaleFunc(&v, 10)

	p := &Vertex04{4, 3}
	p.Scale2(3)
	ScaleFunc(p, 8)
	fmt.Println(v, p)

	// Methods and pointer indirection
	v = Vertex04{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p = &Vertex04{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

	// Choosing a value or pointer receiver
	v2 := &Vertex04{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v2, v2.Abs2())
	v2.Scale2(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v2, v2.Abs2())
}

func (v Vertex04) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex04) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat04) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex04) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex04, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex04) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
