package context

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	getAllKey  = "__getAll__"
	postAllKey = "__postAll__"
	jsonAllKey = "__jsonAll__"
	allKey     = "__all__"
	headersKey = "__headers__"
)

type HandlerFunc func(*Context)

func Handler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := new(Context)
		ctx.Context = c
		_ = c.ShouldBindWith(&ctx.GlobalHeader, binding.Header)
		handler(ctx)
	}
}

type Context struct {
	*gin.Context
	GlobalHeader struct {
		Authorization    string `header:"Authorization"`       // 身份凭证
		VersionCode      int64  `header:"X-Version-Code"`      // 应用版本
		VersionName      string `header:"X-Version-Name"`      // 应用版本名称
		ApiVersion       string `header:"X-API-Version"`       // API版本
		MacCode          string `header:"X-Mac-Code"`          // Mac地址
		OsVersion        string `header:"X-OS-Version"`        // 系统版本
		OsName           string `header:"X-OS-Name"`           // 系统名称
		ResolutionWidth  string `header:"X-Resolution-Width"`  // 分辨率-宽
		ResolutionHeight int64  `header:"X-Resolution-Height"` // 分辨率-高
		PlatformID       int64  `header:"X-Platform-ID"`       // 平台ID
		PlatformName     string `header:"X-Platform-Name"`     // 平台名称
		ChannelID        int64  `header:"X-Channel-ID"`        // 渠道ID
		ChannelName      string `header:"X-Channel-Name"`      // 渠道名称
	}
}

func (c *Context) GetAll() (data map[string]interface{}) {
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

func (c *Context) PostAll() (data map[string]interface{}) {
	if ctxData := c.GetStringMap(postAllKey); ctxData != nil {
		return ctxData
	}
	data = make(map[string]interface{}, 0)
	if c.Request.Method == http.MethodPost {
		if c.ContentType() == "application/json" {
			return c.JsonAll()
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

func (c *Context) JsonAll() (data map[string]interface{}) {
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

func (c *Context) All() (data map[string]interface{}) {
	if ctxData := c.GetStringMap(allKey); ctxData != nil {
		return ctxData
	}
	data = make(map[string]interface{}, 0)
	for k, v := range c.GetAll() {
		data[k] = v
	}
	for k, v := range c.JsonAll() {
		data[k] = v
	}
	for k, v := range c.PostAll() {
		data[k] = v
	}
	c.Set(allKey, data)
	return
}

func (c *Context) Headers() (data map[string]string) {
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

func (c *Context) ResponseWriter() *responseWriter {
	return &responseWriter{
		ResponseWriter: c.Writer,
		Body:           bytes.NewBufferString(""),
	}
}
