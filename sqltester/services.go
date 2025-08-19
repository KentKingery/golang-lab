package main

import (
	_ "github.com/microsoft/go-mssqldb"
)

type Partition struct {
	sourceName  string
	trustedId   string
	description string
	uuid        string
	displayName string
}

func GetPartitions() []Partition {

	r, err := NewPartitionRepository()

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

		for rows.Next() {
		}
		partitions = append(partitions, Partition{sourceName: "sourceName", trustedId: "trustedId", description: "description", uuid: "uuid", displayName: "displayName"})
	}

	return partitions
}
