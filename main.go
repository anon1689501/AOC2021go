package main

import (
	"log"
	"net/http"

	"github.com/anon1689501/AOC2021go/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterAocRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":80", r))
}
