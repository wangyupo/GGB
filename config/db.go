package config

type Database struct {
	Host      string `mapstructure:"host" json:"host" yaml:"host"`
	Port      string `mapstructure:"port" json:"port" yaml:"port"`
	Dbname    string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username  string `mapstructure:"username" json:"username" yaml:"username"`
	Password  string `mapstructure:"password" json:"password" yaml:"password"`
	Charset   string `mapstructure:"charset" json:"charset" yaml:"charset"`
	Collation string `mapstructure:"collation" json:"collation" yaml:"collation"`
}
