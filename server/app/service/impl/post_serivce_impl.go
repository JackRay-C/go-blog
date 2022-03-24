package impl

import (
	"blog/app/model/po"
	"blog/app/model/vo"
	"blog/app/utils/auth"
	"blog/core/global"
	"github.com/gin-gonic/gin"
)

type PostServiceImpl struct {
}



func (p *PostServiceImpl) ISelectAllWeb(c *gin.Context, page *vo.Pager, post *po.Post) error {
	var posts []*po.Post
	db := global.DB.Model(&po.Post{})

	if !auth.CheckLogin(c) {
		db.Where("visibility=2")
	} else {
		db.Where("(user_id=? and visibility=1) or (visibility=2)")
	}

	if err := db.Count(&page.TotalRows).Error; err != nil {
		return err
	}

	if err := db.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize).Find(&posts).Error; err != nil {
		return err
	}

	page.MustList(&posts)

	return nil
}

func (p *PostServiceImpl) ISelectOneWeb(c *gin.Context, post *po.Post) error {
	panic("implement me")
}

func (p *PostServiceImpl) IUpdateOneWeb(c *gin.Context, post *po.Post) error {
	panic("implement me")
}

func (p *PostServiceImpl) ISearchWeb(c *gin.Context, page *vo.Pager) error {
	panic("implement me")
}

func (p *PostServiceImpl) IStaged(c *gin.Context, post *po.Post) error {
	panic("implement me")
}

func (p *PostServiceImpl) ICommit(c *gin.Context, post *po.Post) error {
	panic("implement me")
}

func (p *PostServiceImpl) IPublish(c *gin.Context, post *po.Post) error {
	panic("implement me")
}

func (p *PostServiceImpl) IPull(c *gin.Context, post *po.Post) error {
	panic("implement me")
}

func (p *PostServiceImpl) ISelectAll(c *gin.Context, page *vo.Pager) error {
	panic("implement me")
}

func (p *PostServiceImpl) ISelectOne(c *gin.Context, post *po.Post) error {
	panic("implement me")
}

func (p *PostServiceImpl) ICreateOne(c *gin.Context, post *po.Post) error {
	panic("implement me")
}

func (p *PostServiceImpl) IUpdateOne(c *gin.Context, post *po.Post) error {
	panic("implement me")
}

func (p *PostServiceImpl) IDeleteOne(c *gin.Context, post *po.Post) error {
	panic("implement me")
}

