// es buena practica tener un archivo con el nombre del paquete, en este caso model.go,
// para tener todas las constantes, errores y metodos globales del paquete model
package handler

import "curso-go-edteam/04-api/03-api/model"

// Storage interfaz para el almacenamiento de personas
type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}