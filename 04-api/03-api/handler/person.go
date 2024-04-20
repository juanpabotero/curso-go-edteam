package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"curso-go-edteam/04-api/03-api/model"
)

// Storage interfaz para el almacenamiento de personas
type person struct {
	storage Storage
}

// newPerson crea una nueva instancia de person
func newPerson(storage Storage) person {
	return person{storage}
}

// create handler para crear una nueva persona
func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	// decodificamos el cuerpo de la petición y lo guardamos en la variable data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al crear la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Persona creada correctamente", nil)
	responseJSON(w, http.StatusCreated, response)
}

// update handler para actualizar una persona
func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || ID <= 0 {
		response := newResponse(Error, "El id debe ser un número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	// decodificamos el cuerpo de la petición y lo guardamos en la variable data
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(ID, &data)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Ocurrió un error al actualizar el registro", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Persona actualizada correctamente", nil)
	responseJSON(w, http.StatusOK, response)
}

// delete handler para eliminar una persona
func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || ID <= 0 {
		response := newResponse(Error, "El id debe ser un número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Ocurrió un error al eliminar el registro", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", nil)
	responseJSON(w, http.StatusOK, response)
}

// getByID handler para obtener una persona por su ID
func (p *person) getByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetByID(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Ocurrió un error al elminar el registro", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, response)
}

// getAll handler para obtener todas las personas
func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Hubo un problema al obtener todas las personas", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, response)
}