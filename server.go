package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	log.Printf("Listening on :%s...", httpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil)
	if err != nil {
		log.Fatal(err)
	}
}
