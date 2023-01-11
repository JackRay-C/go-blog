package service

import (
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service/impl"
	"github.com/gin-gonic/gin"
)

type PostsTagsService interface {
	ISelectTagsByPost(c *gin.Context, tags *[]*po.Tag, post *po.Post) error
	ISelectPostsByTag(c *gin.Context, pager *vo.Pager, tag *po.Tag) error
}

func NewPostsTagsService() PostsTagsService {
	return &impl.PostsTagsServiceImpl{}
}