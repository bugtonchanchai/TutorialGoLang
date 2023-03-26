package main

import "fmt"

type Game struct {
	Title    string
	Category string
	Platform string
}

type Movie struct {
	Title    string
	Category string
}

type GameOrMovie interface {
	GetValue()
}

func (g *Game) GetValue() {
	fmt.Println(g.Title, g.Category, g.Platform)
}

func (m *Movie) GetValue() {
	fmt.Println(m.Title, m.Category)
}

func GetValueGameOrMovie[T GameOrMovie](t T) {
	t.GetValue()
}

func main() {
	//--- tutorial 1
	g := &Game{
		Title:    "Game 1",
		Category: "Category game 1",
		Platform: "Platform game 1",
	}

	f := &Movie{
		Title:    "Movie 1",
		Category: "Category movie 1",
	}

	g.GetValue()
	f.GetValue()

	GetValueGameOrMovie(g)
	GetValueGameOrMovie(f)

	//--- tutorial 2
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
