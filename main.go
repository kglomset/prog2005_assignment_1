package main

import (
	"Assignment_1/handlers"
	"Assignment_1/util"
	"log"
	"net/http"
)

func main() {
	// Error handling her?
	handlers.StartUptimeTracking()
	http.HandleFunc(util.BookCountEndPoint, handlers.BookCountHandler)
	http.HandleFunc(util.ReadershipEndPoint, handlers.ReadershipHandler)
	http.HandleFunc(util.StatusEndPoint, handlers.GetStatusFromEndPoints)
	log.Fatal(http.ListenAndServe(util.DefaultPort, nil))
}
