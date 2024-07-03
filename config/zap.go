package config

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
	MaxSize       int    `mapstructure:"max-size" json:"max-size" yaml:"max-size"`                   // 大小限制（单位：M）
	MaxBackups    int    `mapstructure:"max-backups" json:"max-backups" yaml:"max-backups"`          // 日志备份数
	MaxAge        int    `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // 存放时间（单位：天）
	Compress      bool   `mapstructure:"compress" json:"compress" yaml:"compress"`                   // 是否压缩
}
