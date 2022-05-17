package impl

import (
	"blog/pkg/model/common"
)

type DictServiceImpl struct {
	common.BaseServiceImpl
}

//func (d *DictServiceImpl) IUpdateOne(c *gin.Context, dict *po.Dict) error {
//	if !auth.CheckLogin(c) {
//		return vo.NotLogin
//	}
//	if !auth.CheckPermission(c, "dicts", "update") {
//		return vo.Forbidden
//	}
//
//	if err := global.DB.Model(&po.Dict{}).Where("id=?", dict.ID).Updates(dict).Error; err != nil {
//		return vo.DatabaseUpdateError
//	}
//	return nil
//}
//
//func (d *DictServiceImpl) ICreateOne(c *gin.Context, dict *po.Dict) error {
//	if !auth.CheckLogin(c) {
//		return vo.NotLogin
//	}
//
//	if !auth.CheckPermission(c, "dicts", "add") {
//		return vo.Forbidden
//	}
//
//	// todo: 判断是否有同名的dict
//	if err := global.DB.Model(&po.Dict{}).Create(dict).Error; err != nil {
//		return vo.DatabaseInsertError
//	}
//
//	return nil
//}
//
//func (d *DictServiceImpl) IDeleteOne(c *gin.Context, dict *po.Dict) error {
//	if !auth.CheckLogin(c) {
//		return vo.NotLogin
//	}
//
//	if !auth.CheckPermission(c, "dicts", "add") {
//		return vo.Forbidden
//	}
//
//	if err := global.DB.Model(&po.Dict{}).Where("id=?", dict.ID).Delete(dict).Error; err != nil {
//		return vo.DatabaseDeleteError
//	}
//	return nil
//}
//
//func (d *DictServiceImpl) ISelectOne(c *gin.Context, dict *po.Dict) error {
//	db := global.DB.Model(&po.Dict{})
//	if dict.ID != 0 {
//		db.Where("id=? ", dict.ID)
//	}
//	if dict.Name != "" {
//		db.Where("name=?", dict.Name)
//	}
//	return db.First(&dict).Error
//}
//
//func (d *DictServiceImpl) ISelectAll(c *gin.Context, pager *vo.Pager, dict *po.Dict) error {
//	offset := (pager.PageNo - 1) * pager.PageSize
//	limit := pager.PageSize
//	var dicts []*po.Dict
//	var count int64
//
//
//	db := global.DB.Model(&po.Dict{})
//	if dict.Name != "" {
//		db.Where("name=?", dict.Name)
//	}
//
//	if err := db.Count(&count).Error; err != nil {
//		return vo.DatabaseSelectError
//	}
//
//	if err := db.Offset(offset).Limit(limit).Find(&dicts).Error; err != nil {
//		return vo.DatabaseSelectError
//	}
//
//	global.Log.Infof("%s", dicts)
//	pager.TotalRows = count
//	pager.PageCount  = int( (count+int64(pager.PageSize) - 1)/ int64(pager.PageSize))
//	pager.MustList(&dicts)
//
//	return nil
//}

