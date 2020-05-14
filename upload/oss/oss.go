/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package oss

import (
	"bytes"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"go-gulu/upload"
	"strconv"
)

type Oss struct {
	bucket *oss.Bucket
	domain string
}

func New(endpoint, accessKeyID, accessKeySecret, bucketName string, haveDomain string) (*Oss, error) {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, errors.New("oss-new 初始化失败")
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, errors.New("oss-bucket 初始化失败")
	}
//-internal
	domain := ""
	if haveDomain == "" {
		domain = "https://" + bucketName + "." + endpoint + "/"
	} else if string(haveDomain[len(haveDomain)-1:]) != "/" {
		domain = haveDomain + "/"
	}

	return &Oss{
		bucket: bucket,
		domain: domain,
	}, nil
}

func (oss *Oss) Upload(c *gin.Context, attribute *upload.FileAttribute) *upload.File {
	u := upload.File{}

	f, image, err := c.Request.FormFile(attribute.Key)
	if err != nil {
		u.Error = errors.New("未上传文件")
		return &u
	}
	defer f.Close()
	if !upload.CheckFileExt(image.Filename, &attribute.Exts) {
		u.Error = errors.New("不支持此文件类型")
		return &u
	}

	if image.Size > attribute.MaxSize {
		u.Error = errors.New("文件大小最大允许" + strconv.FormatInt(attribute.MaxSize/1024/1024, 10) + "M")
		return &u
	}

	run := 0
	finalDirFileName := ""
	imageName := ""
	for {
		imageName = upload.RandFileName(image.Filename)
		finalDirFileName = attribute.DirPath + imageName
		if isExist, _ := oss.bucket.IsObjectExist(finalDirFileName); !isExist {
			break
		}

		run++
		if run >= 5 {
			u.Error = errors.New("文件名无法构造")
			return &u
		}
	}

	buf := make([]byte, image.Size)
	_, err = f.ReadAt(buf, 0)
	if err != nil {
		u.Error = errors.New("文件IO错误")
		return &u
	}

	err = oss.bucket.PutObject(finalDirFileName, bytes.NewReader(buf))
	if err != nil {
		u.Error = errors.New("无法保存文件")
		return &u
	}

	u.Url = oss.domain + finalDirFileName
	u.Save = finalDirFileName

	return &u
}

func (oss *Oss) Remove(saveFile string) (bool, error) {
	if err := oss.bucket.DeleteObject(saveFile); err != nil {
		return false, err
	}

	return true, nil
}
