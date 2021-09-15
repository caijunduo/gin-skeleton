package response

import (
	"net/http"
)

var (
	Example              = &Response{status: http.StatusOK, code: 0, message: "Example"}
	OK                   = &Response{status: http.StatusOK, code: 0, message: "OK"}
	NotFound             = &Response{status: http.StatusNotFound, code: 404, message: "Not Found"}
	InternalServerError  = &Response{status: http.StatusInternalServerError, code: 500, message: "Internal Server Error"}
	InvalidParameters    = &Response{status: http.StatusInternalServerError, code: 10000, message: "Invalid Parameter"}
	InvalidAuthorization = &Response{status: http.StatusUnauthorized, code: 10002, message: "Invalid authorization"}
	UnAuthorization      = &Response{status: http.StatusInternalServerError, code: 10003, message: "Authorization generate failed"}
)
