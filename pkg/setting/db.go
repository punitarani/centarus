package setting

// DbCfg database configuration settings
type DbCfg struct {
	Driver   string `toml:"driver"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`

	DSN string `toml:"dsn"`
}
