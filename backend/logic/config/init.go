package config

import (
	"github.com/spf13/viper"
	"strings"
)

var config Config = Config{
	Debug: &DebugConfig{
		Enable:  true,
		Verbose: true,
	},
	DB: &DBConfig{
		Host:     "127.0.0.1",
		Port:     33061,
		User:     "sms_backend",
		Password: "password",
		DbName:   "sms_backend",
		TimeZone: "Local",
	},
	HTTP: &HTTPConfig{
		Listen: "0.0.0.0",
		Port:   8888,
	},
	Rpc: &RpcConfig{
		RPCHost:               "https://api.s0.b.hmny.io",
		LedgerContractAddress: "0x60e77693a8054eb786ef7631e47aab8a28b3fbd2",
		VaultContractAddress:  "0xd3446851deb19bcf700dadef258ba90834c8472a",
		ChainId:               1666700000,
	},
}

func init() {
	instance := viper.New()

	// only for dev
	instance.AddConfigPath("/etc/veric-backend/")
	instance.AddConfigPath(".")

	instance.SetConfigType("yaml")
	instance.SetConfigName("config.veric-backend.yaml")

	instance.SetEnvPrefix("vb")
	instance.AutomaticEnv()
	instance.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := instance.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}

	err = instance.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	return &config
}
