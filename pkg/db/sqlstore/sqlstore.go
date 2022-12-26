package sqlstore

import "github.com/punitarani/centarus/pkg/setting"

// SQLStore holds the database configuration and connection
type SQLStore struct {
	Cfg *setting.DbCfg // Database configuration
	Db  *DbSession     // Database connection
}
