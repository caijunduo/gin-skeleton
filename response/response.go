package response

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type response struct {
	status  int
	code    int
	message string
	data    interface{}
	err     error

	temp map[string]interface{}
}

func (r *response) GetStatus() int {
	return r.status
}

func (r *response) GetCode() int {
	return r.code
}

func (r *response) GetMessage() string {
	return cast.ToString(r.getTemp("message", r.message))
}

func (r *response) GetData() interface{} {
	return r.getTemp("data", r.data)
}

func (r *response) GetError() error {
	if err, ok := r.getTemp("err", r.err).(error); ok && err != nil {
		return err
	}
	return nil
}

func (r *response) SetMessage(message string) *response {
	r.setTemp("message", message)
	return r
}

func (r *response) SetData(data interface{}) *response {
	r.setTemp("data", data)
	return r
}

func (r *response) SetError(err error) *response {
	r.setTemp("err", err)
	return r
}

func (r *response) setTemp(key string, value interface{}) {
	if r.temp == nil {
		r.resetTemp()
	}
	r.temp[key] = value
}

func (r *response) getTemp(key string, def interface{}) interface{} {
	if val, ok := r.temp[key]; ok {
		return val
	} else {
		return def
	}
}

func (r *response) resetTemp() {
	r.temp = make(map[string]interface{}, 0)
}

func (r *response) Slice() (int, interface{}) {
	res := gin.H{
		"code":    r.code,
		"message": r.GetMessage(),
	}
	if data := r.GetData(); data != nil {
		res["data"] = data
	}
	if gin.Mode() != gin.ReleaseMode {
		if err := r.GetError(); err != nil {
			res["error"] = err
		}
	}
	r.resetTemp()
	return r.status, res
}
