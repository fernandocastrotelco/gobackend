package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	println("=== test simple channel ===")
	simpleChannel()
	println("=== test simple channel ===")

	println("=== test buffered channel con tamaño 1 ===")
	bufChannel()
	println("=== test buffered channel ===")

	println("=== test buffered channel con tamaño 2 ===")
	buffChannel2()
	println("=== test buffered channel 2 ===")

	println("=== test for loop in channels ===")
	forLoopChannel()
	println("=== test for loop channel ===")
}

func simpleChannel() {
	// definimos un channel
	channel := make(chan string)
	// vamos a esperar la goroutine con esta variable
	var waitGroup sync.WaitGroup
	// le decimos q vamos a esperar una goroutine
	waitGroup.Add(1)
	// esta goroutine se ejecuta en paralelo, no bloquea el main thread
	go func() {
		// enviamos un mensaje por el channel
		channel <- "Hello World!"
		// esto no se va a ejecutar hasta q el channel se use
		println("Finishing goroutine")
		// marcamos como completa la tarea
		waitGroup.Done()
	}()
	// hacemos una pausa
	time.Sleep(time.Second)
	// el receptor del channel recien esta disponible despues de la pausa
	message := <-channel
	// una vez q el channel recibe el mensaje se desbloquea tmb el otro thread y se ejecuta al mismo tiempo la linea 31
	fmt.Println(message) // los mensajes se pueden imprimir en cualquier orden, ya que se ejecutan en dos threads diferentes

	waitGroup.Wait()
}

func bufChannel() {
	// con bufered channels indicamos la cantidad de mensajes antes de que el bloqueo se realice
	channel := make(chan string, 1)
	// ejecutamos la goroutine en otro thread
	go func() {
		// enviamos 1 mensaje por el channel
		channel <- "Hello World!"
		// como indicamos 1 en el buffered channel, este thread no se bloquea y sigue con la siguiente funcion
		println("Finishing goroutine") // esta linea se va a imprimir antes de la pausa
	}()
	// realizamos una pausa
	time.Sleep(time.Second)
	// recibimos el mensaje del channel
	message := <-channel
	// lo imprimimos al final
	fmt.Println(message)
}

func buffChannel2() {
	// indicamos tamaño del buff en 1 pero vamos a probar q pasa si mandamos 2 mensajes
	channel := make(chan string, 1)
	// mandamos la goroutine
	go func() {
		// enviamos los dos mensajes
		channel <- "Hello World! 1"
		// enviar otro mensaje va a bloquear este thread
		channel <- "Hello World! 2"
		// este mensaje no se va a imprimir hasta que el mensaje 1 sea recibido
		println("Finishing goroutine")
	}()
	// hacemos una pausa
	time.Sleep(time.Second)
	// recibir este mensaje despues de la pausa va a liberar el otro thread
	message := <-channel
	// todo continua y se imprime el mensaje 1
	fmt.Println(message)
	// el programa termina asi que nadie recibe el mensaje 2
}

func sendString(ch chan<- string, s string) {
	ch <- s
}

func receiver(helloCh, goodbyeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			println(msg)
		case msg := <-goodbyeCh:
			println(msg)
		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting")
			quitCh <- true
			break
		}
	}
}

func forLoopChannel() {
	ch := make(chan int)

	go func() {
		ch <- 1
		time.Sleep(time.Second)

		ch <- 2

		close(ch)
	}()
	for v := range ch {
		println(v)
	}
}
