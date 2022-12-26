package sqlstore

import "github.com/punitarani/centarus/pkg/setting"

// SQLStore holds the database configuration and connection
type SQLStore struct {
	Cfg *setting.DbCfg // Database configuration
	Db  *DbSession     // Database connection
}

// TransactionFunc is the function signature for database transaction callbacks.
//
// param sess: Database session to use for the transaction.
//
// return error: Any error that occurred during the transaction.
type TransactionFunc func(sess *DbSession) error
