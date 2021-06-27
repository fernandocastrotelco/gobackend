package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// server simple
type MyServer struct{}

// funcion simple sin decorators solo implementa la interfaz Handler del pkg http
func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator!")
}

// primer decorator que va a extender la funcion handler
type LoggerMiddleware struct {
	Handler   http.Handler
	LogWriter io.Writer
}

// agrega logs antes de ejecutar la funcion simple de Handler
func (l *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(l.LogWriter, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(l.LogWriter, "Host: %s\n", r.Host)
	fmt.Fprintf(l.LogWriter, "Content Length: %d\n", r.ContentLength)
	fmt.Fprintf(l.LogWriter, "Method: %s\n", r.Method)
	fmt.Fprintf(l.LogWriter, "-----------------------------------\n")
	l.Handler.ServeHTTP(w, r)
}

// agrega authentication al server
type SimpleAuthMiddleware struct {
	Handler  http.Handler
	User     string
	Password string
}

func (s *SimpleAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()

	if ok {
		if user == s.User && pass == s.Password {
			s.Handler.ServeHTTP(w, r)
		} else {
			fmt.Fprintf(w, "User or password incorrect\n")
		}
	} else {
		fmt.Fprintln(w, "Error trying to retrieve data from Basic auth")
	}
}

func main() {
	fmt.Println("Enter the type number of server you want to launch from the" +
		" following:")
	fmt.Println("1.- Plain server")
	fmt.Println("2.- Server with logging")
	fmt.Println("3.- Server with logging and authentication")

	var selection int
	fmt.Fscanf(os.Stdin, "%d", &selection)

	var mySuperServer http.Handler

	switch selection {
	case 1:
		// instancio mi server de la manera mas simple
		mySuperServer = new(MyServer)
	case 2:
		// instancio el server con el logger, pero adentro instancio el server simple
		mySuperServer = &LoggerMiddleware{
			Handler:   new(MyServer),
			LogWriter: os.Stdout,
		}
	case 3:
		var user, password string

		fmt.Println("Enter user and password separated by a space")
		fmt.Fscanf(os.Stdin, "%s %s", &user, &password)
		// instancio el server con el logger, adentro le agrego el decorator auth y adentro del auth agrego el server simple
		mySuperServer = &LoggerMiddleware{
			Handler: &SimpleAuthMiddleware{
				Handler:  new(MyServer),
				User:     user,
				Password: password,
			},
			LogWriter: os.Stdout,
		}
	default:
		mySuperServer = new(MyServer)
	}

	http.Handle("/", mySuperServer)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
