package main

import (
	"bookstore/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	fmt.Println("http listening...")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8808", r))

}
