package storage

import (
	"encoding/json"
	"mime/multipart"
)

type FileInfo struct {
	FileName string // 文件名
	Path string // 文件路径
	FileType FileType // 文件类型
	Ext string // 文件后缀
	AccessUrl string
}

func (f FileInfo) String() string {
	marshal, err := json.Marshal(f)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type Storage interface {
	Save(p string, header *multipart.FileHeader) error
	Delete(name ...string) (int, error)
	GetAccessUrl(prefix, name string) string
}
