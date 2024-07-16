package db

import (
	"fmt"
	"gamestatsticregistry/mrgrease.com/models"
	"os"
)

type DbManager interface {
	Init()
	SaveEndOfSessionDataToDb(AppName string, stats models.GameStats)
}

var DB DbManager

func GetCurrentDbPointer() DbManager {
	_, dbPresent := os.LookupEnv("DBTYPE")

	if !dbPresent {
		panic("Incomplete Configuration Error: No Db provided!")
	}

	return DB
}

func InitCurrentDb() {
	_, dbPresent := os.LookupEnv("DBTYPE")

	if !dbPresent {
		panic("Incomplete Configuration Error: No Db provided!")
	}

	fmt.Println("Creating current DB Manager")

	switch os.Getenv("DBTYPE") {
	case "MONGO":
		DB = new(MongoManager)
	}
}
