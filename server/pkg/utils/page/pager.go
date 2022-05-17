package page

import (
	"blog/pkg/utils/convert"
	"github.com/gin-gonic/gin"
)

func GetPageNo(c *gin.Context) int {
	pageNo := c.DefaultQuery("page_no", "1")
	return convert.StrTo(pageNo).MustInt()
}

func GetPageSize(c *gin.Context) int {
	pageSize := c.DefaultQuery("page_size", "10")
	return convert.StrTo(pageSize).MustInt()
}

