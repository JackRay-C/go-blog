package qiniu

import (
	"blog/internal/config"
	"mime/multipart"
)

type Qiniu struct {

}

func New(setting *config.Qiniu) (*Qiniu, error)  {
	return &Qiniu{}, nil
}

func (q *Qiniu) Save(header *multipart.FileHeader) (string,error) {
	panic("implement me")
}

func (q *Qiniu) Delete(name ...string) (int, error) {
	panic("implement me")
}

func (q *Qiniu) GetAccessUrl(name string)(string, error){
	panic("implement me")
}

