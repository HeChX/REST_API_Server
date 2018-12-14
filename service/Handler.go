package service

import (
	"fmt"
	"net/http"
	_ "unsafe"

	"github.com/gorilla/mux"

	"github.com/Howlyao/REST_API_Server/database"
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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, myDB.QueryPeople(id))

}

func queryPlanet(w http.ResponseWriter, r *http.Request) {
	myDB := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, myDB.QueryPlanet(id))
}

func queryFilm(w http.ResponseWriter, r *http.Request) {
	myDB := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, myDB.QueryFilm(id))
}

func querySpecies(w http.ResponseWriter, r *http.Request) {
	myDB := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, myDB.QuerySpecies(id))
}

func queryStarship(w http.ResponseWriter, r *http.Request) {
	myDB := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, myDB.QueryStarship(id))
}

func queryVehicle(w http.ResponseWriter, r *http.Request) {
	myDB := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, myDB.QueryVehicle(id))
}
