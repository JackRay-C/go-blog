package aliyun

import (
	"blog/internal/config"
	"blog/internal/storage/utils"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"path"
	"strconv"
	"time"
)

type AliyunOSS struct {
	Client    *oss.Client
	Bucket    *oss.Bucket
	generator *utils.Generator
	config    *config.AliyunOSS
}

func New(setting *config.AliyunOSS) (*AliyunOSS,  error) {
	client, err := oss.New(setting.Endpoint, setting.AccessKeyId, setting.AccessKeySecret, oss.Timeout(setting.HTTPTimeout, setting.ReadWriteTimeout), oss.EnableCRC(setting.EnableCRC))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("inialization aliyun oss client failed: %s", err))
	}

	exist, err := client.IsBucketExist(setting.BucketName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("check bucket exists failed: %s", err))
	}
	if !exist {
		return nil, errors.New(fmt.Sprintf("aliyun oss bucket %s is not exist. ", setting.BucketName))
	}

	bucket, err := client.Bucket(setting.BucketName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("inialization aliyun oss bucket failed: %s", err))
	}

	return &AliyunOSS{
		Client:    client,
		Bucket:    bucket,
		generator: utils.NewFilenameGenerator(),
		config:    setting,
	}, nil
}

func (a *AliyunOSS) Save(header *multipart.FileHeader) (string, error) {
	name, err := a.generator.NewName()
	if err != nil {
		return "", err
	}

	p := path.Join(name.Year, name.Month, name.Day, name.Name+"."+path.Ext(header.Filename))

	open, err := header.Open()
	if err != nil {
		return "", errors.New(fmt.Sprintf("aliyun oss save failed: %s", err))
	}
	if err := a.Bucket.PutObject(p, open); err != nil {
		return "", errors.New(fmt.Sprintf("aliyun oss client PutObject failed: %s", err))
	}


	return name.Name + "." + path.Ext(header.Filename), nil
}

func (a *AliyunOSS) Delete(name ...string) (int, error) {

	var names []string
	for _, s := range name {
		timestamp, _, err := a.generator.ParseName(s)
		if err != nil {
			continue
		}
		unix := time.Unix(timestamp, 0)
		n := path.Join(strconv.FormatInt(int64(unix.Year()), 10), unix.Month().String(), strconv.FormatInt(int64(unix.Day()), 10), s)
		names = append(names, n)
	}

	objects, err := a.Bucket.DeleteObjects(names)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("aliyun oss client DeleteObjects failed: %s", err))
	}
	return len(objects.DeletedObjects), nil
}

func (a *AliyunOSS) GetAccessUrl(name string) (string, error) {
	timestamp, _, err := a.generator.ParseName(name)
	unix := time.Unix(timestamp, 0)
	n := path.Join(strconv.FormatInt(int64(unix.Year()), 10), unix.Month().String(), strconv.FormatInt(int64(unix.Day()), 10), name)
	url, err := a.Bucket.SignURL(n, oss.HTTPGet, 60)
	if err != nil {
		return "", errors.New(fmt.Sprintf("aliyun oss client get access url failed: %s", err))
	}
	return url, nil
}
