package config

type Config struct {
	Application Application `mapstructure:"app" json:"app" yaml:"app"`
}
type Application struct {
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	RunMode  string `mapstructure:"run_mode" json:"run_mode" yaml:"run_mode"`
	RootPath string `mapstructure:"rootPath" json:"rootPath" yaml:"rootPath"`
}
