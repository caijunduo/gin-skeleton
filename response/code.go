package response

import (
	"net/http"
)

var (
	OK                   = &response{status: http.StatusOK, code: 0, message: "OK"}
	NotFound             = &response{status: http.StatusNotFound, code: 404, message: "Not Found"}
	InvalidParameters    = &response{status: http.StatusInternalServerError, code: 10000, message: "Invalid Parameter"}
	InvalidAuthorization = &response{status: http.StatusUnauthorized, code: 10002, message: "Invalid authorization"}
	UnAuthorization      = &response{status: http.StatusInternalServerError, code: 10003, message: "Authorization generate failed"}
)
