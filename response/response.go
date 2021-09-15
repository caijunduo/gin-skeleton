package response

import "github.com/gin-gonic/gin"

type Response struct {
	status  int
	code    int
	message string
	data    interface{}
	err     error
}

func (r *Response) GetStatus() int {
	return r.status
}

func (r *Response) GetCode() int {
	return r.code
}

func (r *Response) GetMessage() string {
	return r.message
}

func (r *Response) GetData() interface{} {
	return r.data
}

func (r *Response) GetError() error {
	return r.err
}

func (r *Response) SetMessage(message string) *Response {
	r.message = message
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.data = data
	return r
}

func (r *Response) SetError(err error) *Response {
	r.err = err
	return r
}

func (r *Response) Slice() (int, interface{}) {
	res := gin.H{
		"code":    r.code,
		"message": r.message,
	}
	if r.data != nil {
		res["data"] = r.data
	}
	if gin.Mode() != gin.ReleaseMode && r.err != nil {
		res["error"] = r.err.Error()
	}
	return r.status, res
}
