package context

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w responseWriter) ToStringE() (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &data)
	return data, err
}
