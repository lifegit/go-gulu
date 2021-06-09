/**
* @Author: TheLife
* @Date: 2021/5/28 下午4:10
 */
package viperine_test

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-gulu/viperine"
	"testing"
	"time"
)

type GlobalConf struct {
	App struct {
		Env string `toml:"env"`
	} `toml:"app"`
	Server struct {
		Addr string `toml:"addr"`
	} `toml:"server"`
	Db struct {
		Database string `toml:"database"`
	} `toml:"db"`
}

var Global GlobalConf

func TestName(t *testing.T) {
	confBase := "./conf/base.toml"
	_, _ = viperine.LocalConfToViper([]string{confBase}, &Global, nil)
	v, err := viperine.LocalConfToViper([]string{confBase, fmt.Sprintf("./conf/%s/conf.toml", Global.App.Env)}, &Global, func(event fsnotify.Event, viper *viper.Viper) {
		if event.Op != fsnotify.Remove {
			_ = viper.Unmarshal(&Global)
		}
	})
	if err != nil {
		logrus.WithError(err).Fatal(err)
	}

	fmt.Println(Global, v)

	<-time.After(time.Hour)
}
