package service

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{Name: "people", Method: "GET", Pattern: "/people/{id}", HandlerFunc: queryPeople},
	Route{Name: "planets", Method: "GET", Pattern: "/planets/{id}", HandlerFunc: queryPlanet},
	Route{Name: "films", Method: "GET", Pattern: "/films/{id}", HandlerFunc: queryFilm},
	Route{Name: "species", Method: "GET", Pattern: "/species/{id}", HandlerFunc: querySpecies},
	Route{Name: "starships", Method: "GET", Pattern: "/starships/{id}", HandlerFunc: queryStarship},
	Route{Name: "vehicles", Method: "GET", Pattern: "/vehicles/{id}", HandlerFunc: queryVehicle},
}
