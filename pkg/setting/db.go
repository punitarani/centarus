package setting

// DbCfg database configuration settings
type DbCfg struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     int
	Name     string

	// [Db.<db>.Params] section
	Params DbCfgParams

	DSN string
}

type DbCfgParams struct {
	SslMode            string
	SslCert            string
	SslKey             string
	SslRootCert        string
	SslCrl             string
	AppName            string
	FallbackAppName    string
	ConnectTimeout     int
	Keepalives         bool
	KeepalivesIdle     int
	KeepalivesInterval int
	KeepalivesCount    int
}
