/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package upload

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/nice/file"
	"io"
	"os"
	"path"
)

type Local struct {
	baseDir string
	domain  string
}

func NewLocal(baseDir, domain string) *Local {
	return &Local{
		baseDir: baseDir,
		domain:  domain,
	}
}

//	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
func (l Local) Upload(c *gin.Context, attribute FileAttribute) File {
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

	if err = ReadyDir(path.Join(l.baseDir, attribute.DirPath)); err != nil {
		u.Error = errors.New("系统无权限保存文件")
		return u
	}

	run := 0
	finalDirFileName, simpleDirFileName := "", ""
	for {
		imageName := RandFileName(image.Filename)
		simpleDirFileName = path.Join(attribute.DirPath, imageName)
		finalDirFileName = path.Join(l.baseDir, simpleDirFileName)
		if !file.IsExist(finalDirFileName) {
			break
		}

		run++
		if run >= 5 {
			u.Error = errors.New("文件名无法构造")
			return u
		}
	}

	out, err := os.Create(finalDirFileName)
	if err != nil {
		u.Error = errors.New("文件IO错误")
		return u
	}
	defer out.Close()

	if _, err = io.Copy(out, f); err != nil {
		u.Error = errors.New("无法保存文件")
		return u
	}

	u.Url = path.Join(l.domain, simpleDirFileName)
	u.Save = finalDirFileName

	return u
}

func (l Local) Remove(filename string) error {
	return os.Remove(filename)
}
