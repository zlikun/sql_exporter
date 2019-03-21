package main

import (
	"flag"
	"fmt"
	"log"
	"testing"
)

func Test_YamlConfig_Load(t *testing.T) {

	// 需要指定命令行参数，单元测试中直接使用代码指定
	flag.Set("global_path", "../../example/conf/global.yml")
	flag.Set("query_dir", "../../example/conf/queries")

	yamlConfig, err := New()
	if err != nil {
		log.Fatal(err)
	}

	if conf, err := yamlConfig.Load(); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("默认配置", conf.Defaults)
		fmt.Println("数据源配置", conf.DataSources)
		if conf.DataSources != nil {
			for k, v := range conf.DataSources {
				fmt.Println("key = ", k, "value=", v)
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
