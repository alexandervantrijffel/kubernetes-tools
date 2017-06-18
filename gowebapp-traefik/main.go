package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Request: ", r.URL.Path)
	fmt.Fprint(w, "Hello World! V1.0")
}

func main() {
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		ex := recover()
		if ex != nil {
			// "%v" prints the value of ex
			// for strings, it is the string, for errors .Error() method, for Stringer the .String() etc
			// Errorf returns an error instead of a string
			log.Fatal("Unhandled exception", ex)
		}
	}()

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}
	log.Printf("Starting server on port %v", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}
