package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", http.StripPrefix("/", fs))
	fmt.Println("http://localhost:1234")
	log.Fatal(http.ListenAndServe("0.0.0.0:1234", nil))
}
