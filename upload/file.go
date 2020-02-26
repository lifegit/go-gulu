/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package upload

import (
	"fmt"
	"go-gulu/crypto"
	"go-gulu/file"
	"go-gulu/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type File struct {
	Error error
	Url   string
	Save  string
}

type FileAttribute struct {
	Key     string
	Exts    []string
	DirPath string
	MaxSize int64
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
		if strings.ToUpper(string(allowExt)) == strings.ToUpper(ext) {
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
