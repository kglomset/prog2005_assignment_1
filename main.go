package main

import (
	"Assignment_1/handlers"
	"Assignment_1/util"
	"errors"
	"log"
	"net/http"
)

// The main function starts the timer for how long the service have been active and starts the service
// with all endpoints
func main() {
	handlers.StartUptimeTracking()

	http.HandleFunc(util.BookCountEndPoint, handlers.BookCountHandler)
	http.HandleFunc(util.ReadershipEndPoint, handlers.ReadershipHandler)
	http.HandleFunc(util.StatusEndPoint, handlers.GetStatusFromEndPoints)

	if err := http.ListenAndServe(util.DefaultPort, nil); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server gracefully shut down")
		} else {
			log.Fatalf("Server failed to start: %v", err)
		}
	}
}
