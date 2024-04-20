package storage

import (
	"curso-go-edteam/06-testing/clase5/api/model"
	"testing"
)

func TestCreate(t *testing.T) {
	m := NewMemory()
	table := []struct {
		person *model.Person
		err    error
	}{
		{nil, model.ErrPersonCanNotBeNil},
		{&model.Person{Name: "Alvaro"}, nil},
		{&model.Person{Name: "Beto"}, nil},
		{&model.Person{Name: "Alexys"}, nil},
	}
	for _, v := range table {
		err := m.Create(v.person)
		if err != v.err {
			t.Errorf("se esperaba %v, se obtuvo %v", v.err, err)
		}
	}

	currentID := len(table) - 1
	if m.currentID != currentID {
		t.Errorf("el currentID debería ser %d, y es %d", currentID, m.currentID)
	}
}
