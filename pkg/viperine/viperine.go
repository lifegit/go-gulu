package viperine

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func LocalConfToViper(pathNames []string, callData interface{}, watchChange func(event fsnotify.Event, viper *viper.Viper)) (v *viper.Viper, err error) {
	v = viper.New()
	for _, value := range pathNames {
		v.SetConfigFile(value)
		if err = v.MergeInConfig(); err != nil {
			return
		}
		if watchChange != nil {
			v.WatchConfig()
			v.OnConfigChange(func(in fsnotify.Event) {
				v.SetConfigFile(in.Name)
				if err = v.MergeInConfig(); err == nil {
					watchChange(in, v)
				}
			})
		}
	}
	err = v.Unmarshal(&callData)

	return
}
