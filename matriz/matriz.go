package main

import (
	"fmt"
	"reflect"
)

type NodoPersona struct {
	ESTE, OESTE, NORTE, SUR interface{}
	Nombre                  string
	Edad                    int
}

type NodoCabeceraVertical struct {
	ESTE  interface{}
	NORTE interface{}
	SUR   interface{}
	OESTE interface{}
	Edad  int
}

type NodoCabeceraHorizontal struct {
	ESTE, OESTE, SUR, NORTE interface{}
	Letra                   string
}

type Matriz struct {
	CabH *NodoCabeceraHorizontal
	CabV *NodoCabeceraVertical
}

func (this *Matriz) getHorizontal(nombre string) interface{} {
	if this.CabH == nil {
		return nil
	}
	var aux interface{} = this.CabH
	for aux != nil {
		if aux.(*NodoCabeceraHorizontal).Letra == string(nombre[0]) {
			return aux
		}
		aux = aux.(*NodoCabeceraHorizontal).ESTE
	}
	return nil
}

func (this *Matriz) getVertical(edad int) interface{} {
	if this.CabV == nil {
		return nil
	}
	var aux interface{} = this.CabV
	for aux != nil {
		if aux.(*NodoCabeceraVertical).Edad == edad {
			return aux
		}
		aux = aux.(*NodoCabeceraVertical).SUR
	}
	return nil
}

func (this *Matriz) crearVertical(edad int) *NodoCabeceraVertical {
	if this.CabV == nil {
		nueva := &NodoCabeceraVertical{
			ESTE:  nil,
			OESTE: nil,
			SUR:   nil,
			NORTE: nil,
			Edad:  edad,
		}
		this.CabV = nueva
		return nueva
	}
	var aux interface{} = this.CabV
	if edad < aux.(*NodoCabeceraVertical).Edad {
		nueva := &NodoCabeceraVertical{
			ESTE:  nil,
			OESTE: nil,
			SUR:   nil,
			NORTE: nil,
			Edad:  edad,
		}
		nueva.SUR = this.CabV
		this.CabV.NORTE = nueva
		this.CabV = nueva
		return nueva
	}
	for aux.(*NodoCabeceraVertical).SUR != nil {
		if edad > aux.(*NodoCabeceraVertical).Edad && edad < aux.(*NodoCabeceraVertical).SUR.(*NodoCabeceraVertical).Edad {
			nueva := &NodoCabeceraVertical{
				ESTE:  nil,
				OESTE: nil,
				SUR:   nil,
				NORTE: nil,
				Edad:  edad,
			}
			tmp := aux.(*NodoCabeceraVertical).SUR
			tmp.(*NodoCabeceraVertical).NORTE = nueva
			nueva.SUR = tmp
			aux.(*NodoCabeceraVertical).SUR = nueva
			nueva.NORTE = aux
			return nueva
		}
		aux = aux.(*NodoCabeceraVertical).SUR
	}
	nueva := &NodoCabeceraVertical{
		ESTE:  nil,
		OESTE: nil,
		SUR:   nil,
		NORTE: nil,
		Edad:  edad,
	}
	aux.(*NodoCabeceraVertical).SUR = nueva
	nueva.NORTE = aux
	return nueva
}

func (this *Matriz) crearHorizontal(nombre string) *NodoCabeceraHorizontal {
	if this.CabH == nil {
		nueva := &NodoCabeceraHorizontal{
			ESTE:  nil,
			OESTE: nil,
			SUR:   nil,
			NORTE: nil,
			Letra: string(nombre[0]),
		}
		this.CabH = nueva
		return nueva
	}
	var aux interface{} = this.CabH
	if nombre <= aux.(*NodoCabeceraHorizontal).Letra {
		nueva := &NodoCabeceraHorizontal{
			ESTE:  nil,
			OESTE: nil,
			SUR:   nil,
			NORTE: nil,
			Letra: string(nombre[0]),
		}
		nueva.ESTE = this.CabH
		this.CabH.OESTE = nueva
		this.CabH = nueva
		return nueva
	}
	for aux.(*NodoCabeceraHorizontal).ESTE != nil {
		if string(nombre[0]) > aux.(*NodoCabeceraHorizontal).Letra && string(nombre[0]) <= aux.(*NodoCabeceraHorizontal).ESTE.(*NodoCabeceraHorizontal).Letra {
			nueva := &NodoCabeceraHorizontal{
				ESTE:  nil,
				OESTE: nil,
				SUR:   nil,
				NORTE: nil,
				Letra: string(nombre[0]),
			}
			tmp := aux.(*NodoCabeceraHorizontal).ESTE
			tmp.(*NodoCabeceraHorizontal).OESTE = nueva
			nueva.ESTE = tmp
			aux.(*NodoCabeceraHorizontal).ESTE = nueva
			nueva.OESTE = aux
			return nueva
		}
		aux = aux.(*NodoCabeceraHorizontal).ESTE
	}
	nueva := &NodoCabeceraHorizontal{
		ESTE:  nil,
		OESTE: nil,
		SUR:   nil,
		NORTE: nil,
		Letra: string(nombre[0]),
	}
	aux.(*NodoCabeceraHorizontal).ESTE = nueva
	nueva.OESTE = aux
	return nueva
}

func (this *Matriz) obtenerUltimoV(cabecera *NodoCabeceraHorizontal, edad int) interface{} {
	if cabecera.SUR == nil {
		return cabecera
	}
	aux := cabecera.SUR
	if edad <= aux.(*NodoPersona).Edad {
		return cabecera
	}
	for aux.(*NodoPersona).SUR != nil {
		if edad > aux.(*NodoPersona).Edad && edad <= aux.(*NodoPersona).SUR.(*NodoPersona).Edad {
			return aux
		}
		aux = aux.(*NodoPersona).SUR

	}
	if edad <= aux.(*NodoPersona).Edad {
		return aux.(*NodoPersona).NORTE
	}
	return aux
}
func (this *Matriz) obtenerUltimoH(cabecera *NodoCabeceraVertical, nombre string) interface{} {
	if cabecera.ESTE == nil {
		return cabecera
	}
	aux := cabecera.ESTE
	if nombre <= aux.(*NodoPersona).Nombre {
		return cabecera
	}
	for aux.(*NodoPersona).ESTE != nil {
		if nombre > aux.(*NodoPersona).Nombre && nombre <= aux.(*NodoPersona).ESTE.(*NodoPersona).Nombre {
			return aux
		}
		aux = aux.(*NodoPersona).ESTE
	}
	//tx:=aux.(*NodoPersona)
	if nombre <= aux.(*NodoPersona).Nombre {
		return aux.(*NodoPersona).OESTE
	}
	return aux
}

func (this *Matriz) Add(nueva *NodoPersona) {
	vertical := this.getVertical(nueva.Edad)
	horizontal := this.getHorizontal(nueva.Nombre)
	if vertical == nil {
		vertical = this.crearVertical(nueva.Edad)
	}
	if horizontal == nil {
		horizontal = this.crearHorizontal(nueva.Nombre)
	}
	izquierda := this.obtenerUltimoH(vertical.(*NodoCabeceraVertical), nueva.Nombre)
	superior := this.obtenerUltimoV(horizontal.(*NodoCabeceraHorizontal), nueva.Edad)
	if reflect.TypeOf(izquierda).String() == "*main.NodoPersona" {
		if izquierda.(*NodoPersona).ESTE == nil {
			izquierda.(*NodoPersona).ESTE = nueva
			nueva.OESTE = izquierda
		} else {
			tmp := izquierda.(*NodoPersona).ESTE
			izquierda.(*NodoPersona).ESTE = nueva
			nueva.OESTE = izquierda
			tmp.(*NodoPersona).OESTE = nueva
			nueva.ESTE = tmp
		}
	} else {
		if izquierda.(*NodoCabeceraVertical).ESTE == nil {
			izquierda.(*NodoCabeceraVertical).ESTE = nueva
			nueva.OESTE = izquierda
		} else {
			tmp := izquierda.(*NodoCabeceraVertical).ESTE
			izquierda.(*NodoCabeceraVertical).ESTE = nueva
			nueva.OESTE = izquierda
			tmp.(*NodoPersona).OESTE = nueva
			nueva.ESTE = tmp
		}
	}

	/*SUperior*/
	if reflect.TypeOf(superior).String() == "*main.NodoPersona" {
		if superior.(*NodoPersona).SUR == nil {
			superior.(*NodoPersona).SUR = nueva
			nueva.NORTE = superior
		} else {
			tmp := superior.(*NodoPersona).SUR
			superior.(*NodoPersona).SUR = nueva
			nueva.NORTE = superior
			tmp.(*NodoPersona).NORTE = nueva
			nueva.SUR = tmp
		}
	} else {
		if superior.(*NodoCabeceraHorizontal).SUR == nil {
			superior.(*NodoCabeceraHorizontal).SUR = nueva
			nueva.NORTE = superior
		} else {
			tmp := superior.(*NodoCabeceraHorizontal).SUR
			superior.(*NodoCabeceraHorizontal).SUR = nueva
			nueva.NORTE = superior
			tmp.(*NodoPersona).NORTE = nueva
			nueva.SUR = tmp
		}
	}
}

func (this *Matriz) Imprimir() {
	var aux interface{} = this.CabV
	for aux != nil {
		fmt.Print(aux.(*NodoCabeceraVertical).Edad, "***************")
		tmp := aux.(*NodoCabeceraVertical).ESTE
		for tmp != nil {
			fmt.Printf("%v,%v------", tmp.(*NodoPersona).Nombre, tmp.(*NodoPersona).Edad)
			tmp = tmp.(*NodoPersona).ESTE
		}
		fmt.Print("\n")
		aux = aux.(*NodoCabeceraVertical).SUR
	}
}

func (this *Matriz) Imprimir2() {
	var aux interface{} = this.CabH
	for aux != nil {
		fmt.Print(aux.(*NodoCabeceraHorizontal).Letra, "*****************")
		tmp := aux.(*NodoCabeceraHorizontal).SUR
		for tmp != nil {
			fmt.Printf("%v,%v-------", tmp.(*NodoPersona).Nombre, tmp.(*NodoPersona).Edad)
			tmp = tmp.(*NodoPersona).SUR
		}
		fmt.Println("")
		aux = aux.(*NodoCabeceraHorizontal).ESTE
	}
}
