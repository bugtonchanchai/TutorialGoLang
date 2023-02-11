package pg

import (
	"fmt"
	"math"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

var powd = []int{1, 2, 4, 8, 16, 32, 64, 128}

type Vt struct {
	Lat, Long float64
}

var ms map[string]Vt

var mg = map[string]Vt{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func Pg03(isActive bool) {
	if !isActive {
		return
	}

	fmt.Println("Welcome to playground 03")

	i, j := 42, 2701
	p := &i
	fmt.Println(*p) // 42
	*p = 21
	fmt.Println(i) // 21
	p = &j
	*p = *p / 37
	fmt.Println(j) // 72
	fmt.Println(p) // 0xc0000180d0

	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	v = Vertex{1, 2}
	pa := &v
	pa.X = 1e9
	fmt.Println(v)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var sl []int = primes[1:4]
	fmt.Println(sl)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	na := names[0:2]
	nb := names[1:3]
	fmt.Println(na, nb)
	nb[0] = "XXX"
	fmt.Println(na, nb)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	afs := []int{2, 3, 5, 7, 11, 13}
	afs = afs[1:4]
	fmt.Println(afs)
	afs = afs[:2]
	fmt.Println(afs)
	afs = afs[1:]
	fmt.Println(afs)

	gs := []int{2, 3, 5, 7, 11, 13}
	printSlice(gs)
	gs = gs[:0]
	printSlice(gs)
	gs = gs[:4]
	printSlice(gs)
	gs = gs[2:]
	printSlice(gs)

	var fsa []int
	fmt.Println(fsa, len(fsa), cap(fsa))
	if fsa == nil {
		fmt.Println("nil!")
	}

	gaa := make([]int, 5)
	printSlice2("gaa", gaa)
	gab := make([]int, 0, 5)
	printSlice2("gab", gab)
	gac := gab[:2]
	printSlice2("gac", gac)
	gad := gac[2:5]
	printSlice2("gad", gad)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var ys []int
	printSlice(ys)
	ys = append(ys, 0)
	printSlice(ys)
	ys = append(ys, 1)
	printSlice(ys)
	ys = append(ys, 2, 3, 4, 5, 7)
	printSlice(ys)

	for i, v := range powd {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	powz := make([]int, 10)
	for i := range powz {
		powz[i] = 1 << uint(i)
	}
	for _, value := range powz {
		fmt.Printf("%d\n", value)
	}

	ms = make(map[string]Vt)
	ms["Bell Labs"] = Vt{
		40.68433, -74.39967,
	}
	fmt.Println(ms["Bell Labs"])

	fmt.Println(mg)

	mk := make(map[string]int)
	mk["Answer"] = 42
	fmt.Println("The value:", mk["Answer"])
	mk["Answer"] = 48
	fmt.Println("The value:", mk["Answer"])
	delete(mk, "Answer")
	fmt.Println("The value:", mk["Answer"])
	va, ok := mk["Answer"]
	fmt.Println("The value:", va, "Present?", ok)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
