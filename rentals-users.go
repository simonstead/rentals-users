package main

import (
	"fmt"
	"github.com/simonstead/rentals-users/handlers"
	"net/http"
	"os"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	PORT, err := determineListenAddress()
	if err != nil {
		PORT = ":4000"
	}

	http.HandleFunc("/", handlers.RootHandler)

	// server
	fmt.Printf("-\n- Listening on port %v\n-\n", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		panic(err)
	}
}
