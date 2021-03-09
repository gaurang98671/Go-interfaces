package main

type game struct {
	id   int
	name string
}

func (g *game) print() string {
	return g.name
}
