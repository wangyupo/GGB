package config

import "fmt"

type Mysql struct {
	Database `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Dbname,
		m.Charset,
		m.Collation,
	)
}
