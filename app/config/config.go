package config

import (
	// import built-in packages
	"fmt"
	// import third-party packages
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type server struct {
	Port int
}

type database struct {
	Type string
}

type config struct {
	Server   server
	Database database
}

func ReadConfig() config {
	viper := viper.New()
	viper.SetConfigFile("./conf/config.toml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("The config file is not existd.")
	}
	viper.WatchConfig()
	var conf config
	viper.Unmarshal(&conf)
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&conf); err != nil {
			fmt.Println(err.Error())
		} else {
			return
		}
	})
	// return
	return conf
}

var (
	ServerPort   = ReadConfig().Server.Port
	DatabaseType = ReadConfig().Database.Type
)
