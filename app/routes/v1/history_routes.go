package v1

import (
	v1 "blog/app/api/console/v1"
	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitHistoryRouter(group *gin.RouterGroup) {
	historyGroup := group.Group("history")
	{
		history := v1.NewHistory()
		historyGroup.GET("", wrapper.Wrapper(history.List))
		historyGroup.POST("", wrapper.Wrapper(history.Post))
		historyGroup.GET("/:id", wrapper.Wrapper(history.Get))
		historyGroup.PUT("/:id", wrapper.Wrapper(history.Put))
		historyGroup.DELETE("/:id", wrapper.Wrapper(history.Delete))

	}
}