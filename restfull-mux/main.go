package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	http.Handle("/", router)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

type usuario struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var usuarios = []usuario{
	{
		Name: "Rafael",
		Age:  32,
	},
	{
		Name: "Erick",
		Age:  29,
	},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(usuarios)
}
