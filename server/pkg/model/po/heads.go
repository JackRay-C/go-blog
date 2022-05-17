package po

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Head struct {
	ID           int            `json:"id" gorm:"type:int;primary_key;autoIncrement;common:主键ID;" form:"id"` // 主键ID，博客的唯一ID
	RepositoryID int            `json:"repository_id" gorm:"type:int" form:"repository_id"`                  // 当前存储库的id
	Visibility   int            `json:"visibility" gorm:"type:tinyint(2)" form:"visibility"`                 // 是否公开 1、私有 2、公开
	Status       int            `json:"status" gorm:"type:tinyint(2)" form:"status"`                         // 博客状态 1、已暂存 2、已提交 3、已发布
	SubjectID    int            `json:"subject_id" gorm:"type:int;index:idx_subject_id" form:"subject_id"`   // 专题ID
	CoverImageId int            `json:"cover_image_id"`                                                      // 封面图片ID
	Likes        int            `json:"likes"  gorm:"type:int"`                                              // 点赞
	Views        int            `json:"views" gorm:"type:int"`                                               // 阅读量
	UserID       int            `json:"user_id" gorm:"type:int;index:idx_user_id" form:"user_id"`            // 用户ID
	CreatedAt    time.Time      `json:"created_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"`        // 创建时间
	UpdatedAt    time.Time      `json:"updated_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"`        // 更新时间
	DeletedAt    gorm.DeletedAt `json:"deleted_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"`        // 删除时间
}

func (h *Head) String() string {
	marshal, err := json.Marshal(h)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (*Head) TableName() string {
	return "heads"
}
