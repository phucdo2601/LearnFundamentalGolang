package main

import "fmt"

type car struct {
	make   string
	model  string
	height int
	width  int

	wheel struct {
		radius   int
		material string
	}
}

func main() {
	newCarA := car{
		make:   "titanium",
		model:  "s23",
		height: 234,
		width:  353,
		wheel: struct {
			radius   int
			material string
		}{
			radius:   231,
			material: "titan",
		},
	}

	fmt.Println("New car A: ", newCarA)
}
