// es buena practica tener un archivo con el nombre del paquete, en este caso model.go,
// para tener todas las constantes, errores y metodos globales del paquete model
package model

import "errors"

var (
	// ErrPersonCanNotBeNil la persona no puede ser nula
	ErrPersonCanNotBeNil = errors.New("la persona no puede ser nula")
	// ErrIDPersonDoesNotExists la persona no existe
	ErrIDPersonDoesNotExists = errors.New("la persona no existe")
)