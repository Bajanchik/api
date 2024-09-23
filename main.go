package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloHandler).Methods("GET")
	http.ListenAndServe(":8090", router)
}
