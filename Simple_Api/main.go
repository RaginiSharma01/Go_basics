package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // query parametrs

	if name == "" {
		name = " Ragini"
	}

	res := Response{
		Message: "Hello" + name,
	}

	w.Header().Set("Conten-Type", "application/json")

	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/greet", greetHandler)

	http.ListenAndServe(":3000", nil)
}
