package upload

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"net/url"
	"path"
)

type Oss struct {
	bucket *oss.Bucket
	domain string
}

func NewOss(endpoint, accessKeyID, accessKeySecret, bucketName string) (*Oss, error) {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, errors.New("oss-new 初始化失败")
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, errors.New("oss-bucket 初始化失败")
	}

	return &Oss{
		bucket: bucket,
	}, nil
}

func (oss *Oss) URL(domain, filePath string) (*url.URL, error) {
	if domain == "" {
		domain = fmt.Sprintf("https://%s.%s/", oss.bucket.BucketName, oss.bucket.Client.Config.Endpoint)
	}

	return url.Parse(fmt.Sprintf("%s/%s", domain, filePath))
}

func (oss *Oss) Upload(c *gin.Context, attr FileAttribute) (f File, e error) {
	// init http file
	rFile, rFileHeader, err := c.Request.FormFile(attr.Key)
	if err != nil {
		return f, errors.New("未上传文件")
	}
	defer rFile.Close()

	// check file
	if e = attr.CheckFile(rFileHeader); e != nil {
		return f, e
	}

	// make file information
	for {
		f.Name = RandFileName(rFileHeader.Filename)
		f.Save = path.Join(attr.DirPath, f.Name)
		if isExist, _ := oss.bucket.IsObjectExist(f.Save); !isExist {
			break
		}
	}
	f.Size = rFileHeader.Size

	// create file
	err = oss.bucket.PutObject(f.Save, rFile)
	if err != nil {
		return f, errors.New("无法保存文件")
	}

	return
}

func (oss *Oss) Remove(filename string) error {
	return oss.bucket.DeleteObject(filename)
}
