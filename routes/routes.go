package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/stats/game/:name", GetAllGameStats)            //Get paginated list of all game results for given game name
	server.GET("/stats/game/:name/session/:id", GetGameStats)   //Get stats for a single game session
	server.GET("/stats/game/:name/players", GetAllPlayerStats)  //Get paginated list of all player stats for given game name
	server.GET("/stats/game/:name/players/:id", GetPlayerStats) //Get paginated list of all player stats for given game name and given player id

	server.POST("/stats/game/:name/endofsession", EndOfSessionSave) //Save results for end of session
	server.POST("/stats/game/:name/midsession", MidSessionSave)     // Save results during the session
}
