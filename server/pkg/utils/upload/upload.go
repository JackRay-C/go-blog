package upload

import (
	"blog/pkg/global"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func GetFileExt(filename string) string  {
	return path.Ext(filename)
}

func CheckMaxSize(file multipart.File) bool  {
	content, _ := ioutil.ReadAll(file)
	l := len(content)
	if int64(l) >= global.App.Storage.AllowUploadMaxSize.Size() {
		return false
	}
	return true
}

func CheckExts(filename string) bool  {
	ext := path.Ext(filename)
	for _, allowExt := range global.App.Storage.AllowUploadFileExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckPathExist(p string) bool  {
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		return false
	}
	return true
}