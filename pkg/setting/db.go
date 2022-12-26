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
	Params map[DbCfgParam]string

	DSN string
}

// DbCfgParam database configuration parameter
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
type DbCfgParam string

// ValidateDbCfg checks if the DbCfg struct is valid.
func ValidateDbCfg(cfg *DbCfg) error {
	// Check Driver
	switch cfg.Driver {
	case "postgresql":
		// Valid Driver
	default:
		return fmt.Errorf("invalid Driver: %s", cfg.Driver)
	}

	// Ensure that Username, Password, Host, and Name are not empty.
	for key, val := range []string{cfg.Username, cfg.Password, cfg.Host, cfg.Name} {
		if val == "" {
			return fmt.Errorf("%v is empty", key)
		}
	}

	// Check Port range
	if !(1 <= cfg.Port && cfg.Port <= 65535) {
		return fmt.Errorf("invalid port: %d", cfg.Port)
	}

	// Validate Params
	var invalidParams []string
	for k := range cfg.Params {
		if !isValidDbCfgParam(string(k)) {
			invalidParams = append(invalidParams, string(k))
		}
	}
	if len(invalidParams) > 0 {
		return fmt.Errorf("invalid DbCfg params: %v", invalidParams)
	}

	return nil
}

// isValidDbCfgParam checks if the DbCfgParam is valid.
func isValidDbCfgParam(param string) bool {
	switch DbCfgParam(param) {
	case
		"SslMode",
		"SslCert",
		"SslKey",
		"SslRootCert",
		"SslCrl",
		"AppName",
		"FallbackAppName",
		"ConnectTimeout",
		"Keepalives",
		"KeepalivesIdle",
		"KeepalivesInterval",
		"KeepalivesCount":
		return true
	default:
		return false
	}
}

// BuildDSN builds the Data Source Name (DSN) connection url from the DbCfg struct.
func BuildDSN(cfg *DbCfg) string {
	var dsn string

	// Build the base DSN
	dsn = fmt.Sprintf("%s://%s:%s@%s:%d/%s", cfg.Driver, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	// Add the Params
	if len(cfg.Params) > 0 {
		values := url.Values{}
		for k, v := range cfg.Params {
			values.Add(string(k), v)
		}
		dsn = fmt.Sprintf("%s?%s", dsn, values.Encode())
	}

	// Update the DbCfg struct
	cfg.DSN = dsn

	return dsn
}
