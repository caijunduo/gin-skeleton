package request

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DefaultValue interface {
	Default()
}

type ValidateValue interface {
	Validate() error
}

const (
	getAllKey  = "__getAll__"
	postAllKey = "__postAll__"
	jsonAllKey = "__jsonAll__"
	allKey     = "__all__"
	headersKey = "__headers__"
)

func GetAll(c *gin.Context) (data map[string]interface{}) {
	if ctxData := c.GetStringMap(getAllKey); ctxData != nil {
		return ctxData
	}
	data = make(map[string]interface{}, 0)
	if c.Request.Method == http.MethodGet {
		if c.Request.URL.Query() != nil {
			for k, v := range c.Request.URL.Query() {
				data[k] = v[0]
			}
			c.Set(getAllKey, data)
		}
	}
	return
}

func PostAll(c *gin.Context) (data map[string]interface{}) {
	if ctxData := c.GetStringMap(postAllKey); ctxData != nil {
		return ctxData
	}
	data = make(map[string]interface{}, 0)
	if c.Request.Method == http.MethodPost {
		if c.ContentType() == "application/json" {
			return JsonAll(c)
		}
		if c.Request.ParseForm() == nil {
			_ = c.Request.ParseMultipartForm(32 << 20)
		}
		formData := make(url.Values)
		switch c.ContentType() {
		case "multipart/form-data":
			formData = c.Request.PostForm
		case "application/x-www-form-urlencoded":
			formData = c.Request.Form
		}
		if formData != nil {
			for k, v := range formData {
				data[k] = v[0]
			}
		}
		c.Set(postAllKey, data)
	}
	return
}

func JsonAll(c *gin.Context) (data map[string]interface{}) {
	if ctxData := c.GetStringMap(jsonAllKey); ctxData != nil {
		return ctxData
	}
	data = make(map[string]interface{}, 0)
	if c.ContentType() == "application/json" {
		var b []byte
		if cb, ok := c.Get(gin.BodyBytesKey); ok {
			if cbb, ok := cb.([]byte); ok {
				b = cbb
			}
		}
		if b == nil {
			b, _ = c.GetRawData()
			c.Set(gin.BodyBytesKey, b)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		}
		_ = json.Unmarshal(b, &data)
		c.Set(jsonAllKey, data)
	}
	return
}

func All(c *gin.Context) (data map[string]interface{}) {
	if ctxData := c.GetStringMap(allKey); ctxData != nil {
		return ctxData
	}
	data = make(map[string]interface{}, 0)
	for k, v := range GetAll(c) {
		data[k] = v
	}
	for k, v := range JsonAll(c) {
		data[k] = v
	}
	for k, v := range PostAll(c) {
		data[k] = v
	}
	c.Set(allKey, data)
	return
}

func Headers(c *gin.Context) (data map[string]string) {
	if ctxData := c.GetStringMapString(headersKey); ctxData != nil {
		return ctxData
	}
	data = make(map[string]string, 0)
	for k := range c.Request.Header {
		data[k] = c.Request.Header.Get(k)
	}
	c.Set(headersKey, data)
	return
}
