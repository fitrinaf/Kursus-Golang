package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := 1491
	http.Handle("/", http.FileServer(http.Dir("polymer")))
	fmt.Printf("Server Running on Port %d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
