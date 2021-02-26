package main

func main() {
	a := NewArbol()
	for i := 0; i < 20; i++ {
		a.Insertar(i)
	}
	a.Generar()
}
