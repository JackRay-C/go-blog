package tencent

import (
	"blog/internal/config"
	"mime/multipart"
)

type TecentOSS struct {

}

func New(setting *config.TencentOSS) (*TecentOSS, error)  {
	return &TecentOSS{}, nil
}

func (t *TecentOSS) Save(header *multipart.FileHeader) (string, error) {
	panic("implement me")
}

func (t *TecentOSS) Delete(name ...string) (int, error) {
	panic("implement me")
}

func (t *TecentOSS) GetAccessUrl( name string) (string, error){
	panic("implement me")
}
