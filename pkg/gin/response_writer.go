package gin

import (
    "bytes"
    "github.com/gin-gonic/gin"
)

var Writer *responseWriter

type responseWriter struct {
    gin.ResponseWriter
    Body *bytes.Buffer
}

func NewResponseWriter(rw gin.ResponseWriter) {
    Writer = &responseWriter{
        ResponseWriter: rw,
        Body:           bytes.NewBufferString(""),
    }
}

func (w responseWriter) Write(b []byte) (int, error) {
    w.Body.Write(b)
    return w.ResponseWriter.Write(b)
}
