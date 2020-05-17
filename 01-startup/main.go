package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	err := runServer(logger)
	if err == nil {
		logger.Println("finished clean")
		os.Exit(0)
	} else {
		logger.Printf("Got error: %v", err)
		os.Exit(1)
	}
}

func runServer(logger *log.Logger) error {
	// =========================================================================
	// App Starting

	logger.Printf("main : Started")
	defer logger.Println("main : Completed")

	// Convert the Echo function to a type that implements http.Handler
	h := http.HandlerFunc(Echo)

	// Start a server listening on port 8000 and responding using Echo.
	if err := http.ListenAndServe("localhost:8000", h); err != nil {
		logger.Printf("error: listening and serving: %s", err)
		return err
	}
	return nil
}

// Echo is a basic HTTP Handler.
func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}