package errno

import (
    "net/http"
    "skeleton/pkg/exception"
)

var (
    InvalidParameters    = exception.New(10000, "Invalid parameter").SetStatus(http.StatusBadRequest)
    InvalidSignature     = exception.New(10001, "Signature verification failed").SetStatus(http.StatusBadRequest)
    InvalidAuthorization = exception.New(10002, "Invalid authorization").SetStatus(http.StatusUnauthorized)
    UnAuthorization      = exception.New(10003, "Authorization generate failed").SetStatus(http.StatusInternalServerError)
)
