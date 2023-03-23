package main

import {
  "fmt"
  "net/http"
}

const port = ":3000"

func main() {
  mux := http.NewServerMux()
  mux.HandleFunc("/home", index)

	mux.Handle ("/static/", http.StripPrefix)

	fmt.Println("Start server")
  http.ListenAndServe(port, mux)
}
