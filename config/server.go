package config

// Server 运行 Gin 时需要的配置信息
type Server struct {
	Listen string `mapstructure:"listen" json:"listen" yaml:"listen"` // 监听地址
	Port   string `mapstructure:"port" json:"port" yaml:"port"`       // 监听端口
}
