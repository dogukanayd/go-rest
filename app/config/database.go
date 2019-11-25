package config

import (
	"os"
)

var connections = map[string]string{
	"mysql": os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@" + os.Getenv("MYSQL_HOST"),
}

// GetDBDSN returns connection dsn
func GetDBDSN(dbName string) string {
	return connections["mysql"] + "/" + dbName
}
