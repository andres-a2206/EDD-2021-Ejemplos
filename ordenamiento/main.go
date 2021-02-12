package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(100))
	}
	fmt.Println(a)
	/*fmt.Println("Burbuja")
	burbuja(a)*/
	fmt.Println("Insercion")
	insercion(a)
	fmt.Println(busqueda(&a, 87))
}

func burbuja(arreglo []int) {
	var i, j, aux int
	for i = 0; i < len(arreglo)-1; i++ {
		for j = 0; j < len(arreglo)-i-1; j++ {
			if arreglo[j+1] < arreglo[j] {
				aux = arreglo[j+1]
				arreglo[j+1] = arreglo[j]
				arreglo[j] = aux
			}
		}
	}
	fmt.Println(arreglo)
}

func insercion(arreglo []int) {
	var p, j int
	var aux int
	for p = 1; p < len(arreglo); p++ {
		aux = arreglo[p]
		j = p - 1
		for (j >= 0) && (aux < arreglo[j]) {
			arreglo[j+1] = arreglo[j]
			j--
		}
		arreglo[j+1] = aux
	}
	fmt.Println(arreglo)
}

func Quicksort(arreglo []int) {
	quicksort1(&arreglo, 0, len(arreglo)-1)
	fmt.Println(arreglo)
}
func quicksort1(arreglo *[]int, start int, end int) {
	part := partition(arreglo, start, end)
	if (part - 1) > start {
		quicksort1(arreglo, start, part-1)
	}
	if (part + 1) < end {
		quicksort1(arreglo, part+1, end)
	}
}

func partition(arreglo *[]int, start int, end int) int {
	pivote := (*arreglo)[end]
	for i := start; i < end; i++ {
		if (*arreglo)[i] < pivote {
			tmp := (*arreglo)[start]
			(*arreglo)[start] = (*arreglo)[i]
			(*arreglo)[i] = tmp
			start++
		}
	}
	tmp := (*arreglo)[start]
	(*arreglo)[start] = pivote
	(*arreglo)[end] = tmp
	return start
}

func busquedaB(arreglo *[]int, dato int) bool {
	return BusquedaBinaria(arreglo, dato, 0, len(*arreglo)-1)
}
func BusquedaBinaria(arreglo *[]int, dato int, inicio int, final int) bool {
	medio := (final + inicio) / 2
	if final >= 0 && final < len(*arreglo) && inicio < final && inicio >= 0 {
		if (*arreglo)[medio] == dato {
			return true
		} else if dato < (*arreglo)[medio] {
			return BusquedaBinaria(arreglo, dato, inicio, medio-1)
		}
		return BusquedaBinaria(arreglo, dato, medio+1, final)
	}
	return false
}
