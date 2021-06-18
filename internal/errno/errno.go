package errno

import (
    "net/http"
    "skeleton/pkg/exception"
)

var (
    InvalidParameters    = exception.New(10000, "无效的参数").SetStatus(http.StatusBadRequest)
    InvalidAuthorization = exception.New(10001, "无效的Authorization").SetStatus(http.StatusUnauthorized)
    UnAuthorization      = exception.New(10002, "Authorization生成失败")
)
