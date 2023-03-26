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
}
