package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"vaahan/track"

	glog "github.com/golang/glog"
)

var staticFs = http.FileServer(http.Dir("/workspace/src/vaahan/static"))

func main() {
	// Overriding glog's logtostderr flag's value to print logs to stderr.
	flag.Lookup("logtostderr").Value.Set("true")
	// Calling flag.Parse() so that all flag changes are picked.
	flag.Parse()

	http.HandleFunc("/api/get_track", func(w http.ResponseWriter, r *http.Request) {
		trackData, err := track.GetTrack()
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(trackData)
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
