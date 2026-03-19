package main

import (
	"fmt"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func basic() {
	http.HandleFunc("/health", HealthHandler)
	fmt.Println("starting the server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server isnt working")
	}
}
