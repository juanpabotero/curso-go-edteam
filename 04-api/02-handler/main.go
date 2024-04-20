package main

import (
	"fmt"
	"net/http"
)

func main () {
	// registrar ruta y handler
	http.HandleFunc("/greeting", greetingHandler)
	// levantar servidor
	http.ListenAndServe(":8080", nil)
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))
	fmt.Fprintf(w, "Hello, World!")
}