package utils

import (
	"github.com/dogukanayd/go-rest/app/config"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

type MySQLInterface interface {
	Connect(name string) MySQLInterface
	Ping() error
	Close() error
}

type MySQL struct {
	connection sqlbuilder.Database
}

// Connect to database
func (m *MySQL) Connect(name string) MySQLInterface {
	m.connection = m.connect(name)

	return m
}

// NewConnection returns new mysql database connection instance
func NewMySQL() MySQLInterface {
	return &MySQL{}
}

// Ping mysql connection
func (m *MySQL) Ping() error {
	return m.connection.Ping()
}

func (m *MySQL) Close() error {
	return m.connection.Close()
}

func (m *MySQL) connect(dbName string) sqlbuilder.Database {
	url, err := mysql.ParseURL(config.GetDBDSN(dbName))

	if err != nil {
		panic(err)
	}

	session, err := mysql.Open(url)

	if err != nil {
		panic(err)
	}

	return session
}
