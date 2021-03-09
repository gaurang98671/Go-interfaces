package main

import "fmt"

type printer interface {
	print() string
}

type list []printer

func main() {

	var (
		a = game{1, "asd"}
		b = Books{2, "book1"}
	)

	var store list

	store = append(store, &a)
	store = append(store, &b)

	for _, a := range store {
		fmt.Println(a)
	}

}
