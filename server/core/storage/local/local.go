package local

import (
	"blog/app/encrypt"
	"blog/core/setting"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type Local struct {
	*setting.App
	*setting.Local
}

func (l *Local) GetAccessUrl(prefix, name string) string {
	return path.Join(l.Path, prefix, name)
}

func NewLocalStorage(appSetting *setting.App, localSetting *setting.Local) *Local {
	return &Local{
		App:   appSetting,
		Local: localSetting,
	}
}

func (l *Local) Save(p string, header *multipart.FileHeader) error {
	name := path.Base(p)
	dir := path.Join(l.Path, strings.Split(p, name)[0])

	// 3、检查本地路径，没有则创建
	if checkLocalPath(dir) {
		if err := createLocalPath(dir, os.ModePerm); err != nil {
			return errors.New("failed to create local directory. ")
		}
	}

	// 4、检查是否有文件写入权限
	if checkPermission(dir) {
		return errors.New("insufficient file permissions. ")
	}

	// 6、写入本地路径
	err := l.saveLocal(path.Join(dir, name), header)
	if err != nil {
		return errors.New("failed to flush to local path. ")
	}

	// 7、文件信息
	return nil
}

func (l *Local) saveLocal(p string, header *multipart.FileHeader) error {
	src, err := header.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(p)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}

func generateFileName(header *multipart.FileHeader) string {
	ext := path.Ext(header.Filename)
	name := strings.TrimSuffix(header.Filename, ext)
	return encrypt.MD5(name) + ext
}

func checkLocalPath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}
func createLocalPath(dst string, perm os.FileMode) error {
	return os.MkdirAll(dst, perm)
}
func checkPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func (l *Local) Delete(name ...string) (int, error) {
	// 1、获取文件类型及前缀
	col := 0
	for _, s := range name {
		err := os.Remove(s)
		if err == nil {
			col += 1
		}
	}
	return col, nil
}
