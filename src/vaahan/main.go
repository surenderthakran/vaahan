package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"vaahan/mapper"

	glog "github.com/golang/glog"
)

var staticFs = http.FileServer(http.Dir("/workspace/src/vaahan/static"))

func main() {
	// Overriding glog's logtostderr flag's value to print logs to stderr.
	flag.Lookup("logtostderr").Value.Set("true")
	// Calling flag.Parse() so that all flag changes are picked.
	flag.Parse()

	http.HandleFunc("/api/get_map", func(w http.ResponseWriter, r *http.Request) {
		mapData, err := mapper.GetMap()
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(mapData)
		if err != nil {
			log.Print(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	http.Handle("/", http.StripPrefix("/", staticFs))

	log.Fatal(http.ListenAndServe(":18770", nil))
}
