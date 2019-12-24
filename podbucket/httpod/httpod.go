package httpod

import (
	"net/http"
)

type stringFunc func() string

// Response comment
func Response(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
