package console



import (

	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitDraftRouter(group *gin.RouterGroup) {
	draftGroup := group.Group("drafts")
	{
		draft := console.NewDraft()
		draftGroup.GET("", wrapper.Wrapper(draft.List))
		draftGroup.POST("", wrapper.Wrapper(draft.Post))
		draftGroup.GET("/:id", wrapper.Wrapper(draft.Get))
		draftGroup.PUT("/:id", wrapper.Wrapper(draft.Put))
		draftGroup.POST("/:id/published", wrapper.Wrapper(draft.Get))
		draftGroup.DELETE("/:id", wrapper.Wrapper(draft.Delete))
	}
}
