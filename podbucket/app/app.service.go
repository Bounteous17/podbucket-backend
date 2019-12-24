package app

import (
	"encoding/json"
	"net/http"

	"github.com/bounteous/podbucket-backend/podbucket/httpod"
)

type appStatus struct {
	Message string `json:"status"`
}

// Handler : simple http method switch
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		response, err := getAppStatus()
		if err == nil {
			httpod.Response(w, response)
		}
	}
}

func getAppStatus() ([]byte, error) {
	message := &appStatus{
		Message: "Up and running ...",
	}
	return json.Marshal(message)
}
