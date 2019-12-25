package httpod

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bounteous/podbucket-backend/podbucket/clioutput"
)

const (
	statusOK                  = http.StatusOK
	statusInternalServerError = http.StatusInternalServerError
	statusBadRequest          = http.StatusBadRequest
)

func send(w http.ResponseWriter, response []byte, status int) {
	log(response, status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func log(response []byte, status int) {
	responseoutput := strings.Join([]string{strconv.Itoa(status), " ", string(response)}, "")
	switch status {
	case statusOK:
		clioutput.Success(responseoutput)
	case statusBadRequest:
		clioutput.Error(responseoutput)
	}
}

// StatusOK http send
func StatusOK(w http.ResponseWriter, response []byte) {
	send(w, response, statusOK)
}

// StatusInternalServerError http send
func StatusInternalServerError(w http.ResponseWriter, response []byte) {
	send(w, response, statusInternalServerError)
}

// StatusBadRequest http send
func StatusBadRequest(w http.ResponseWriter, response []byte) {
	send(w, response, statusBadRequest)
}
