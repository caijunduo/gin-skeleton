package accountInternal

import (
	"github.com/gin-gonic/gin"
	"skeleton/response"
)

func GetInfo(c *gin.Context) {
	c.IndentedJSON(response.OK.SetData("Hello，Skeleton.").Slice())
}

func GetMessageInfo(c *gin.Context) {
	c.IndentedJSON(response.OK.SetData("Hello, SKeleton Message.").Slice())
}
