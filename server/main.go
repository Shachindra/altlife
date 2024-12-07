package main

import (
	"net/http"
)

func deploy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("../ui")))
	http.HandleFunc("/api/v1.0/deploy", deploy)
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		panic(err)
	}
}
