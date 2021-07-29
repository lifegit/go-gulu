package tpl

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const backQuote = "_[BACKQUOTE]_"

type tplNode struct {
	NameFormat string
	TplContent string
}

func (n *tplNode) ParseExecute(appDir, pathArg string, data interface{}) error {
	var p string
	if pathArg != "" {
		if strings.Index(n.NameFormat, "handlers/") == 0 && pathArg[0:3] == "tb_" {
			pathArg = pathArg[3:]
		}
		p = strings.Replace(n.NameFormat, "%s", pathArg, -1)
	} else {
		p = n.NameFormat
	}
	p = filepath.Clean(filepath.Join(appDir, p))
	err := os.MkdirAll(filepath.Dir(p), 0777)
	if err != nil {
		return err
	}
	tplFormat := strings.Replace(n.TplContent, backQuote, "`", -1)
	tmpl, err := template.New(p).Parse(tplFormat)
	file, err := os.Create(p)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
