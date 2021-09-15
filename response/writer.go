package response

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

var Writer = &writer{
	Body: bytes.NewBufferString(""),
}

type writer struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w writer) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}
