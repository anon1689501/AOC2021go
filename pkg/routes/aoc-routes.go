package routes

import (
	"fmt"
	"net/http"

	"github.com/anon1689501/AOC2021go/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAocRoutes = func(router *mux.Router) {
	router.HandleFunc("/day/", controllers.Day1).Methods("POST")
	router.HandleFunc("/day/", controllers.GetForm).Methods("GET")
	// router.HandleFunc("/day/{id}")
	router.Handle("/", http.FileServer(http.Dir("./static/")))
	fmt.Printf("starting server at port 80")
}
