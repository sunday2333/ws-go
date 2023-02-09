package main

import (
	"log"
	"net/http"
)

func main() {
	setup()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setup() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

}
