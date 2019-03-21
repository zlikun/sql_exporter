package main

import (
	"fmt"
	"github.com/zlikun/sql_exporter/config"
	"log"
	"plugin"
)

func init() {

}

func main() {

	// 加载配置插件
	dll, err := plugin.Open("plugins/config/yaml.so")
	if err != nil {
		log.Fatal(err)
	}

	// 查找插件中的 New 函数（插件约定）
	fn, err := dll.Lookup("New")
	if err != nil {
		log.Fatal(err)
	}

	// New函数签名在这里使用（类型断言）
	loader, err := fn.(func() (config.ConfigLoader, error))()
	if err != nil {
		log.Fatal(err)
	}

	// 加载配置
	if conf, err := loader.Load(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("默认配置", conf.Defaults)
		fmt.Println("数据源配置", conf.DataSources)
		if conf.DataSources != nil {
			for k, v := range conf.DataSources {
				fmt.Println("key = ", k, "value = ", v)
			}
		}
		fmt.Println("查询配置", conf.Queries)
		if conf.Queries != nil {
			for k, v := range conf.Queries {
				fmt.Println("key = ", k, "value=", v)
			}
		}
	}

}
