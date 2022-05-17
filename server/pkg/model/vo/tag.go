package vo

import (
	"blog/pkg/model/po"
	"time"
)

type VTag struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	UserId       int          `json:"user_id"`
	CoverImageId int          `json:"cover_image_id"`
	CoverImage   *po.File `json:"cover_image"`
	CreatedAt    time.Time    `json:"created_at"`
}
