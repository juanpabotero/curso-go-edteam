package main

import (
	"curso-go-edteam/04-api/03-api/handler"
	"curso-go-edteam/04-api/03-api/storage"
	"log"
	"net/http"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)

	log.Println("El servidor está corriendo en http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}