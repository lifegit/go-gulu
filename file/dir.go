/**
* @Author: TheLife
* @Date: 2020-2-25 9:38 下午
 */
package file

import (
	"os"
	"path/filepath"
)

// 判断一个文件夹是否存在,不存在则创建
func IsNotExistMkDir(src string) error {
	if !IsExist(src) {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// 创建一个文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// 判断指定的路径是否为目录
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}

	if nil != err {
		//logger.Warnf("Determines whether [%s] is a directory failed: [%v]", path, err)
		return false
	}

	return fio.IsDir()
}

// 将source目录复制到dest目录
func CopyDir(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, err := os.Open(source)
	if err != nil {
		return err
	}

	defer directory.Close()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(srcFilePath, destFilePath)
			if err != nil {
				//logger.Error(err)
			}
		} else {
			err = CopyFile(srcFilePath, destFilePath)
			if err != nil {
				//logger.Error(err)
			}
		}
	}

	return nil
}
