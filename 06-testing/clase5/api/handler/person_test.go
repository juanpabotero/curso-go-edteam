package handler

import (
	"bytes"
	"curso-go-edteam/06-testing/clase5/api/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

type responsePerson struct {
	MessageType string       `json:"message_type"`
	Message     string       `json:"message"`
	Data        model.Person `json:"data"`
}

func TestCreate_wrong_storage(t *testing.T) {
	// data simula una estructura correcta
	data := []byte(`{"name": "Alexys", "age": 40, "communities": []}`)
	// w es un ResponseWriter mock
	w := httptest.NewRecorder()
	// r es una Request mock
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	// se setea el header de la petición
	r.Header.Set("Content-Type", "application/json")
	// crea un nuevo contexto de echo
	e := echo.New()
	ctx := e.NewContext(r, w)
	// se crea una instancia de person con un storage mock que devuelve errores
	p := newPerson(&PersonStorageWrongMock{})
	err := p.create(ctx)
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Código de estado, se esperaba %d, se obtuvo %d", http.StatusInternalServerError, w.Code)
	}

	resp := responsePerson{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("no se pudo leer el cuerpo: %v", err)
	}

	wantMessage := "Hubo un problema al crear la persona"
	if resp.Message != wantMessage {
		t.Errorf("mensaje de error equivocado: se esperaba %q, se obtuvo %q", wantMessage, resp.Message)
	}
}

func TestCreate_wrong_structure(t *testing.T) {
	// data simula una estructura incorrecta
	data := []byte(`{"name": 5, "age": 40, "communities": []}`)
	// w es un ResponseWriter mock
	w := httptest.NewRecorder()
	// r es una Request mock
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	// se setea el header de la petición
	r.Header.Set("Content-Type", "application/json")
	// crea un nuevo contexto de echo
	e := echo.New()
	ctx := e.NewContext(r, w)
	// se crea una instancia de person con un storage mock que no devuelve errores
	p := newPerson(&PersonStorageOKMock{})
	err := p.create(ctx)
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo %v", err)
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("Código de estado, se esperaba %d, se obtuvo %d", http.StatusBadRequest, w.Code)
	}

	resp := responsePerson{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("no se pudo leer el cuerpo: %v", err)
	}

	wantMessage := "La persona no tiene una estructura correcta"
	if resp.Message != wantMessage {
		t.Errorf("mensaje de error equivocado: se esperaba %q, se obtuvo %q", wantMessage, resp.Message)
	}
}

func TestCreate(t *testing.T) {
	// data simula una estructura
	data := []byte(`{"name": "Alexys", "age": 40, "communities": []}`)
	// w es un ResponseWriter mock
	w := httptest.NewRecorder()
	// r es una Request mock
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	// se setea el header de la petición
	r.Header.Set("Content-Type", "application/json")
	// crea un nuevo contexto de echo
	e := echo.New()
	ctx := e.NewContext(r, w)
	// se crea una instancia de person con un storage mock que no devuelve errores
	storage := PersonStorageOKMock{}
	p := newPerson(&storage)
	err := p.create(ctx)
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo %v", err)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("Código de estado, se esperaba %d, se obtuvo %d", http.StatusCreated, w.Code)
	}

	resp := responsePerson{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("no se pudo leer el cuerpo: %v", err)
	}

	wantMessage := "Persona creada correctamente"
	if resp.Message != wantMessage {
		t.Errorf("mensaje de error equivocado: se esperaba %q, se obtuvo %q", wantMessage, resp.Message)
	}

	dataStorage, _ := storage.GetAll()
	t.Logf("Mock storage: %v", dataStorage)
}
