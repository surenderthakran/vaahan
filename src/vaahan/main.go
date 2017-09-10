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

type Sim struct {
	Car   *car.Car     `json:"car"`
	Track *track.Track `json:"track"`
}

var staticFs = http.FileServer(http.Dir("/workspace/src/vaahan/static"))

func main() {
	// Overriding glog's logtostderr flag's value to print logs to stderr.
	flag.Lookup("logtostderr").Value.Set("true")
	// Calling flag.Parse() so that all flag changes are picked.
	flag.Parse()

	http.HandleFunc("/api/get_sim", func(w http.ResponseWriter, r *http.Request) {
		glog.Info("Request: /api/get_sim")
		track, err := track.GetTrack()
		if err != nil {
			glog.Errorf("unable to get sim: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		car, err := car.GetCar()
		if err != nil {
			glog.Errorf("unable to get sim: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sim := &Sim{
			Car:   car,
			Track: track,
		}

		response, err := json.Marshal(sim)
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	http.HandleFunc("/api/init_car", func(w http.ResponseWriter, r *http.Request) {
		glog.Info("Request: /api/init_car")
		car, err := car.InitCar()
		if err != nil {
			glog.Errorf("unable to initialize car: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
		glog.Info("Request: /api/get_car")
		car, err := car.GetCar()
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(car)
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	http.HandleFunc("/api/drive_car", func(w http.ResponseWriter, r *http.Request) {
		glog.Info("Request: /api/drive_car")
		car, err := car.GetCar()
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		car.Drive()

		response, err := json.Marshal(car)
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	http.HandleFunc("/api/stop_car", func(w http.ResponseWriter, r *http.Request) {
		glog.Info("Request: /api/stop_car")
		car, err := car.GetCar()
		if err != nil {
			glog.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		car.Stop()

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
