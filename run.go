package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// Internal
	"github.com/bounteous/podbucket-backend/custom/mycustomservice"
	"github.com/bounteous/podbucket-backend/podbucket/app"
)

func main() {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", app.Handler)
	muxRouter.HandleFunc("/mycustomservices", mycustomservice.Handler)
	log.Fatal(http.ListenAndServe(":1337", muxRouter))
}
