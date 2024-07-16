package db

import (
	"fmt"
	"gamestatsticregistry/mrgrease.com/models"
)

type MongoManager struct {
}

func (manager *MongoManager) Init() {
	fmt.Println("Initializing mongo")
}

func (manager *MongoManager) SaveEndOfSessionDataToDb(AppName string, stats models.GameStats) {
	fmt.Println("Saving End of session data to db")
}
