package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(200)
		w.Write([]byte("Home Page"))

	})

	http.ListenAndServe(":8080", nil)
}
