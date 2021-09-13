package gin

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/url"
	"skeleton/request"
)

type Context struct {
	Gin      *gin.Context
	body     map[string]interface{}
	getBody  map[string]interface{}
	postBody map[string]interface{}
	jsonBody map[string]interface{}
	header   request.Header
	isHeader bool
}

func (c *Context) Get() map[string]interface{} {
	c.getBody = make(map[string]interface{}, 0)

	if c.Gin.Request.Method == "GET" {
		if c.Gin.Request.URL.Query() != nil {
			for k, v := range c.Gin.Request.URL.Query() {
				c.getBody[k] = v[0]
			}
		}
	}

	return c.getBody
}

func (c *Context) Post() map[string]interface{} {
	c.postBody = make(map[string]interface{}, 0)

	if c.Gin.Request.Method == "POST" {
		if c.Gin.ContentType() == "application/json" {
			jsonBody := c.Json()
			for k, v := range jsonBody {
				c.postBody[k] = v
			}
			return c.postBody
		}

		if c.Gin.Request.ParseForm() == nil {
			_ = c.Gin.Request.ParseMultipartForm(32 << 20)
		}

		formData := make(url.Values)
		switch c.Gin.ContentType() {
		case "multipart/form-data":
			formData = c.Gin.Request.PostForm
		case "application/x-www-form-urlencoded":
			formData = c.Gin.Request.Form
		}
		if formData != nil {
			for k, v := range formData {
				c.postBody[k] = v[0]
			}
		}
	}

	return c.postBody
}

func (c *Context) Json() map[string]interface{} {
	c.jsonBody = make(map[string]interface{}, 0)

	if c.Gin.ContentType() == "application/json" {
		var b []byte
		if cb, ok := c.Gin.Get(gin.BodyBytesKey); ok {
			if cbb, ok := cb.([]byte); ok {
				b = cbb
			}
		}
		if b == nil {
			b, _ = c.Gin.GetRawData()
			c.Gin.Set(gin.BodyBytesKey, b)
			c.Gin.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		}
		var body map[string]interface{}
		_ = json.Unmarshal(b, &body)
		for k, v := range body {
			c.jsonBody[k] = v
		}
	}

	return c.jsonBody
}

func (c *Context) All() map[string]interface{} {
	c.body = make(map[string]interface{}, 0)

	for k, v := range c.Get() {
		c.body[k] = v
	}
	for k, v := range c.Json() {
		c.body[k] = v
	}
	for k, v := range c.Post() {
		c.body[k] = v
	}

	return c.body
}

func (c *Context) Header() request.Header {
	if !c.isHeader {
		c.Gin.ShouldBindHeader(&c.header)
		c.isHeader = true
	}
	return c.header
}
