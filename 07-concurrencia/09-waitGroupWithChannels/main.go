package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"http://localhost:1234?duration=3s",
	"http://localhost:1234?duration=1s",
	"http://localhost:1234?duration=5s",
}

func main() {
	fetchConcurrentCancelation(urls)
}

func fetchSequential(urls []string) {
	for _, url := range urls {
		fetch(url)
	}
}

func fetchConcurrent(urls []string) {
	// instancia de WaitGroup
	var wg sync.WaitGroup
	// definir el contador, en este caso, agregar 3 Goroutines
	wg.Add(3)

	for _, url := range urls {
		// crear Goroutine
		go func(u string) {
			fetch(u)
			// cuando la Goroutine termina, se llama a Done y se resta 1 del contador
			wg.Done()
		}(url)
	}
	// esperar a que todas las Goroutines terminen, es decir, que el contador llegue a 0
	wg.Wait()
}

func fetchConcurrentCSP(urls []string) {
	// crear canal
	signal := make(chan struct{})

	for _, url := range urls {
		// crear Goroutine
		go func(u string) {
			fetch(u)
			// enviar valor al canal
			signal <- struct{}{}
		}(url)
	}
	// recibir valores del canal
	<-signal
	<-signal
	<-signal
}

func fetchConcurrentCancelation(urls []string) {
	// crear canal
	done := make(chan struct{})

	for _, url := range urls {
		// crear Goroutine
		go func(u string) {
			fetch(u)
			// select: permite recibir valores de varios canales, funciona como un switch pero para canales
			select {
			// recibir valores del canal
			case <-done:
				return
			}
		}(url)
	}

	select {
	case <-time.After(time.Second * 4):
		close(done)
	}
}

func fetch(url string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatalf("failed url: %s, err: %v", url, err)
	}
	log.Println(url, ": ", resp.StatusCode)
}
