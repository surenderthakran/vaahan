package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"vaahan/car"
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
		trackID := r.URL.Query()["id"][0]
		track, err := track.GetTrack(trackID)
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(track)
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	http.HandleFunc("/api/init_car", func(w http.ResponseWriter, r *http.Request) {
		// trackID := r.URL.Query()["id"][0]
		// track, err := track.GetTrack(trackID)
		// if err != nil {
		// 	glog.Error(err)
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		car := car.New()

		response, err := json.Marshal(car)
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	http.HandleFunc("/api/get_car", func(w http.ResponseWriter, r *http.Request) {
		car := car.GetCar()

		response, err := json.Marshal(car)
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	http.Handle("/", http.StripPrefix("/", staticFs))

	log.Print("Listening on 18770...")
	log.Fatal(http.ListenAndServe(":18770", nil))
}
