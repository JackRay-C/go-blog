package impl

import (
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"github.com/gin-gonic/gin"
)

type CommentServiceImpl struct {

}

func (s *CommentServiceImpl) ISelectOne(c *gin.Context, comment *po.Comment) error {
	return nil
}

func (s *CommentServiceImpl) ISelectAll(c *gin.Context, pager *vo.Pager, comment *po.Comment) error {
	panic("implement me")
}

func (s *CommentServiceImpl) ICreate(c *gin.Context, comment *po.Comment) error {
	panic("implement me")
}

func (s *CommentServiceImpl) IDelete(c *gin.Context, comment *po.Comment) error {
	panic("implement me")
}

func (s *CommentServiceImpl) ISelectAllByPostId(ctx *gin.Context) error {
	panic("implement me")
}



