package console

import (

	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitHistoryRouter(group *gin.RouterGroup) {
	historyGroup := group.Group("history")
	{
		history := console.NewHistory()
		historyGroup.GET("", wrapper.Wrapper(history.List))
		historyGroup.POST("", wrapper.Wrapper(history.Post))
		historyGroup.GET("/:id", wrapper.Wrapper(history.Get))
		historyGroup.PUT("/:id", wrapper.Wrapper(history.Put))
		historyGroup.DELETE("/:id", wrapper.Wrapper(history.Delete))

	}
}