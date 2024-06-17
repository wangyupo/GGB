package config

type System struct {
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
}
