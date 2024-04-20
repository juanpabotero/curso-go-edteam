package handler

import (
	"curso-go-edteam/06-testing/clase5/api/model"
	"errors"
)

// Mock de PersonStorage sin errores
type PersonStorageOKMock struct{}

func (psm *PersonStorageOKMock) Create(person *model.Person) error {
	return nil
}
func (psm *PersonStorageOKMock) Update(ID int, person *model.Person) error {
	return nil
}
func (psm *PersonStorageOKMock) Delete(ID int) error {
	return nil
}
func (psm *PersonStorageOKMock) GetByID(ID int) (model.Person, error) {
	return model.Person{}, nil
}
func (psm *PersonStorageOKMock) GetAll() (model.Persons, error) {
	return nil, nil
}

// Mock de PersonStorage con errores
type PersonStorageWrongMock struct{}

func (psm *PersonStorageWrongMock) Create(person *model.Person) error {
	return errors.New("error")
}
func (psm *PersonStorageWrongMock) Update(ID int, person *model.Person) error {
	return errors.New("error")
}
func (psm *PersonStorageWrongMock) Delete(ID int) error {
	return errors.New("error")
}
func (psm *PersonStorageWrongMock) GetByID(ID int) (model.Person, error) {
	return model.Person{}, errors.New("error")
}
func (psm *PersonStorageWrongMock) GetAll() (model.Persons, error) {
	return nil, errors.New("error")
}
