package config

type Database struct {
	Host      string `mapstructure:"host" json:"host" yaml:"host"`                // 数据库地址
	Port      string `mapstructure:"port" json:"port" yaml:"port"`                // 数据库端口号
	Username  string `mapstructure:"username" json:"username" yaml:"username"`    // 数据库账号
	Password  string `mapstructure:"password" json:"password" yaml:"password"`    // 数据库密码
	Dbname    string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`       // 数据库名
	Charset   string `mapstructure:"charset" json:"charset" yaml:"charset"`       // 数据库编码方式
	Collation string `mapstructure:"collation" json:"collation" yaml:"collation"` // 数据库字符集
}
