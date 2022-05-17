package wrapper

import (

	"blog/pkg/model/vo"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)


func Wrapper(f func(c *gin.Context)(*vo.Response, error)) gin.HandlerFunc  {
	return func(c *gin.Context) {
		resp, err := f(c)
		if err == nil {
			c.JSON(http.StatusOK, resp)
			return
		}
		of := reflect.TypeOf(err)
		
		if of.Kind() == reflect.Ptr {
			of = of.Elem()
		}
		if of.Name() == "vo.Error" {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 500, "message": err})
		}
	}
}
