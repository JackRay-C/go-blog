package tencent_oss

import (
	"blog/core/setting"
	"io/ioutil"
	"mime/multipart"
	"path"
	"strings"
)

type TencentOss struct {
	*setting.App
	*setting.TencentOSS
}



func (t *TencentOss) Save(p string, header *multipart.FileHeader) error {
	panic("implement me")
}

func (t *TencentOss) Delete(name ...string) (int, error) {
	panic("implement me")
}

func (t *TencentOss) CheckExts(filename string) bool {
	for _, s := range  t.UploadAllowExts{
		if strings.ToUpper(s) == strings.ToUpper(path.Ext(filename)) {
			return true
		}
	}
	return false
}

func (t *TencentOss) CheckSize(file multipart.File) bool {
	content, _ := ioutil.ReadAll(file)
	fileSize := len(content)

	if int64(fileSize) <= int64(t.UploadMaxSize) {
		return true
	}

	return false
}
func (t *TencentOss) GetAccessUrl(prefix, name string) string {
	return "/" + prefix + "/" + name;
}