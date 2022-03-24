package po

import "encoding/json"

type HeadsTags struct {
	HeadId int `json:"head_id" gorm:"type:int;index:idx_headId_tagId,unique;common:head id"`
	TagId  int `json:"tag_id" gorm:"type:int;index:idx_headId_tagId,unique;common:tag id"`
}

func (p *HeadsTags) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (*HeadsTags) TableName() string {
	return "heads_tags"
}
