package setting

// DbCfg database configuration settings
type DbCfg struct {
	Driver   string `toml:"driver"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`

	// [dbs.db.params] section
	Params DbCfgParams `toml:"params"`

	DSN string `toml:"dsn"`
}

type DbCfgParams struct {
	SslMode            string `toml:"ssl_mode"`
	SslCert            string `toml:"ssl_cert"`
	SslKey             string `toml:"ssl_key"`
	SslRootCert        string `toml:"ssl_root_cert"`
	SslCrl             string `toml:"ssl_crl"`
	AppName            string `toml:"app_name"`
	FallbackAppName    string `toml:"fallback_app_name"`
	ConnectTimeout     int    `toml:"connect_timeout"`
	Keepalives         bool   `toml:"keepalives"`
	KeepalivesIdle     int    `toml:"keepalives_idle"`
	KeepalivesInterval int    `toml:"keepalives_interval"`
	KeepalivesCount    int    `toml:"keepalives_count"`
}
