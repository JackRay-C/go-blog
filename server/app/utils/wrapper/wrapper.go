package wrapper

import (
	"blog/app/response"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Wrapper(f func(c *gin.Context)(*response.Response, error)) gin.HandlerFunc  {
	return func(c *gin.Context) {
		resp, err := f(c)
		if err == nil {
			c.JSON(http.StatusOK, resp)
			return
		}
		c.JSON(http.StatusOK, err)
	}
}
