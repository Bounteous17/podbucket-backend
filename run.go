package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	// Internal
	"github.com/bounteous/podbucket-backend/custom/mycustomservice"
	"github.com/bounteous/podbucket-backend/podbucket/app"
	"github.com/bounteous/podbucket-backend/podbucket/clioutput"
)

const (
	// Port listening port for http incoming traffic
	Port int = 1337
)

func main() {
	clioutput.Info(strings.Join([]string{"Server listening on port ", strconv.Itoa(Port)}, ""))
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", app.Handler)
	muxRouter.HandleFunc("/mycustomservices", mycustomservice.Handler)
	log.Fatal(http.ListenAndServe(strings.Join([]string{":", strconv.Itoa(Port)}, ""), muxRouter))
}
