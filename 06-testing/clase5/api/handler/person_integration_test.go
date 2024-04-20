package handler_test

import (
	"bytes"
	"curso-go-edteam/06-testing/clase5/api/handler"
	"curso-go-edteam/06-testing/clase5/api/model"
	"curso-go-edteam/06-testing/clase5/api/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestCreate_integration(t *testing.T) {
	// Cleanup se ejecuta al finalizar el test
	t.Cleanup(func() {
		cleanData(t)
	})
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

	store := storage.NewPsql()
	p := handler.NewPerson(&store)
	err := p.Create(&p, ctx)
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo %v", err)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("Código de estado, se esperaba %d, se obtuvo %d", http.StatusCreated, w.Code)
	}
}

type responsePerson struct {
	MessageType string        `json:"message_type"`
	Message     string        `json:"message"`
	Data        model.Persons `json:"data"`
}

func TestGetAll_integration(t *testing.T) {
	// Cleanup se ejecuta al finalizar el test
	t.Cleanup(func() {
		cleanData(t)
	})
	insertData(t)
	// w es un ResponseWriter mock
	w := httptest.NewRecorder()
	// r es una Request mock
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	// se setea el header de la petición
	r.Header.Set("Content-Type", "application/json")
	// crea un nuevo contexto de echo
	e := echo.New()
	ctx := e.NewContext(r, w)

	store := storage.NewPsql()
	p := handler.NewPerson(&store)
	err := p.GetAll(&p, ctx)
	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo %v", err)
	}

	var response responsePerson
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatalf("no se pudo hacer marshal de la respuesta: %v", err)
	}

	lenPersonsWant := 2
	lenPersonsGot := len(response.Data)
	if lenPersonsGot != lenPersonsWant {
		t.Errorf("Se esperaban %d personas, se obtuvieron %d", lenPersonsWant, lenPersonsGot)
	}
}

func insertData(t *testing.T) {
	store := storage.NewPsql()
	err := store.Create(&model.Person{Name: "Alexys", Age: 40})
	if err != nil {
		t.Fatalf("no se pudo registrar la persona %v", err)
	}

	err = store.Create(&model.Person{Name: "Matthew", Age: 4})
	if err != nil {
		t.Fatalf("no se pudo registrar la persona %v", err)
	}

	store.CloseDB()
}

func cleanData(t *testing.T) {
	store := storage.NewPsql()
	err := store.DeleteAll()
	if err != nil {
		t.Fatalf("no se pudo limpiar la tabla %v", err)
	}

	store.CloseDB()
}
