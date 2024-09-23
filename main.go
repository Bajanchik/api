package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var message string

type requestBody struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello "+message)
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {

	var req requestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Успешно")
	message = req.Message
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/message", MessageHandler).Methods("POST")
	http.ListenAndServe(":8090", router)
}
