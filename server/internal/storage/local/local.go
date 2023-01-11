package local

import (
	"blog/internal/config"
	"blog/internal/storage/utils"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Local struct {
	Path      string
	host      string
	generator *utils.Generator
}

func New(setting *config.App) (*Local, error) {
	localPath := setting.Storage.Local.Path

	// 1、如果配置的存储路径是绝对路径的话，将路径加上homepath
	if !filepath.IsAbs(setting.Storage.Local.Path) {
		localPath = filepath.Join(setting.AppHomePath, setting.Storage.Local.Path)
	}
	// 2、判断是否存在该路径，不存在创建
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		if err := os.MkdirAll(localPath, os.ModePerm); err != nil {
			return nil, errors.New(fmt.Sprintf("mkdir local storage path failed: %s", err))
		}
	}

	return &Local{Path: localPath, host: setting.Storage.Local.Host, generator: utils.NewFilenameGenerator()}, nil
}

// Save 保存文件
// example： save(headers)
func (l *Local) Save(header *multipart.FileHeader) (string, error) {
	// 1、创建本地路径
	name, err := l.generator.NewName()
	if err != nil {
		return "",  err
	}
	dir := filepath.Join(l.Path, name.Year, name.Month, name.Day)

	// 2、判断路径是否存在，不存在则创建
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", errors.New(fmt.Sprintf("create local path: %s failed", dir))
		}
	}

	// 3、检查文件是否有写入权限
	if os.IsPermission(err) {
		return "",  errors.New("insufficient file permissions. ")
	}

	// 4、写入文件
	if err := saveLocal(filepath.Join(dir, name.Name+filepath.Ext(header.Filename)), header); err != nil {
		return "",  errors.New(fmt.Sprintf("failed flush to local path. "))
	}

	return name.Name + filepath.Ext(header.Filename), nil
}

// Delete
// ** 根据名称删除文件
// ** 1、解析名称获取文件时间戳
func (l *Local) Delete(name ...string) (int, error) {
	col := 0
	for _, s := range name {
		timestamp, _, err := l.generator.ParseName(s)
		if err != nil {
			log.Println(err)
			continue
		}
		unix := time.Unix(timestamp, 0)
		log.Println(filepath.Join(l.Path, strconv.FormatInt(int64(unix.Year()), 10), strconv.FormatInt(int64(unix.Month()), 10), strconv.FormatInt(int64(unix.Day()), 10), s))
		err = os.Remove(filepath.Join(l.Path, strconv.FormatInt(int64(unix.Year()), 10), strconv.FormatInt(int64(unix.Month()), 10), strconv.FormatInt(int64(unix.Day()), 10), s))
		if err == nil {
			log.Println(err)
			col += 1
		}
	}
	return col, nil
}

// GetAccessUrl
// ** 根据名称解析时间路径
func (l *Local) GetAccessUrl(name string) (string, error) {
	timestamp, _, err := l.generator.ParseName(strings.TrimSuffix(name, filepath.Ext(name)))
	if err != nil {
		return "", err
	}
	unix := time.Unix(timestamp, 0)

	return l.host + "/" + filepath.Join(l.Path, strconv.FormatInt(int64(unix.Year()), 10), strconv.FormatInt(int64(unix.Month()), 10), strconv.FormatInt(int64(unix.Day()), 10), name), nil
}

func saveLocal(p string, header *multipart.FileHeader) error {
	src, err := header.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		_ = src.Close()
	}(src)

	out, err := os.Create(p)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)
	_, err = io.Copy(out, src)
	return err
}
