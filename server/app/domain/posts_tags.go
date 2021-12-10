package domain

import (
	"blog/core/global"
	"encoding/json"
	"gorm.io/gorm/clause"
)

type PostsTags struct {
	PostId int `json:"post_id" gorm:"type:int;index:idx_postId_tagId,unique;comment:post id"`
	TagId  int `json:"tag_id" gorm:"type:int;index:idx_postId_tagId,unique;comment:tag id"`
}

func (p *PostsTags) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (*PostsTags) TableName() string {
	return "posts_tags"
}

// 插入所有
func (p *PostsTags) InsertAll(postsTags []PostsTags) error {
	return global.DB.Model(p).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(postsTags, 1000).Error
}

func (p *PostsTags) Insert() error  {
	return global.DB.Model(p).Clauses(clause.OnConflict{DoNothing: true}).Create(p).Error
}
func (p *PostsTags) Delete() error  {
	return global.DB.Model(p).Where("post_id=? and tag_id=?", p.PostId, p.TagId).Delete(p).Error
}

// 根据tag id 删除所有关系
func (p *PostsTags) DeleteByTags(tagId int64) error {
	return global.DB.Model(p).Where("tag_id=?", tagId).Delete(p).Error
}

func (p *PostsTags) DeleteByPosts(id int) error {
	return global.DB.Model(p).Where("post_id=?", id).Delete(p).Error
}

