package main

import (
	"log"
	"net/http"
)

func main() {
	const port = ":3000"
	mux := http.NewServerMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/post", post)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
