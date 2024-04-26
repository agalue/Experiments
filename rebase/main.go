package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	log.Println("Starting web server")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err.Error())
	}
}
