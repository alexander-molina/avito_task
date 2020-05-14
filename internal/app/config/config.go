package config

import "time"

// AppConfig represents main configs for application
var appConfig *Config = &Config{"24", 100, time.Minute * 2}

// Config represent app config data
type Config struct {
	SubnetPrefixSize string
	ReqestLimit      int
	BlockTime        time.Duration
}

// GetConfig create new config
func GetConfig() Config {
	return *appConfig
}

// SetConfig set new config
func SetConfig(c Config) {
	appConfig = &c
}
