package main

import (
	"fmt"
	"sync"
)

func main() {
  // instancia de WaitGroup
	var wg sync.WaitGroup
  // definir el contador, en este caso, agregar 5 Goroutines
	wg.Add(5)
	for i := 0; i < 5; i++ {
		// crea una Goroutine
		go func(j int) {
			fmt.Println(j)
			// cuando la Goroutine termina, se llama a Done y se resta 1 del contador
			wg.Done()
		}(i)
	}
  // esperar a que todas las Goroutines terminen, es decir, que el contador llegue a 0
	wg.Wait()
}