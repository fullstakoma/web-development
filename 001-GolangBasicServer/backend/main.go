package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // added CORS header
	w.Header().Set("Content-Type", "application/json")
	resp := Response{Message: "Hello from Go!"}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	// ポート8080でHTTPサーバーを起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
