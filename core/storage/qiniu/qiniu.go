package qiniu

import (
	"blog/core/setting"
	"io/ioutil"
	"mime/multipart"
	"path"
	"strings"
)

type Qiniu struct {
	*setting.App
	*setting.Qiniu
}

func (q *Qiniu) Save(p string, header *multipart.FileHeader) error {
	panic("implement me")
}

func (q *Qiniu) Delete(name ...string) (int, error) {
	panic("implement me")
}

func (q *Qiniu) CheckExts(filename string) bool {
	for _, s := range  q.UploadAllowExts{
		if strings.ToUpper(s) == strings.ToUpper(path.Ext(filename)) {
			return true
		}
	}
	return false
}

func (q *Qiniu) CheckSize(file multipart.File,) bool {
	content, _ := ioutil.ReadAll(file)
	fileSize := len(content)

	if int64(fileSize) <= int64(q.UploadMaxSize) {
		return true
	}

	return false
}

func (q *Qiniu) GetAccessUrl(prefix, name string) string {
	return q.Bucket + "/" + prefix + "/" + name;
}