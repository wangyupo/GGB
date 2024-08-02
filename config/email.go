package config

type Email struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // SMTP 发送邮件服务器地址
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             // SMTP 发送邮件服务器端口号
	Username string `mapstructure:"username" json:"username" yaml:"username"` // SMTP 帐户（你的QQ邮箱完整的地址）
	Password string `mapstructure:"password" json:"password" yaml:"password"` // SMTP 密码（生成的授权码）
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // 发件人地址
}
