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
	Params map[DbCfgParams]string

	DSN string
}

// DbCfgParams database configuration parameters value
//
//	 Supported parameters:
//		SslMode            string
//		SslCert            string
//		SslKey             string
//		SslRootCert        string
//		SslCrl             string
//		AppName            string
//		FallbackAppName    string
//		ConnectTimeout     int
//		Keepalives         bool
//		KeepalivesIdle     int
//		KeepalivesInterval int
//		KeepalivesCount    int
//
// All parameters must be cast to string.
type DbCfgParams string
