package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = ":8080"

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is port%s web page ", port)
}
func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		log.Printf("TIME [%v]:,RequestMethod:%s,ReuestURI:%s", t, r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	})

}

func main() {
	http.ListenAndServe(port, logHandler(http.HandlerFunc(greeting)))
}
