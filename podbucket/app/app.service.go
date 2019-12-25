package app

import (
	"encoding/json"
	"net/http"

	"github.com/bounteous/podbucket-backend/podbucket/httpod"
)

type getApp struct {
	Message string `json:"status"`
}

type postApp struct {
	Message string `json:"message"`
}

// Handler : simple http method switch
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		response, err := getAppStatus()
		if err == nil {
			httpod.StatusOK(w, response)
		}
	case "POST":
		response, err := postAppStatus()
		if err == nil {
			httpod.StatusBadRequest(w, response)
		}
	}
}

func getAppStatus() ([]byte, error) {
	message := &getApp{
		Message: "Up and running ...",
	}
	return json.Marshal(message)
}

func postAppStatus() ([]byte, error) {
	message := &postApp{
		Message: "Not implemented yet",
	}
	return json.Marshal(message)
}
