package impl

import (
	"blog/pkg/global"
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"github.com/gin-gonic/gin"
)

type CommentServiceImpl struct {
	common.BaseServiceImpl
}


func (s *CommentServiceImpl) ISelectAllByPostId(ctx *gin.Context, comments *[]*po.Comment) error {
	postId, ok := ctx.GetQuery("post_id")
	if !ok {
		return vo.InvalidParams.SetMsg("this query param post_id is required. ")
	}

	var post *po.Post
	if err := global.DB.Model(&po.Post{}).Where("id=?", postId).First(&post).Error; err != nil {
		return vo.RecordNotFound.SetMsg("not found this post: %s", postId)
	}

	var all []*po.Comment
	if err := global.DB.Model(&po.Comment{}).Where("post_id=?", postId).Find(&all).Error; err != nil {
		return vo.DatabaseSelectError
	}

	m := make(map[int64]*po.Comment)
	for _, comment := range all {
		if comment.ParentID == 0 {
			*comments = append(*comments, comment)
		}
		m[comment.ID] = comment
	}

	for _, comment := range all {
		if comment.ParentID!= 0 {
			if c, ok := m[comment.ParentID]; ok {
				c.Child = append(c.Child, comment)
			}
		}
	}
	return nil
}
