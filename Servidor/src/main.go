package main

//go get github.com/githubnemo/CompileDaemon
/*
{
	Nombre,
	Edad
}
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Mensaje struct {
	Nombre string `json:Nombre`
	Edad   int    `json:Edad`
}

type Mensajito struct {
	Fecha string `json:Fecha`
	Texto string `json:Texto`
}

type MensajeDestino struct {
	Origen  string      `json:Origen`
	Destino string      `json:Destino`
	Msg     []Mensajito `json:Msg`
}

type MensajeGlobal struct {
	Mensajes []MensajeDestino `json:Mensajes`
}

func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Funcionan EDD")
}

func agregar(w http.ResponseWriter, r *http.Request) {
	var ms MensajeGlobal
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &ms)
	json.NewEncoder(w).Encode(ms)
	fmt.Println(ms.Mensajes[0].Destino)
}
func numero(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b, _ := strconv.Atoi(vars["id"])
	a := Mensaje{"El numero que me mandaste es ", b}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(a)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/agregar", agregar).Methods("POST")
	router.HandleFunc("/numero/{id}", numero).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
