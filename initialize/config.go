package initialize

import (
	"file-manager/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func InitConfig(configFile string) {
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			log.Fatal("配置解析失败:", err.Error())
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		log.Fatal("配置解析失败:", err.Error())
	}
	global.VP = v
}
