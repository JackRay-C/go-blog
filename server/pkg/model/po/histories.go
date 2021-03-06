package po

import (
	"database/sql"
	"encoding/json"
)

type History struct {
	HeadID           int          `json:"head_id" gorm:"type:int;index:idx_head_id" form:"head_id" binding:"required"`
	RepositoryID     int          `json:"repository_id" gorm:"type:int;index:idx_repository_id"`
	PrevRepositoryID int          `json:"prev_repository_id"`
	UserID           int          `json:"user_id"`
	StagedAt         sql.NullTime `json:"staged_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"`
	CommitedAt       sql.NullTime `json:"commited_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"`
	PublishedAt      sql.NullTime `json:"published_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"`
}

func (r *History) String() string {
	marshal, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (*History) TableName() string {
	return "histories"
}
