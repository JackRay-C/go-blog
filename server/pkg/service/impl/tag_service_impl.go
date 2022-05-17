package impl

import (
	"blog/pkg/model/common"
)

type TagServiceImpl struct {
	common.BaseServiceImpl
}


//func (t *TagServiceImpl) ISelectOne(c *gin.Context, tag *po.Tag) error {
//	if err := global.DB.Model(&po.Tag{}).Where("id=? and user_id=?", tag.ID, tag.UserId).Error; err != nil {
//		return vo.DatabaseSelectError.SetMsg("%s", err)
//	}
//	return nil
//}
//
//func (t *TagServiceImpl) ISelectList(c *gin.Context, pager *vo.Pager, tag *po.Tag) error {
//	var tags []*po.Tag
//	offset := (pager.PageNo - 1) *pager.PageSize
//	limit := pager.PageSize
//
//	db := global.DB.Model(&po.Tag{}).Where("user_id=?", tag.UserId)
//
//	if err := db.Count(&pager.TotalRows).Error; err != nil {
//		return vo.DatabaseSelectError.SetMsg("%s", err)
//	}
//
//	if err := db.Offset(offset).Limit(limit).Find(&tags).Error; err != nil {
//		return vo.DatabaseSelectError.SetMsg("%s", err)
//	}
//
//	pager.PageCount = int((pager.TotalRows + int64(pager.PageSize) - 1) / int64(pager.PageSize))
//	pager.MustList(tags)
//
//	return nil
//}
//
//func (t *TagServiceImpl) IUpdateOne(c *gin.Context, tag *po.Tag, tags *dto.PutTags) error {
//
//	if tag.ID != tags.ID {
//		return vo.InvalidParams.SetMsg("ID [%d] miss match update tag ID [%d]", tag.ID, tags.ID)
//	}
//
//	if err := transform.Transition(tags, tag); err != nil {
//		return vo.TransformFailed.SetMsg("%t transform failed to %t: %s", tags, tag, err)
//	}
//
//	if err := global.DB.Model(&po.Tag{}).Where("id=? and user_id=?", tag.ID, tag.UserId).Updates(tags).Error; err != nil {
//		return vo.DatabaseUpdateError.SetMsg("%s", err)
//	}
//
//	return nil
//}
//
//func (t *TagServiceImpl) IDeleteOne(c *gin.Context, tag *po.Tag) error {
//	panic("implement me")
//}
//
//func (t *TagServiceImpl) ICreateOne(c *gin.Context, tag *po.Tag) error {
//	panic("implement me")
//}
//
//func (t *TagServiceImpl) ISelectOneWeb(c *gin.Context, tag *po.Tag) error {
//	panic("implement me")
//}
//
//func (t *TagServiceImpl) ISelectListWeb(c *gin.Context, pager *vo.Pager, tag *po.Tag) error {
//	panic("implement me")
//}
