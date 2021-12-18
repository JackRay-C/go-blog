package dto

type ListSubjects struct {
	PageNo     int    `json:"page_no" form:"page_no"`
	PageSize   int    `json:"page_size" form:"page_size"`
	UserId     int    `json:"user_id" form:"user_id"`
	Visibility int    `json:"visibility" form:"visibility"`
	Search     string `json:"search" form:"search"`
}

type AddSubjects struct {
	Title        string `json:"title"`
	AvatarId     int    `json:"avatar_id"`
	CoverImageId int    `json:"cover_image_id"`
	Description  string `json:"description"`
	Visibility   int    `json:"visibility"`
}

type PutSubjects struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	AvatarId     int    `json:"avatar_id"`
	CoverImageId int    `json:"cover_image_id"`
	Description  string `json:"description"`
	Visibility   int   `json:"visibility"`
	Views        int    `json:"views"`
}
