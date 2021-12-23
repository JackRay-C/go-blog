package vo

import (
	"blog/app/domain"
	"time"
)

type VTag struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	UserId       int          `json:"user_id"`
	CoverImageId int          `json:"cover_image_id"`
	CoverImage   *domain.File `json:"cover_image"`
	CreatedAt    time.Time    `json:"created_at"`
}
