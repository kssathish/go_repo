package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to GO")
	resp, err := http.Get("https://www.google.com")

	if err != nil {
		log.Fatal(err)
	}

	if int(resp.StatusCode) == http.StatusOK {
		fmt.Printf("Connected to google.com with status code: %d", resp.StatusCode)
	}
}
