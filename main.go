package main

import (
	"fmt"
)

func main() {
	Go := Course{
		Name:    "Go desde Cero",
		Price:   1.34,
		IsFree:  false,
		UserIDs: []uint{1, 2, 4, 5},
		Classes: map[uint]string{
			1: "Introduction",
			2: "Another",
			3: "Another one",
		},
	}
	Css := Course{
		Name:   "Css desde cero",
		IsFree: true,
	}
	js := Course{}
	js.Name = "Js desde Cero"

	fmt.Println(Go.Name)
	fmt.Printf("%+v\n", Css)
	fmt.Printf("%+v\n", js)

	Go.ImprimirClasses()

	Go.ChangePrice(10.5)
	fmt.Println(Go.Price)
}
