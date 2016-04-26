package main

import (
	"fmt"
	"html"
	"log"
	"os"

	"net/http"
	"github.com/gorilla/mux"

	"encoding/json"
	"runtime"
)

func main() {
	var port string
	port = os.Getenv("PORT")
	if (port == "") {
		port = "8080"
	}

	root,err := GetRootFolder()
	if (err != nil) {
		 log.Fatal("Cannot find working directory")
	}

	log.Printf("Starting server: %s [%s] ", port, root)

	index, err := LoadModel(root)
	if(err != nil) {
		panic(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HandleIndexRequest)
	router.HandleFunc("/events", HandleEvents(index))
	router.HandleFunc("/sessions/{eventId}", HandleEventSessions(root, index))

	log.Fatal(http.ListenAndServe(":" + port, router))
}

func HandleIndexRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server : %q\n", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "       : %s %s \n", runtime.GOOS, runtime.GOARCH)
	fmt.Fprintf(w, "       : %s \n", runtime.Version())
	fmt.Fprintf(w, "       : %s \n", os.Getenv("GOPATH"))
}

// Requests: Events -----------------------------------------------------------

func HandleEvents(index EventsIndex) func (w http.ResponseWriter, r *http.Request) {

	var events = Events{}
	for _, event := range index {
		events = append(events, event)
	}
	SortEventsById(events)

	/*
	 * handler
	 */
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(events)
	}
}

// Requests: Sessions ---------------------------------------------------------

func HandleEventSessions(root string, index EventsIndex) func (w http.ResponseWriter, r *http.Request) {

	/*
	 * handler
	 */
	return func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		eventId := vars["eventId"]
		if event, ok := index[eventId]; ok {
			log.Printf("+Sessions: %q", html.EscapeString(eventId))
			if data, err := ReadEventFile(root, event); err == nil {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, string(data[:]))
			} else {
				http.NotFound(w, r)
			}
		} else {
			log.Printf("-Sessions: %q", html.EscapeString(eventId))
			http.NotFound(w, r)
		}
	}
}

// ----------------------------------------------------------------------------
