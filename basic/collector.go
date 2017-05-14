package basic

import (
	"net/http"
	"log"
	"time"
)

// buffered channel to send WorkRequests
var WorkQueue = make(chan WorkRequest, 1000)

func Collector(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("INFO", "Only POST requests allowed")
		return
	}


	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Wrong delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	work := WorkRequest{Name: name, Delay: delay}

	WorkQueue <- work
	log.Println("INFO", "Work request queued of Name: "+name)
	w.WriteHeader(http.StatusCreated)
	return
}