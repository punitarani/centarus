package setting

import (
	"fmt"
	"net/url"
)

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

// BuildDSN builds the Data Source Name (DSN) connection url from the DbCfg struct.
func BuildDSN(cfg *DbCfg) string {
	var dsn string

	// Build the Base DSN
	dsn = fmt.Sprintf("%s://%s:%s@%s:%d/%s", cfg.Driver, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	// Add the Params
	if len(cfg.Params) > 0 {
		values := url.Values{}
		for k, v := range cfg.Params {
			values.Add(string(k), v)
		}
		dsn = fmt.Sprintf("%s?%s", dsn, values.Encode())
	}

	return dsn
}
