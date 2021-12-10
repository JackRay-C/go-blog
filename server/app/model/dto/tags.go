package dto

import "encoding/json"

type AddTags struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CoverImage  int    `json:"cover_image"`
}

func (l *AddTags) String() string {
	marshal, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type PutTags struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CoverImage  int    `json:"cover_image"`
}

func (l *PutTags) String() string {
	marshal, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(marshal)
}
