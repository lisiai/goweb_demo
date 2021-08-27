package settings

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func InitModules() (err error) {
	err = viper.UnmarshalKey("mysql", &MysqlC)
	err = viper.UnmarshalKey("redis", &RedisC)
	err = viper.UnmarshalKey("log", &LogC)
	err = viper.UnmarshalKey("web_app", &AppC)
	if err != nil {
		log.Println("反序列化失败", err)
	}
	return
}
func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./settings/")
	err = viper.ReadInConfig()
	if err != nil {
		log.Println("加载配置信息错误：", err)
	}
	err = InitModules()

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = InitModules()
	})
	return err
}
