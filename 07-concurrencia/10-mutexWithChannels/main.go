package main

import (
	"fmt"
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

type bankOperation struct {
	amount int
	done   chan struct{}
}

func main() {
	// crear canales
	signal := make(chan struct{})
	transaction := make(chan *bankOperation)

	alexys := account{name: "Alexys", balance: 500}
	beto := account{name: "Beto", balance: 900}

	// cajero
	go func() {
		for {
			// recibir valores del canal
			request := <-transaction
			transfer(request.amount, &alexys, &beto)
			// enviar valores al canal
			request.done <- struct{}{}
		}
	}()

	for _, value := range []int{300, 300} {
		go func(amount int) {
			requestTransaction := bankOperation{amount: amount, done: make(chan struct{})}
			// enviar valores al canal
			transaction <- &requestTransaction
			// recibir valores del canal
			signal <- <-requestTransaction.done
		}(value)
	}
	// recibir valores del canal
	<-signal
	<-signal
}