package config

// MonaPanelConfig 应用根配置
type MonaPanelConfig struct {
	Server   Server   `mapstructure:"server" json:"env" yaml:"env"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
}
