package mycustomservice

import (
	"encoding/json"
	"net/http"

	"github.com/bounteous/podbucket-backend/podbucket/httpod"
)

type myServiceList struct {
	List string `json:"list"`
}

// Handler : simple http method switch
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		response, err := getCustomServices()
		if err == nil {
			httpod.Response(w, response)
		}
	}
}

func getCustomServices() ([]byte, error) {
	list := &myServiceList{
		List: "No custom services already created ...",
	}
	return json.Marshal(list)
}
