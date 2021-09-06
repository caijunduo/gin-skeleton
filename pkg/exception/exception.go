package exception

import "net/http"

type exception struct {
    status  int
    code    int
    message string
    err     error
}

type Exception interface {
    Status() int
    Code() int
    Message() string
    Error() error
}

func New(code int, message string) *exception {
    return &exception{http.StatusInternalServerError, code, message, nil}
}

func (e exception) Status() int {
    return e.status
}

func (e *exception) SetStatus(status int) *exception {
    e.status = status
    return e
}

func (e exception) Code() int {
    return e.code
}

func (e exception) Message() string {
    return e.message
}

func (e *exception) SetMessage(message string) *exception {
    e.message = message
    return e
}

func (e exception) Error() error {
    return e.err
}

func (e *exception) SetError(err error) *exception {
    e.err = err
    return e
}
