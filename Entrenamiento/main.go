package main

import (
	"fmt"
)

type Animal struct {
	nombre string
	color  string
}

func (this Animal) saludar() {
	fmt.Println("hola")
}

type Pez struct {
	Animal
	aletas int
}

func main() {
	pez := Pez{Animal: Animal{"Filemon", "rojo"}, aletas: 5}
	fmt.Println(pez)
	pez.saludar()
}
