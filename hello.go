package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rsc.io/quote"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	return
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	id := 0
	fmt.Println("test : ", id)

	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/test/{id}", TestHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
