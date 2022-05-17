package v1

import (
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
)

type Post struct {
	log            logger.Logger
	postService    service.PostService
	commentService service.CommentService
}

func NewPost() *Post {
	return &Post{
		log:            global.Logger,
		postService:    service.NewPostService(),
		commentService: service.NewCommentService(),
	}
}

// Get 获取单个博客信息
func (p *Post) Get(c *gin.Context) (*response.Response, error) {
	panic("implement me")
}

// List 分页获取所有博客信息
func (p *Post) List(c *gin.Context) (*response.Response, error) {
	panic("implement me")
}
