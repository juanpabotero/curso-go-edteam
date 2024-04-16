package main

import (
	"fmt"
	"sync"
)

func main() {
	courses := make(map[string]string)
  // instancia de WaitGroup
	var wg sync.WaitGroup
	// instancia de Mutex
	var mu sync.Mutex
  // definir el contador, en este caso, agregar 2 Goroutines
	wg.Add(2)

	// crea una Goroutine
	go func() {
  	// bloquear el Mutex, para que solo una Goroutine pueda acceder a la sección crítica
		mu.Lock()
		courses["go desde cero"] = "Intermedio"
  	// desbloquear el Mutex, para que otra Goroutine pueda acceder a la sección crítica
		mu.Unlock()
  	// cuando la Goroutine termina, se llama a Done y se resta 1 del contador
		wg.Done()
	}()

	go func() {
		mu.Lock()
		courses["go concurrencia"] = "Avanzado"
		mu.Unlock()
		wg.Done()
	}()
  // esperar a que todas las Goroutines terminen, es decir, que el contador llegue a 0
	wg.Wait()

	fmt.Println(courses)
}