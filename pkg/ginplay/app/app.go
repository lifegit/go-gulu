package app

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/pkg/ginplay/tpl"
	"os"
	"os/exec"
	"path/filepath"
)

type App struct {
	GinPlay
	Resources []Resource `json:"-"`
	Files     []string
}

func (app *App) generateCodeBase() error {
	for _, tplNode := range tpl.ParseOneList {
		err := tplNode.ParseExecute(app.AppDir, "", app)
		if err != nil {
			return fmt.Errorf("parse [%s] template failed with error : %s", tplNode.NameFormat, err)
		}
	}

	for _, resource := range app.Resources {
		resource.AppPkg = app.AppPkg
		tableName := resource.TableName
		//generate model from resource

		for _, tplNode := range tpl.ParseObjList {
			err := tplNode.ParseExecute(app.AppDir, tableName, resource)
			if err != nil {
				return fmt.Errorf("parse [%s] template failed with error : %s", tplNode.NameFormat, err)
			}
		}
	}

	return nil
}
func (app *App) goFmtCodeBase() error {
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = app.AppDir
	cmd.Env = append(os.Environ(), "GOPROXY=https://goproxy.io")
	bb, err := cmd.CombinedOutput()
	if err != nil {
		//print gin-goinc/autols failure
		// fix it :::  https://github.com/gin-gonic/gin/issues/1673
		return fmt.Errorf("%s   %s", string(bb), err)
	}
	return nil
}
func (app *App) ListAppFileTree() error {
	return filepath.Walk(app.AppDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		if !info.IsDir() {
			app.Files = append(app.Files, path)
		}
		return nil
	})
}
