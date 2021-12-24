package web

import (
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
)

type Comment struct {
	log         logger.Logger
	postService *service.PostService
}

func NewComment() *Post {
	return &Post{
		log:         global.Logger,
		postService: service.NewPostService(),
	}
}


// 根据post id 获取comment
func (c *Comment) List()  {

}

// 获取单条评论
func (c *Comment) Get()  {

}