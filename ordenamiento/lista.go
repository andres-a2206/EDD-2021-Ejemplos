package main

import (
	"fmt"
	"os"
	"strings"
)

type Nodo struct {
	Nombre    string
	edad      int
	siguiente *Nodo
}

type Lista struct {
	Cabeza *Nodo
	Cola   *Nodo
}

func nuevaLista() *Lista {
	return &Lista{nil, nil}
}
func (this *Lista) insertar(n *Nodo) {
	if this.Cabeza == nil {
		this.Cabeza = n
		this.Cola = n
	} else {
		this.Cola.siguiente = n
		this.Cola = n
	}
}

func (this *Lista) Graficar() string {
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=record];\n")
	graficar(this.Cabeza, &cadena, nil)
	fmt.Fprintf(&cadena, "}")
	guardarArchivo(cadena.String())
	return cadena.String()
}
func graficar(anterior *Nodo, s *strings.Builder, actual *Nodo) {
	if anterior != nil {
		fmt.Fprintf(s, "node%p[label=\"%v|%v\"];\n", &(*anterior), anterior.Nombre, anterior.edad)
		if actual != nil {
			fmt.Fprintf(s, "node%p->node%p;\n", &(*actual), &(*anterior))
			fmt.Fprintf(s, "node%p->node%p;\n", &(*anterior), &(*actual))
		}
		graficar(anterior.siguiente, s, anterior)
	}
}
func guardarArchivo(cadena string) {
	f, err := os.Create("lista.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(cadena)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	l := nuevaLista()
	n := Nodo{"Marvin", 24, nil}
	a := Nodo{"Marvin", 24, nil}
	b := Nodo{"Marvin", 24, nil}
	l.insertar(&n)
	l.insertar(&a)
	l.insertar(&b)
	fmt.Println(l.Graficar())
}
