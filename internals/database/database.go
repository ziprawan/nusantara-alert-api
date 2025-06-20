package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbInstance *sql.DB
	once       sync.Once
)

func ConnectDB() *sql.DB {
	once.Do(func() {
		var err error
		dbInstance, err = sql.Open("sqlite3", "database/data.db")
		if err != nil {
			panic(fmt.Errorf("failed to connect to database: %v", err))
		}

		dbInstance.SetMaxOpenConns(15)
		dbInstance.SetMaxIdleConns(5)
	})

	return dbInstance
}

func GetDB() *sql.DB {
	if dbInstance == nil {
		// panic(errors.New("database is not initialized! Call ConnectDB() first"))
		ConnectDB()
	}

	return dbInstance
}
