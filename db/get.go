package db

const (
	getCacheData = `SELECT data FROM http_cache WHERE rhost = ? AND rport = ? AND uri = ?`
	getInstalled = `SELECT installed FROM verified WHERE software_name = ? AND rhost = ? AND rport = ?`
)

// Look for an HTTP response in the db cache.
func GetHTTPResponse(rhost string, rport int, path string) (string, bool) {
	if GlobalSQLHandle == nil {
		return "", false
	}

	var retVal []byte
	err := GlobalSQLHandle.QueryRow(getCacheData, rhost, rport, path).Scan(&retVal)

	return string(retVal), err == nil
}

// Check the database to see if the target has been scanned for specific software. If so, return the result (so we don't do it again)
// Return is <db-value>,<ok>.
func GetVerified(product string, rhost string, rport int) (bool, bool) {
	if GlobalSQLHandle == nil {
		return false, false
	}

	retVal := false
	err := GlobalSQLHandle.QueryRow(getInstalled, product, rhost, rport).Scan(&retVal)

	return retVal, err == nil
}
