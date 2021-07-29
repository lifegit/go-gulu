/**
* @Author: TheLife
* @Date: 2021/7/19 下午5:19
 */
package app

import (
	"errors"
	"github.com/lifegit/go-gulu/v2/pkg/ginplay/tpl"
	"github.com/sirupsen/logrus"
)


type AuthType int
const (
	AuthTypeMobile AuthType = iota + 1
	AuthTypeSecurity
)

type GinPlay struct {
	AppPkg     string `json:"app_pkg" form:"app_pkg"`
	AppDir     string `json:"app_dir" form:"app_dir"`
	AppAddr    string `json:"app_addr" form:"app_addr"`
	AppPort    int    `json:"app_port" form:"app_port"`
	AuthTable  string `json:"auth_table" form:"auth_table"`
	AuthColumn string `json:"auth_column" form:"auth_column"`
	AuthType AuthType `json:"account_auth_type" form:"account_auth_type"`
	DbType     string `json:"db_type" form:"db_type"`
	DbUser     string `json:"db_user" form:"db_user"`
	DbPassword string `json:"db_password" form:"db_password"`
	DbAddr     string `json:"db_addr" form:"db_addr"`
	DbPort     int    `json:"db_port"  form:"db_port"`
	DbName     string `json:"db_name" form:"db_name"`
	DbChar     string `json:"db_char" form:"db_char"`
}


func (g *GinPlay) Run() (*App, error) {
	if g.AppPkg == "" {
		return nil, errors.New("app package name can't be empty string")
	}
	app, err := g.newApp()
	if err != nil {
		return nil, err
	}
	err = app.generateCodeBase()
	if err != nil {
		return nil, err
	}
	//go fmt codebase
	//https://cloud.tencent.com/developer/article/1417112
	err = app.goFmtCodeBase()
	if err != nil {
		logrus.WithError(err).Error("go fmt code base failed")
	}

	return app, err
}


func (g *GinPlay) newApp() (*App, error) {
	if g.AuthType == AuthTypeMobile {
		tpl.ParseOneList = append(tpl.ParseOneList, tpl.ParseTypeMobileList...)
	}else if g.AuthType == AuthTypeSecurity {
		tpl.ParseOneList = append(tpl.ParseOneList, tpl.ParseTypeSecurityList...)
	}
	cols, err := FetchDbColumn(*g)
	if err != nil {
		return nil, err
	}
	resources, err := transformToResources(cols, g.AuthTable, g.AuthColumn)
	if err != nil {
		return nil, err
	}
	return &App{
		GinPlay:    *g,
		Resources: resources,
	}, nil
}