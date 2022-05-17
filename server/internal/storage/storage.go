package storage

import (
	"blog/internal/config"
	"blog/internal/storage/aliyun"
	"blog/internal/storage/local"
	"blog/internal/storage/qiniu"
	"blog/internal/storage/tencent"
	"mime/multipart"
)

type Storage interface {
	Save(header *multipart.FileHeader) (string,error)
	Delete(name ...string) (int, error)
	GetAccessUrl(name string) (string, error)
}

const (
	Local      = "local"
	AliyunOSS  = "aliyun-oss"
	TencentOSS = "tencent-oss"
	Qiniu      = "qiniu"
)

func New(setting *config.App) (Storage, error) {
	switch setting.AppStorageType {
	case Local:
		return local.New(setting)
	case AliyunOSS:
		return aliyun.New(setting.Storage.AliyunOSS)
	case TencentOSS:
		return tencent.New(setting.Storage.TencentOSS)
	case Qiniu:
		return qiniu.New(setting.Storage.Qiniu)
	default:
		return local.New(setting)
	}
}
