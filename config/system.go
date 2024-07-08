package config

type System struct {
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	OssType      string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Language     string `mapstructure:"language" json:"language" yaml:"language"`
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`
}
