package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jacksondprice/amber-portal/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/tutors", handlers.ListTutors)
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, mux))
}
