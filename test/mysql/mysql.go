package mysqltest

// With runs a block with locking the DB and truncating all tables in
// the DB.
func With(dsn string, block func()) error { _ = "STUB: not implemented"; return nil }

// TruncateTables truncates all tables in the DB specified by a DSN.
func TruncateTables(dsn string) error { _ = "STUB: not implemented"; return nil }
