package routes

import (
	"fmt"
	"io"
	"net/http"

	"gamestatsticregistry/mrgrease.com/db"
	"gamestatsticregistry/mrgrease.com/models"
	"gamestatsticregistry/mrgrease.com/parsers"

	"github.com/gin-gonic/gin"
)

func ListPlayerStats(context *gin.Context) {
	//Todo fill this in
	context.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED YET"})
}

func GetGameStats(context *gin.Context) {
	//Todo fill this in
	context.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED YET"})
}

func GetAllPlayerStats(context *gin.Context) {
	//Todo fill this in
	context.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED YET"})
}

func GetPlayerStats(context *gin.Context) {
	//Todo fill this in
	context.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED YET"})
}

func EndOfSessionSave(context *gin.Context) {
	// which game are we playing
	AppName := context.Param("name")

	fmt.Println("App name is ", AppName)
	fmt.Println("Grabbing parser type...")
	var stats models.GameStats = parsers.GetAppParsedType(AppName)

	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// parse json to an inbetween type (for sanity checks and generalization)
	stats.Parse(body)

	// which db are we using
	databaseManager, err := db.GetCurrentDbPointer()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save to db
	databaseManager.SaveEndOfSessionDataToDb(AppName, stats)

	context.JSON(http.StatusOK, gin.H{"message": "End of Session Data Saved"})
}

func MidSessionSave(context *gin.Context) {
	//Todo fill this in
	context.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED YET"})
}
