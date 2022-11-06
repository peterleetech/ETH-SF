package config

type Config struct {
	Debug *DebugConfig
	DB    *DBConfig
	HTTP  *HTTPConfig
	Rpc   *RpcConfig
}

type DebugConfig struct {
	Enable  bool
	Verbose bool
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     uint16
	TimeZone string
}

type HTTPConfig struct {
	Listen string
	Port   uint16
}

type RpcConfig struct {
	RPCHost               string
	LedgerContractAddress string
	VaultContractAddress  string
	ChainId               int64
	AccountPriKey         string
}
