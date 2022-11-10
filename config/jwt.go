package config

type Jwt struct {
	Secret                  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl                  int    `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"`
	JwtBlacklistGracePeriod int    `mapstructure:"jwt_blacklist_grace_period" json:"jwt_blacklist_grace_period" yaml:"jwt_blacklist_grace_period"` // jwt过期宽限时间
}
