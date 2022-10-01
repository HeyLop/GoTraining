package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello http server"))
	})
	http.ListenAndServe(":8081", nil)
}
