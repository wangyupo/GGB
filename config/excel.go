package config

type Excel struct {
	TemplateDir string `mapstructure:"template-dir" json:"template-dir" yaml:"template-dir"`
	OutputDir   string `mapstructure:"output-dir" json:"output-dir" yaml:"output-dir"`
}
