package config

type Config struct {
	Application Application `mapstructure:"app" json:"app" yaml:"app"`
	Zap         Zap         `mapstructure:"zap" json:"zap" yaml:"zap"`
}
type Application struct {
	Port        int    `mapstructure:"port" json:"port" yaml:"port"`
	RunMode     string `mapstructure:"run_mode" json:"run_mode" yaml:"run_mode"`
	RootPath    string `mapstructure:"rootPath" json:"rootPath" yaml:"rootPath"`
	BackUpPath  string `mapstructure:"backUpPath" json:"backUpPath" yaml:"backUpPath"`
	LogicDelete bool   `mapstructure:"logicDelete" json:"logicDelete" yaml:"logicDelete"`
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                          // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                       // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                       // 日志前缀
	Path          string `mapstructure:"path" json:"path"  yaml:"path"`                            // 日志文件目录
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`                // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encodeLevel"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktraceKey"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"logInConsole"`   // 输出控制台
}
