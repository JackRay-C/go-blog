package aliyun_oss

import (
	"blog/core/setting"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

type AliyunOss struct {
	*setting.App
	*setting.AliyunOSS
	client *oss.Client
	Bucket *oss.Bucket
}

func NewAliyunOss(appSetting *setting.App, ossSetting *setting.AliyunOSS) *AliyunOss {
	client, err := oss.New(ossSetting.Endpoint, ossSetting.AccessKeyId, ossSetting.AccessKeySecret, oss.Timeout(ossSetting.HTTPTimeout, ossSetting.ReadWriteTimeout), oss.EnableCRC(ossSetting.EnableCRC))
	if err != nil {
		panic(fmt.Errorf("inialization aliyun oss client failed: %s\n", err))
	}

	isExist, _ := client.IsBucketExist(ossSetting.BucketName)
	if !isExist {
		panic(fmt.Errorf("aliyun oss bucket: %s is not exist. ", ossSetting.BucketName))
	}

	bucket, _ := client.Bucket(ossSetting.BucketName)

	return &AliyunOss{
		App:       appSetting,
		AliyunOSS: ossSetting,
		client:    client,
		Bucket:    bucket,
	}
}

func (a *AliyunOss) Save(p string, header *multipart.FileHeader) error {
	open, err := header.Open()
	if err != nil {
		return  err
	}
	defer open.Close()

	if err := a.Bucket.PutObject(p, open); err != nil {
		return  err
	}
	return nil
}


func (a *AliyunOss) Delete(name ...string) (int, error) {
	objects, err := a.Bucket.DeleteObjects(name)
	if err != nil {
		return 0, err
	}
	return len(objects.DeletedObjects), nil
}

func (a *AliyunOss) GetAccessUrl(prefix, name string) string {
	return a.BucketUrl + "/" + prefix + "/" + name;
}