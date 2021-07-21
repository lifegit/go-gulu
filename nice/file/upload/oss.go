/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package upload

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"path"
)

type Oss struct {
	bucket *oss.Bucket
	domain string
}

func NewOss(endpoint, accessKeyID, accessKeySecret, bucketName string, haveDomain string) (*Oss, error) {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, errors.New("oss-new 初始化失败")
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, errors.New("oss-bucket 初始化失败")
	}
	//-internal
	if haveDomain == "" {
		haveDomain = fmt.Sprintf("https://%s.%s/", bucketName, endpoint)
	}
	return &Oss{
		bucket: bucket,
		domain: haveDomain,
	}, nil
}

func (oss *Oss) Upload(c *gin.Context, attribute FileAttribute) File {
	u := File{}

	f, image, err := c.Request.FormFile(attribute.Key)
	if err != nil {
		u.Error = errors.New("未上传文件")
		return u
	}
	defer f.Close()
	if !CheckFileExt(image.Filename, &attribute.Exts) {
		u.Error = errors.New("不支持此文件类型")
		return u
	}

	if image.Size > attribute.MaxByte {
		u.Error = errors.New(fmt.Sprintf("文件大小最大允许%dM", attribute.MaxByte/1024/1024))
		return u
	}

	run := 0
	finalDirFileName := ""
	for {
		imageName := RandFileName(image.Filename)
		finalDirFileName = path.Join(attribute.DirPath, imageName)
		if isExist, _ := oss.bucket.IsObjectExist(finalDirFileName); !isExist {
			break
		}

		run++
		if run >= 5 {
			u.Error = errors.New("文件名无法构造")
			return u
		}
	}

	buf := make([]byte, image.Size)
	_, err = f.ReadAt(buf, 0)
	if err != nil {
		u.Error = errors.New("文件IO错误")
		return u
	}

	err = oss.bucket.PutObject(finalDirFileName, bytes.NewReader(buf))
	if err != nil {
		u.Error = errors.New("无法保存文件")
		return u
	}

	u.Url = path.Join(oss.domain, finalDirFileName)
	u.Save = finalDirFileName

	return u
}

func (oss *Oss) Remove(filename string) error {
	return oss.bucket.DeleteObject(filename)
}
