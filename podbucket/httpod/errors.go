package httpod

import "encoding/json"

type PodError struct {
	Error string `json:"message"`
}

const (
	INVALID_BODY = "Body is not JSON valid format"
)

func NewError(error *PodError) ([]byte, error) {
	return json.Marshal(error)
}
