package main

import (
	"log"
	"net/http"
)

func main() {
	staticFs := http.FileServer(http.Dir("/workspace/src/vaahan/static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFs))
	log.Fatal(http.ListenAndServe(":18770", nil))
}
