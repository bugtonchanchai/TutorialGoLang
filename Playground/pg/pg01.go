package pg

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	MxInt  int        = 1<<32 - 1
	Z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func Pg01(isActive bool) {
	if !isActive {
		return
	}

	fmt.Println("Welcome to playground 01")
	fmt.Println("The time is ", time.Now().Format("2006-01-01"))
	fmt.Println("My random number is ", rand.Intn(10))
	fmt.Printf("Sqrt of 24 is %g \n", math.Sqrt(24))
	fmt.Println("Pi is ", math.Pi)
	fmt.Println("3 + 2 is ", addNumber(3, 2))

	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	fmt.Println(split(20))

	var i int
	var java, python bool
	fmt.Println(i, java, python)

	var ia, ib int = 1, 2
	var ca, cb = true, "no!"
	fmt.Println(ia, ib, ca, cb)

	k := 3
	da, db, dc := true, false, "!yes"
	fmt.Println(k, da, db, dc)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", MxInt, MxInt)
	fmt.Printf("Type: %T Value: %v\n", Z, Z)

	var ji int
	var jf float64
	var jb bool
	var js string
	fmt.Printf("%v %v %v %q\n", ji, jf, jb, js)

	var ka, kb int = 3, 4
	var kf float64 = math.Sqrt(float64(ka*kb + kb*ka))
	var kz uint = uint(kf)
	fmt.Println(kf, kz)

	eo := 42.3 + 3i
	fmt.Printf("eo is of type %T\n", eo)

	const cz = "Thai"
	const cx = true
	fmt.Println(cz, cx)
}

func addNumber(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum / 10
	y = sum - 2
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
