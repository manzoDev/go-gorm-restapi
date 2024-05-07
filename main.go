package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manzoDev/go-gorm-restapi/routes"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)

}
