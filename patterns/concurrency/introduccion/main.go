package main

import (
	"fmt"
	"strings"
	"sync"
)

// wait nos ayuda a sincronizar las goroutines
var wait sync.WaitGroup

func main() {
	// le avisamos q vamos a correr una goroutine
	wait.Add(1)
	// funcion q va a correr la goroutine
	// pasamos una funcion como parametro para manejar los callbacks
	toUpperAsync("Hello Callbacks!", func(v string) {
		fmt.Printf("Callback: %s\n", strings.ToUpper(v))
		// le avisamos al wait q terminó la goroutine q está esperando
		wait.Done()
	})

	println("Waiting async response...")
	// acá espera a que las goroutines terminen
	wait.Wait()
}

func toUpperAsync(word string, f func(string)) {
	// ejecuta la funcion con el parametro en una goroutine
	go func() {
		f(word)
	}()
}
