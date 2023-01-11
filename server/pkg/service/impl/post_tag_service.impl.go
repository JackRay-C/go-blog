package impl

import (
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"github.com/gin-gonic/gin"
)

type PostsTagsServiceImpl struct{}

func (p *PostsTagsServiceImpl) ISelectTagsByPost(c *gin.Context, tags *[]*po.Tag, post *po.Post) error {
	// 通过post id 获取 tag 列表
	if err := global.DB.Table("tags").Joins("left join posts_tags on tags.id=posts_tags.tag_id").Where("posts_tags.post_id=?", post.ID).Find(&tags).Error; err != nil {
		return err
	}

	return nil
}

func (p *PostsTagsServiceImpl) ISelectPostsByTag(c *gin.Context, pager *vo.Pager, tag *po.Tag) error {
	// 通过tagID 获取post 列表
	db := global.DB.Table("posts").Joins("left join posts_tags on posts.id=posts_tags.posts_id").Where("posts_tags.tag_id=?", tag.ID)

	if err := db.Count(&pager.TotalRows).Error; err != nil {
		return err
	}

	offset := (pager.PageNo - 1) * pager.PageSize
	limit := pager.PageSize
	var posts []*po.Post
	if err := db.Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return err
	}
	pager.MustList(posts)
	return nil
}
