package main

import (
	"flag"
	"fmt"
	glog "github.com/golang/glog"
	"log"
	"net/http"
)

func main() {
	// Overriding glog's logtostderr flag's value to print logs to stderr.
	flag.Lookup("logtostderr").Value.Set("true")
	// Calling flag.Parse() so that all flag changes are picked.
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		glog.Info("A new / request received!")
		fmt.Fprintf(w, "Hello")
	})

	err := http.ListenAndServe(":18770", nil) // Note: Not "localhost:18770" but ":18770"
	log.Fatal(err)
}
