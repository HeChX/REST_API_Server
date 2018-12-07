package service

import (
	"fmt"
	"net/http"
	_ "unsafe"

	"github.com/gorilla/mux"

	"github.com/Howlyao/Server/database"
)

func queryPeople(w http.ResponseWriter, r *http.Request) {
	myDB := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]
	// fmt.Println(myDB.DB_v)
	// people := model.Peoples{
	// 	model.People{Name: "Howl", Height: "177", Mass: "55"},
	// }

	// if err := json.NewEncoder(w).Encode(people); err != nil {
	// 	panic(err)
	// }
	fmt.Fprintln(w, myDB.QueryPeople(id))

}

func queryPlanet(w http.ResponseWriter, r *http.Request) {
	myDB := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, myDB.QueryPlanet(id))
}

func queryFilm(w http.ResponseWriter, r *http.Request) {
	myDB := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, myDB.QueryFilm(id))
}
