package main

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

type Partition struct {
	sourceName string
	trustedId string
	description string
	uuid string
	displayName string
}

func (r PartitionRepository) GetPartitions() []Partition {
    
	rows, err := r.db.Query("select sourceName, trustedId, description, uuid, displayName from student")
    
	if err != nil {
        logger.Error(err.Error())
		panic(err)
    }
    
	defer rows.Close()
    //iterate through the row

	var partitions []Partition

	for rows.Next() {
		var sourceName string
		var trustedId string
		var description string
		var uuid string
		var displayName string

		err = rows.Scan(&sourceName, &trustedId, &description, &uuid, &displayName)
		if err != nil {
			logger.Error(err.Error())
			panic(err)
		}
		
		for rows.Next() {}
			partitions = append( partitions, Partition{ sourceName: "sourceName", trustedId: "trustedId", description: "description", uuid: "uuid", displayName: "displayName"})
	}

	return partitions
}

func getSources() {
	// Create connection string
	
	conn := getConnection()

	rows, err := conn.Query("SELECT * FROM Source")
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	defer rows.Close()

	var sourceId int
		var sourceName string
		var locale sql.NullString
		var trustId sql.NullString
		var description sql.NullString
		var uuid string
		var displayName string
		var prodInd int

	// Iterate through results
	for rows.Next() {
		
		err = rows.Scan(&sourceId, &sourceName, &locale, &trustId, &description, &uuid, &displayName, &prodInd)
		if err != nil {
			logger.Error(err.Error())
			panic(err)
		}
		fmt.Println(description.String)
		// Process results
	}
	if err = rows.Err(); err != nil {
		logger.Error( err.Error())
		panic(err)
	}

	return
}