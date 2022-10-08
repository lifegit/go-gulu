package upload

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/nice/crypto"
	"github.com/lifegit/go-gulu/v2/nice/file"
	"github.com/lifegit/go-gulu/v2/nice/rand"
	"mime/multipart"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var (
	AllowImageExts = []string{".jpeg", ".jpg", ".png", ".bmp", ".gif"}
	AllowAnyExts   = []string{"*"}
)

type FileUpload interface {
	Upload(c *gin.Context, attribute FileAttribute) (File, error)
	Remove(path string) error
	URL(domain, filePath string) (*url.URL, error)
}

type File struct {
	Name string
	Size int64
	Save string
}

type FileAttribute struct {
	Key     string
	Exts    []string
	DirPath string
	MaxByte int64
}

func (a FileAttribute) CheckFile(fileHeader *multipart.FileHeader) (err error) {
	if strings.Join(a.Exts, "") != strings.Join(AllowAnyExts, "") {
		if !CheckFileExt(fileHeader.Filename, &a.Exts) {
			return errors.New("不支持此文件类型")
		}
	}
	if fileHeader.Size > a.MaxByte {
		return errors.New(fmt.Sprintf("文件大小最大允许%dM", a.MaxByte/1024/1024))
	}

	return
}

// 使用md5随机生成一个文件名
func RandFileName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = fileName + "_" + strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.Itoa(rand.Int(10000, 99999))
	fileName = crypto.EncodeMD5(fileName)

	return fileName + ext
}

// 判断一个文件的拓展名是否在切片内
func CheckFileExt(fileName string, exts *[]string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range *exts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// 就绪一个文件夹
func ReadyDir(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
