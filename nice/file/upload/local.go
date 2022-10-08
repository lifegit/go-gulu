package upload

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/nice/file"
	"io"
	"net/url"
	"os"
	"path"
)

type Local struct {
	baseDir string
}

func NewLocal(baseDir string) *Local {
	return &Local{
		baseDir: baseDir,
	}
}

func (l *Local) URL(domain, filePath string) (*url.URL, error) {
	return url.Parse(fmt.Sprintf("%s/%s", domain, filePath))
}

//	r.StaticFS("/upload/resources", http.Dir(upload.GetresourceFullPath()))
func (l *Local) Upload(c *gin.Context, attr FileAttribute) (f File, e error) {
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
		f.Save = path.Join(l.baseDir, attr.DirPath, f.Name)
		if !file.IsExist(f.Save) {
			break
		}
	}
	f.Size = rFileHeader.Size

	// create file
	if err := ReadyDir(path.Join(l.baseDir, attr.DirPath)); err != nil {
		return f, errors.New("系统无权限保存文件")
	}
	out, err := os.Create(f.Save)
	if err != nil {
		return f, errors.New("文件IO错误")
	}
	defer out.Close()
	if _, err = io.Copy(out, rFile); err != nil {
		return f, errors.New("无法保存文件")
	}

	return
}

func (l *Local) Remove(filename string) error {
	return os.Remove(filename)
}
