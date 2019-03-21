package main

import (
	"errors"
	"flag"
	"github.com/zlikun/sql_exporter/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"strings"
)

var (
	globalPath string
	queryDir   string
)

func init() {
	flag.StringVar(&globalPath, "global_path", "", "全局配置，含默认配置、数据源配置")
	flag.StringVar(&queryDir, "query_dir", "", "查询配置目录")
	flag.Parse()
}

// YAML配置接口实现（插件）
type YamlConfig struct {
	// 全局配置路径
	globalPath string
	// 查询配置目录（内部只能存放yaml配置文件）
	queryDir string
}

// 约定插件必须实现的函数，函数签名不能为变更
func New() (config.ConfigLoader, error) {
	if globalPath == "" {
		return nil, errors.New("global_path参数不能为空")
	}
	if queryDir == "" {
		return nil, errors.New("query_dir参数不能为空")
	}
	return &YamlConfig{
		globalPath: globalPath,
		queryDir:   queryDir,
	}, nil
}

func (yc *YamlConfig) Load() (*config.Config, error) {

	conf, err := loadGlobal(yc.globalPath)
	if err != nil {
		return nil, err
	}

	if queries, err := loadQueries(yc.queryDir); err != nil {
		return nil, err
	} else {
		conf.Queries = queries
	}

	return conf, nil
}

// 加载全局配置文件
func loadGlobal(global string) (*config.Config, error) {
	if data, err := ioutil.ReadFile(global); err != nil {
		return nil, err
	} else {
		var conf *config.Config
		if err := yaml.Unmarshal(data, &conf); err != nil {
			return nil, err
		}
		return conf, nil
	}
}

// 加载查询配置文件
func loadQueries(dir string) (map[string]*config.QueryConfig, error) {

	if info, err := ioutil.ReadDir(dir); err != nil {
		return nil, err
	} else {
		var queries = make(map[string]*config.QueryConfig)
		for _, i := range info {
			if !strings.HasSuffix(i.Name(), ".yml") {
				continue
			}
			f := path.Join(dir, i.Name())
			if data, err := ioutil.ReadFile(f); err != nil {
				// 任何错误都将导致整体退出
				return nil, err
			} else {
				var t = make(map[string]*config.QueryConfig)
				if err := yaml.Unmarshal(data, &t); err != nil {
					return nil, err
				} else {
					for k, v := range t {
						queries[k] = v
					}
				}
			}
		}
		return queries, nil
	}

}
