package clase3

import (
	"reflect"
	"testing"
)

func TestPerro(t *testing.T) {
	want := &Perro{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "criollo",
		},
	}
	got := &Perro{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "criolla",
		},
	}

	// t.Logf("want %p, got %p", want, got)
	// reflect.DeepEqual compara los valores de dos estructuras, slices, punteros
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Se esperaba: %v, se obtuvo %v", want, got)
	}
}
