package main

import (
	"fmt"
	"sync"
)

func main() {
	// instancia de Mutex
	mu := sync.Mutex{}
  // instancia de WaitGroup
	wg := sync.WaitGroup{}
  // definir el contador, en este caso, agregar 1 Goroutines
	wg.Add(1)

	data := 1

	// crea una Goroutine
	go func() {
  	// bloquear el Mutex, para que solo una Goroutine pueda acceder a la sección crítica
		mu.Lock()
		data++
  	// desbloquear el Mutex, para que otra Goroutine pueda acceder a la sección crítica
		mu.Unlock()
  	// cuando la Goroutine termina, se llama a Done y se resta 1 del contador
		wg.Done()
	}()
  // esperar a que todas las Goroutines terminen, es decir, que el contador llegue a 0
	wg.Wait()
	mu.Lock()
	fmt.Println(data)
	mu.Unlock()
}