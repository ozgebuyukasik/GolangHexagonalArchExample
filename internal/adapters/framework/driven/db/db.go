package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("Db connection failure: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Db ping failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (dbAdapter Adapter) CloseDbConnection() {
	err := dbAdapter.db.Close()
	if err != nil {
		log.Fatalf("Db connection close failure: %v", err)
	}
	log.Println("Db connection is closed")
}

func (dbAdapter Adapter) AddToHistory(result int32, operation string) error {
	queryString, args, err := sq.Insert("artih_history").Columns("date", "result", "operation").
		Values(time.Now(), result, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = dbAdapter.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
