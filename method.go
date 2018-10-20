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

func NotFound(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 404, message, data, meta)
}

func MethodNotAllowed(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 405, message, data, meta)
}

func UnprocessableEntity(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 422, message, data, meta)
}

func InternalServerError(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 500, message, data, meta)
}

func NotImplemented(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 501, message, data, meta)
}

func BadGateway(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 502, message, data, meta)
}

func ServiceUnavailable(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 503, message, data, meta)
}

func GatewayTimeout(rw http.ResponseWriter, message string, data, meta interface{}) {
	restpec(rw, 504, message, data, meta)
}
