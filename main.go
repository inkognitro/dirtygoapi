package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello from endpoint"))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
