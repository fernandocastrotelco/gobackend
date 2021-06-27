package main

import (
	"sync"
	"time"
)

// struct q implementa mutex para manejar locks
type Counter struct {
	sync.Mutex
	value int
}

func main() {
	counter := Counter{}

	for i := 0; i < 10; i++ {
		// lanzamos 10 goroutines
		go func(i int) {
			// las 10 goroutines van a querer modificar el mismo campo
			// cuando indicamos lock bloqueamos el acceso a la struct Counter
			// las otras goroutines no van a poder acceder a value
			counter.Lock()
			// solo una goroutine va a poder modificar value, el resto se queda esperando
			counter.value++
			// la goroutine presente desbloquea la struct Counter,
			// asi la proxima va a bloquear y acceder a value
			counter.Unlock()
		}(i)
	}
	time.Sleep(time.Second)
	counter.Lock()
	println(counter.value)
	counter.Unlock()
}
