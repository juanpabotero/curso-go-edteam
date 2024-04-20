package handler

import (
	"net/http"
)

// RoutePerson maneja las rutas de la persona
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", h.create)
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/delete", h.delete)
	mux.HandleFunc("/v1/persons/get-all", h.getAll)
	mux.HandleFunc("/v1/persons/get-by-id", h.getByID)
}
