package main

import (
	"fmt"
)

func main() {
	dispersa := &Matriz{nil, nil}
	persona := &NodoPersona{ESTE: nil, OESTE: nil, SUR: nil, NORTE: nil, Nombre: "Marvin", Edad: 24}
	persona2 := &NodoPersona{ESTE: nil, OESTE: nil, SUR: nil, NORTE: nil, Nombre: "Xarvin", Edad: 25}
	persona3 := &NodoPersona{ESTE: nil, OESTE: nil, SUR: nil, NORTE: nil, Nombre: "Alejandro", Edad: 10}
	persona4 := &NodoPersona{ESTE: nil, OESTE: nil, SUR: nil, NORTE: nil, Nombre: "Yaiza", Edad: 10}
	persona5 := &NodoPersona{ESTE: nil, OESTE: nil, SUR: nil, NORTE: nil, Nombre: "Alan", Edad: 10}
	persona6 := &NodoPersona{ESTE: nil, OESTE: nil, SUR: nil, NORTE: nil, Nombre: "Alan", Edad: 2}
	dispersa.Add(persona)
	dispersa.Add(persona2)
	dispersa.Add(persona3)
	dispersa.Add(persona4)
	dispersa.Add(persona5)
	dispersa.Add(persona6)
	dispersa.Imprimir()
	fmt.Println("++++++++")
	dispersa.Imprimir2()
	fmt.Println("A"[0])
}
