package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name  string `json:"name"`
	Age   string `json:"age"`
	Email string `json:"email,omitempty"`
}

// personHandler function acts as a translator between the Internet(Json) and go code(Structs)
func personHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//The translator(json decoding) it decode request body into person struct
	var p Person
	err := json.NewDecoder(r.Body).Decode(&p) //as you find json key related to person struct you store it in an p
	if err != nil {
		http.Error(w, "invalid json"+err.Error(), http.StatusBadRequest)
		return
	}

	//validation
	if p.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	//Shipping department(building an response)
	response := map[string]any{
		"message": "hello" + p.Name,
		"age":     p.Age,
	}
	//this is whare your internal go data leaves the server and travel across the internet to the client 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response) //converts a go value to Json and write it directly to the responseWriter
}
func main() {
	http.HandleFunc("/person", personHandler)
	fmt.Println("istening on :8080")
	http.ListenAndServe(":8080", nil)
}
