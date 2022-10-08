package viperine_test

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/lifegit/go-gulu/v2/pkg/viperine"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
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
	s := `
# toml 转 struct : https://github.com/xuri/toml-to-go

[app]
    env = "dev"
[server]
    addr = "127.0.0.1"

[db]
    database = "tb"
`
	err := ioutil.WriteFile("./conf.toml", []byte(s), 0666)
	defer os.Remove("./conf.toml")
	assert.NoError(t, err)

	v, err := viperine.LocalConfToViper([]string{
		"./conf.toml",
	}, &Global, func(event fsnotify.Event, viper *viper.Viper) {
		if event.Op != fsnotify.Remove {
			_ = viper.Unmarshal(&Global)
		}
	})
	if err != nil {
		logrus.WithError(err).Fatal(err)
	}

	fmt.Println(Global, v)
	assert.NoError(t, err)
}
