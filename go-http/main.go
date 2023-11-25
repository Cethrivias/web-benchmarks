package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, _ := json.Marshal(Hello{Hello: "world"})
		w.Write(res)
	})

	http.ListenAndServe(":8080", nil)
}

type Hello struct {
	Hello string `json:"hello"`
}
