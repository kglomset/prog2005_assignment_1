package main

import (
	"Assignment_1/handlers"
	"Assignment_1/util"
	"log"
	"net/http"
)

func main() {

	// Put in some error handling for starting the server
	handlers.StartUptimeTracking()
	http.HandleFunc(util.BookCountEndPoint, handlers.BookCountHandler)
	http.HandleFunc(util.ReadershipEndPoint, handlers.ReadershipHandler)
	http.HandleFunc(util.StatusEndPoint, handlers.GetStatus)
	log.Fatal(http.ListenAndServe(util.DefaultPort, nil))
}
