package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/wangyupo/GGB/core/internal"
	"github.com/wangyupo/GGB/global"
	"os"
)

// Viper 配置管理器
func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 厂里存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用 gin 模式的 %s 环境名称，config 的路径为 %s\n", gin.Mode(), internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用 gin 模式的 %s 环境名称，config 的路径为 %s\n", gin.Mode(), internal.ConfigDefaultFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用 gin 模式的 %s 环境名称，config 的路径为 %s\n", gin.Mode(), internal.ConfigDefaultFile)
				}
			} else {
				config = configEnv
				fmt.Printf("您正在使用 %s 环境变量，config 的路径为 %s\n", internal.ConfigEnv, config)
			}
		} else {
			fmt.Printf("您正在使用命令行 -c 参数传递的值，config 的路径为 %s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用 func Viper() 传递的值，config 的路径为 %s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		err := v.Unmarshal(&global.GGB_CONFIG)
		if err != nil {
			fmt.Println(err)
		}
	})

	err = v.Unmarshal(&global.GGB_CONFIG)
	if err != nil {
		panic(err)
	}

	return v
}
