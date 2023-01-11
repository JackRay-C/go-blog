package po

import "encoding/json"

type PostsTags struct {
	PostId int64 `json:"post_id" gorm:"type:int;index:idx_postId_tagId,unique;common:head id"`
	TagId  int64 `json:"tag_id" gorm:"type:int;index:idx_postId_tagId,unique;common:tag id"`
}

func (p *PostsTags) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (*PostsTags) TableName() string {
	return "heads_tags"
}
