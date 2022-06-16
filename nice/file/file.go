package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// 获取一个文件的拓展名
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// 判断扩展名是否为图像
func ExtIsImg(extension string) bool {
	ext := strings.ToLower(extension)

	switch ext {
	case ".jpg", ".jpeg", ".bmp", ".gif", ".png", ".svg", ".ico":
		return true
	}
	return false
}

// 写出文件
func WriteFile(b []byte, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	return err
}

// 读取文件
func ReadFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	res, err := ioutil.ReadAll(f)

	return string(res), err
}

// 获取一个文件的字节长度
func GetFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if nil != err {
		//logger.Error(err)
		return -1
	}

	return fi.Size()
}

// 判断一个文件或文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

// 获取一个文件是否有权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// 根据特定模式打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// 将source文件复制到dest文件
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}
	}

	return nil
}

// 最大可能打开一个文件
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
