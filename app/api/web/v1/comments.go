package v1

import (
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
)

type Comment struct {
	log            logger.Logger
	postService    service.PostService
	commentService service.CommentService
	userService *service.UserService
}

func NewComment() *Comment {
	return &Comment{
		log:            global.Logger,
		postService:    service.NewPostService(),
		commentService: service.NewCommentService(),
		userService: service.NewUserService(),
	}
}


// Get 获取一条评论
func (c *Comment) Get(ctx *gin.Context) (*response.Response, error) {
	panic("implement me")
}

// List 根据博客ID获取所有评论
func (c *Comment) List(ctx *gin.Context) (*response.Response, error) {
	panic("implement me")
}
