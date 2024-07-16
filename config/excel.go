package config

type Excel struct {
	TemplateDir string `mapstructure:"template-dir" json:"template-dir" yaml:"template-dir"` // Excel模板存放路径
	OutputDir   string `mapstructure:"output-dir" json:"output-dir" yaml:"output-dir"`       // Excel输出路径
}
