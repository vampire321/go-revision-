package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email,omitempty"`
}
//personHandler function acts as a translator between the Internet(Json) and go code(Structs)
func personHandler(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//The translator(json decoding) it decode request body into person struct
	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil{
		http.Error(w,"invalid json"+err.Error(),http.StatusBadRequest)
		return
	}

	
	