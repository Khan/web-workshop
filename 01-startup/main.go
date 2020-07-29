package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	if err := runServer(logger, port); err != nil {
		logger.Printf("Got error: %v", err)
		os.Exit(1)
	}
	logger.Println("Finished clean")
}

func runServer(logger *log.Logger, port string) error {
	// =========================================================================
	// App Starting

	logger.Printf("main : Listening on :%v", port)
	defer logger.Println("main : Completed")

	// Convert the Echo function to a type that implements http.Handler
	h := http.HandlerFunc(Echo)

	// Start a server listening on port 8000 and responding using Echo.
	if err := http.ListenAndServe(":"+port, h); err != nil {
		logger.Printf("error: listening and serving: %s", err)
		return err
	}
	return nil
}

// Echo is a basic HTTP Handler.
func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}
