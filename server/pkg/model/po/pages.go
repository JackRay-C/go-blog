package po

type Pages struct {
	ID int64 `json:"id" gorm:"type:int;primary_key;autoIncrement;common:主键ID;" form:"id"`
	Slug string `json:"slug" gorm:"type:varchar(255);index:idx_slug,unique" form:"slug"`
	PostId int64 `json:"post_id"`
	Name string `json:"name"`
}
