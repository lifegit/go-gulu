/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package local

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/nice/file"
	upload2 "github.com/lifegit/go-gulu/v2/nice/file/upload"
	"io"
	"os"
	"strconv"
)

type Local struct {
	baseDir string
	domain  string
}

func New(baseDir, haveDomain string) *Local {
	dir := ""
	if string(baseDir[len(baseDir)-1:]) != "/" {
		dir = baseDir + "/"
	}
	domain := ""
	if string(haveDomain[len(haveDomain)-1:]) != "/" {
		domain = haveDomain + "/"
	}
	return &Local{
		baseDir: dir,
		domain:  domain,
	}
}

// // 	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
func (l Local) Upload(c *gin.Context, attribute *upload2.FileAttribute) *upload2.File {
	u := upload2.File{}

	f, image, err := c.Request.FormFile(attribute.Key)
	if err != nil {
		u.Error = errors.New("未上传文件")
		return &u
	}
	defer f.Close()
	if !upload2.CheckFileExt(image.Filename, &attribute.Exts) {
		u.Error = errors.New("不支持此文件类型")
		return &u
	}
	if image.Size > attribute.MaxSize {
		u.Error = errors.New("文件大小最大允许" + strconv.FormatInt(attribute.MaxSize/1024/1024, 10) + "M")
		return &u
	}

	if err = upload2.ReadyDir(l.baseDir + attribute.DirPath); err != nil {
		u.Error = errors.New("系统无权限保存文件")
		return &u
	}

	run := 0
	finalDirFileName := ""
	simpleDirFileName := ""
	imageName := ""
	for {
		imageName = upload2.RandFileName(image.Filename)
		simpleDirFileName = attribute.DirPath + imageName
		finalDirFileName = l.baseDir + simpleDirFileName
		if !file.IsExist(finalDirFileName) {
			break
		}

		run++
		if run >= 5 {
			u.Error = errors.New("文件名无法构造")
			return &u
		}
	}

	out, err := os.Create(finalDirFileName)
	if err != nil {
		u.Error = errors.New("文件IO错误")
		return &u
	}
	defer out.Close()

	if _, err = io.Copy(out, f); err != nil {
		u.Error = errors.New("无法保存文件")
		return &u
	}

	u.Url = l.domain + simpleDirFileName
	u.Save = finalDirFileName

	return &u
}

func (l Local) Remove(saveFile string) (bool, error) {
	if err := os.Remove(saveFile); err != nil {
		return false, err
	}

	return true, nil
}
