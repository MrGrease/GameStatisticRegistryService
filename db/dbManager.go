package db

import (
	"errors"
	"fmt"
	"gamestatsticregistry/mrgrease.com/models"
	"os"
)

type DbManager interface {
	Init()
	SaveEndOfSessionDataToDb(AppName string, stats models.GameStats)
}

var DB DbManager

func GetCurrentDbPointer() (DbManager, error) {
	_, dbPresent := os.LookupEnv("DBTYPE")

	if !dbPresent {
		return nil, errors.New("Incomplete Configuration Error: No Db provided!")
	}

	return DB, nil
}

func InitCurrentDb() error {
	_, dbPresent := os.LookupEnv("DBTYPE")

	if !dbPresent {
		return errors.New("Incomplete Configuration Error: No Db provided!")
	}

	fmt.Println("Creating current DB Manager")

	switch os.Getenv("DBTYPE") {
	case "MONGO":
		DB = new(MongoManager)
	}

	return nil
}
