package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func conf() *viper.Viper {
	// 触发加载本目录下其他文件中的 init 方法
	config := viper.New()
	//配置文件名（不带扩展名）
	config.SetConfigName("config")
	//在项目中查找配置文件的路径，可以使用相对路径，也可以使用绝对路径
	config.AddConfigPath("config")
	//多次调用以添加多个搜索路径
	//设置文件类型，这里是yaml文件
	config.SetConfigType("yaml")
	//定义用于接收配置文件的变量
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	//打印文件读取出来的内容:
	return config
}
func main() {
	fmt.Println(conf().Get("email.host"))
}
