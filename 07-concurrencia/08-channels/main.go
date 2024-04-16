package main

import (
	"fmt"
	"time"
)

func main() {
	// crear canales
	number := make(chan int)
	signal := make(chan struct{})
	// crear gorutinas
	go receive(signal, number)
	send(number)

	signal <- struct{}{}
}

// chan<- int: canal de solo escritura, solo puede enviar valores
func send(number chan<- int) {
	// enviar valores al canal
	number <- 1
	number <- 2
	number <- 3
	number <- 4
	number <- 5
	time.Sleep(time.Nanosecond)
	number <- 6
}

// <-chan int: canal de solo lectura, solo puede recibir valores
func receive(signal <-chan struct{}, number <-chan int) {
	for {
		// select: permite recibir valores de varios canales, funciona como un switch pero para canales
		select {
		// recibir valores del canal
		case v := <-number:
			fmt.Println(v)
		case <-signal:
			return
		default:
			fmt.Println("🤔")
		}
	}
}
