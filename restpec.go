package restpec

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

func restpec(rw http.ResponseWriter, code int, message string, data, meta interface{}) {
	response, _ := json.Marshal(Response{
		message,
		data,
		meta,
	})

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(code)
	rw.Write(response)
}
