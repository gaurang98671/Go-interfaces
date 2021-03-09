package main

type Books struct {
	id    int
	title string
}

func (b *Books) print() string {
	return b.title
}
