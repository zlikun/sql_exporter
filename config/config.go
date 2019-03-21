package config

import "time"

type Config struct {
	Defaults    *DefaultsConfig              `yaml:"defaults" json:"defaults"`
	DataSources map[string]*DataSourceConfig `yaml:"data-sources" json:"data-sources"`
	Queries     map[string]*QueryConfig      `yaml:"queries" json:"queries"`
}

type DefaultsConfig struct {
	Timeout      time.Duration `yaml:"timeout" json:"timeout"`
	Interval     time.Duration `yaml:"interval" json:"interval"`
	ValueOnError string        `yaml:"value-on-error" json:"value-on-error"`
	Driver       string        `yaml:"driver" json:"driver"`
}

type DataSourceConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
}

type QueryConfig struct {
	DataSource   string                 `json:"data-source" yaml:"data-source"`
	SQL          string                 `json:"sql" yaml:"sql"`
	Params       map[string]interface{} `json:"params" yaml:"params"`
	Interval     time.Duration          `json:"interval" yaml:"interval"`
	Timeout      time.Duration          `json:"timeout" yaml:"timeout"`
	Metrics      map[string]string      `json:"metrics" yaml:"metrics"`
	ValueOnError string                 `json:"value-on-error" yaml:"value-on-error"`
}

type ConfigLoader interface {
	Load() (*Config, error)
}
