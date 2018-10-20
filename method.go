package restpec

import "net/http"

func Ok(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 200, message, data, meta)
}

func BadRequest(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 400, message, data, meta)
}

func Unauthorized(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 401, message, data, meta)
}

func Forbidden(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 403, message, data, meta)
}
