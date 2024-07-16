package config

type System struct {
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                   // 数据库类型：mysql（默认）
	OssType      string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`                // Oss类型：本地（默认）
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`                            // 服务监听端口
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"` // 路由前缀
	Language     string `mapstructure:"language" json:"language" yaml:"language"`                // 语言环境（zh|en|uk|fr）
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`             // 是否启用redis（默认否）
}
