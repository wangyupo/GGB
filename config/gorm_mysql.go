package config

type Mysql struct {
	Database `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?charset=" + m.Charset + "&parseTime=True&loc=Local"
}
