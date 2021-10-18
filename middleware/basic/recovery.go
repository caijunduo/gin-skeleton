package basicMiddleware

import (
	"github.com/gin-gonic/gin"
	"skeleton/response"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rc := recover(); rc != nil {
				//var stack string
				//if pc, f, fl, ok := runtime.Caller(2); ok {
				//	stack = fmt.Sprintf("(%s:%d): %s", runtime.FuncForPC(pc).Name(), fl, f)
				//}
				switch v := rc.(type) {
				default:
					c.AbortWithStatusJSON(response.InternalServerError.SetError(v.(error)).Slice())
				}
			}
		}()

		c.Next()
	}
}
