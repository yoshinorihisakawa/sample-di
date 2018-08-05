package database

import (
	"database/sql"
)

type Database interface {
	Con() *sql.DB
}

type database struct {
	driverName     string
	dataSourceName string
}

func NewDatabase(driverName string, dataSourceName string) Database {
	return &database{driverName, dataSourceName}
}

func (d *database) Con() *sql.DB{
	db, err := sql.Open(d.driverName, d.dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return db
}
