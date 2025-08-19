package main

import (
	"database/sql"

	_ "github.com/microsoft/go-mssqldb"
)

type PartitionRepository struct {
	db *sql.DB
}

func NewPartitionRepository() (*PartitionRepository, error) {

	db, err := sql.Open(driverName, connectionString)

	if err != nil {
		return nil, err
	}

	return &PartitionRepository{db: db}, nil
}

func getConnection() *sql.DB {

	connString := "server=localhost;user id=oseries9;password=p@ssw0rd;database=oseries917"

	logger.Debug(connString)

	// Open connection
	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	return conn

}
