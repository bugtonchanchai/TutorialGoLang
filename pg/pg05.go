package pg

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func Pg05(isActive bool) {
	if !isActive {
		return
	}
	fmt.Println("Welcome to playground 05 (interface)")

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	// a = v

	fmt.Println(a.Abs())

	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

type MyFloat float64

type F float64

type Vertex struct {
	X, Y float64
}

type I interface {
	M()
}

type T struct {
	S string
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)

	}
	return float64(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}
