package main

import (
	"fmt"
	"sync"
	"time"
)

type account struct {
	name    string
	balance int
}

func (a *account) String() string {
	return fmt.Sprintf("(%s: Balance: %d)", a.name, a.balance)
}

func transfer(amount int, source, dest *account) {
	if source.balance < amount {
		fmt.Printf("❌: %s\n", fmt.Sprintf("%s %s", source, dest))
		return
	}
	time.Sleep(time.Second)

	dest.balance += amount
	source.balance -= amount

	fmt.Printf("✅: %s\n", fmt.Sprintf("%s %s", source, dest))
}

func main() {
  // instancia de WaitGroup
	wg := sync.WaitGroup{}
	// instancia de Mutex
	mu := sync.Mutex{}
  // definir el contador, en este caso, agregar 2 Goroutines
	wg.Add(2)

	alexys := account{name: "Alexys", balance: 500}
	beto := account{name: "Beto", balance: 900}

	for _, value := range []int{300, 300} {
		go func(amount int) {
			// bloquear el Mutex, para que solo una Goroutine pueda acceder a la sección crítica
			mu.Lock()
			transfer(amount, &alexys, &beto)
			// desbloquear el Mutex, para que otra Goroutine pueda acceder a la sección crítica
			mu.Unlock()
			// cuando la Goroutine termina, se llama a Done y se resta 1 del contador
			wg.Done()
		}(value)
	}
  // esperar a que todas las Goroutines terminen, es decir, que el contador llegue a 0
	wg.Wait()
}