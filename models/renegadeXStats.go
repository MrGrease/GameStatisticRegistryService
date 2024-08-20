package models

import (
	"encoding/json"
	"fmt"
)

// Inbetween types used to pass data around until we decide what DB to use

type RenegadeXPlayer struct {
	UniqueID                string `json:"unique_id"`
	Name                    string `json:"name"`
	PlaytimeMinutes         int    `json:"playtime_minutes"`
	VeterancyRank           int    `json:"veterancy_rank"`
	VeterancyPoints         int    `json:"veterancy_points"`
	Deaths                  int    `json:"deaths"`
	TotalKills              int    `json:"total_kills"`
	OffensiveKills          int    `json:"offensive_kills"`
	NeutralKills            int    `json:"neutral_kills"`
	DefensiveKills          int    `json:"defensive_kills"`
	BeaconKills             int    `json:"beacon_kills"`
	MineKills               int    `json:"mine_Kills"`
	TotalAssists            int    `json:"total_assists"`
	OffensiveAssists        int    `json:"offensive_assists"`
	DefensiveAssists        int    `json:"defensive_assists"`
	NeutralAssists          int    `json:"neutral_assists"`
	TotalVehicleKills       int    `json:"total_vehicle_kills"`
	OffensiveVehicleKills   int    `json:"offensive_vehicle_kills"`
	DefensiveVehicleKills   int    `json:"defensive_vehicle_kills"`
	NeutralVehicleKills     int    `json:"neutral_vehicle_kills"`
	TotalVehicleAssists     int    `json:"total_vehicle_assists"`
	OffensiveVehicleAssists int    `json:"offensive_vehicle_assists"`
	DefensiveVehicleAssists int    `json:"defensive_vehicle_assists"`
	NeutralVehicleAssists   int    `json:"neutral_vehicle_assists"`
	VehicleDamage           int    `json:"vehicle_damage"`
	VehicleRepairs          int    `json:"vehicle_repairs"`
	BuildingRepairs         int    `json:"building_repairs"`
	InfantryRepairs         int    `json:"infantry_repairs"`
	BeaconDamage            int    `json:"beacon_damage"`
	BuildingDamage          int    `json:"building_damage"`
	BuildingArmourDamage    int    `json:"building_armour_damage"`
	TechCaptures            int    `json:"tech_captures"`
	VehicleEmps             int    `json:"vehicle_emps"`
}

type RenegadeXTeamInput struct {
	Name           string                  `json:"name"`
	Score          int                     `json:"score"`
	BuildingsAlive int                     `json:"buildingsAlive"`
	BuildingsMax   int                     `json:"buildingsMax"`
	Players        map[int]RenegadeXPlayer `json:"players"`
}

type RenegadeXTeam struct {
	Name           string            `json:"name"`
	Score          int               `json:"score"`
	BuildingsAlive int               `json:"buildingsAlive"`
	BuildingsMax   int               `json:"buildingsMax"`
	Players        []RenegadeXPlayer `json:"players"`
}

type RenegadeXStatsInput struct {
	Winner    string             `json:"winner"`
	WinMethod string             `json:"winMethod"`
	Team0     RenegadeXTeamInput `json:"team0"`
	Team1     RenegadeXTeamInput `json:"team1"`
}

type RenegadeXStats struct {
	Winner    string        `json:"winner"`
	WinMethod string        `json:"winMethod"`
	Team0     RenegadeXTeam `json:"team0"`
	Team1     RenegadeXTeam `json:"team1"`
}

func getTeam(input RenegadeXTeamInput) RenegadeXTeam {
	team := RenegadeXTeam{
		Name:           input.Name,
		Score:          input.Score,
		BuildingsAlive: input.BuildingsAlive,
		BuildingsMax:   input.BuildingsMax,
	}
	for _, p := range input.Players {
		team.Players = append(team.Players, p)
	}
	return team
}

func (stats *RenegadeXStats) Parse(rawData []byte) error {
	fmt.Println("Attempting to parse raw data")

	statsInput := RenegadeXStatsInput{}
	err := json.Unmarshal(rawData, &statsInput)
	if err != nil {
		return err
	}
	stats.Winner = statsInput.Winner
	stats.WinMethod = statsInput.WinMethod
	stats.Team0 = getTeam(statsInput.Team0)
	stats.Team1 = getTeam(statsInput.Team1)
	// TODO do the necessary checks and changes if needed
	return nil

	// //Who won this game? integer represents teams
	// winner, err := strconv.Atoi(rawData["winner"].(string))
	// if err != nil {
	// 	return err
	// }
	// statsData.Winner = int16(winner)
	// //What was the win method? should be a string
	// winMethod := rawData["winMethod"].(string)
	// if winMethod == "" {
	// 	return errors.New("win method should exist")
	// }
	// statsData.WinMethod = winMethod

	// //Parse team data for team 0
	// team0Data := rawData["team0"].(map[string]interface{})

	// statsData.Team0.Name = team0Data["name"].(string)
	// statsData.Team0.Score = int64(team0Data["score"].(float64))

	// buildingsAlive := team0Data["buildingsAlive"].(float64)

	// buildingsMax := team0Data["buildingsMax"].(float64)

	// statsData.Team0.BuildingsAlive = int16(buildingsAlive)
	// statsData.Team0.BuildingsMax = int16(buildingsMax)
	// //parse each player
	// team0Players := team0Data["players"].(map[string]interface{})

	// statsData.Team0.Players = ParsePlayerJsonData(team0Players)
	// //Parse team data for team 1
	// team1Data := rawData["team1"].(map[string]interface{})

	// statsData.Team1.Name = team1Data["name"].(string)
	// statsData.Team1.Score = int64(team1Data["score"].(float64))

	// buildingsAlive = team1Data["buildingsAlive"].(float64)

	// buildingsMax = team1Data["buildingsMax"].(float64)

	// statsData.Team1.BuildingsAlive = int16(buildingsAlive)
	// statsData.Team1.BuildingsMax = int16(buildingsMax)
	// //parse each player
	// statsData.Team1.Players = ParsePlayerJsonData(team0Players)
	// fmt.Println(statsData)

	// return err
}

// func ParsePlayerJsonData(rawData map[string]interface{}) []RenegadeXPlayerData {
// 	var playerDataArray []RenegadeXPlayerData

// 	for key, value := range rawData {
// 		fmt.Println("parsing key ", key)
// 		mappedValue := value.(map[string]interface{})
// 		var playerData RenegadeXPlayerData
// 		playerData.UniqueId = mappedValue["unique_id"].(string)
// 		playerData.Name = mappedValue["name"].(string)
// 		playerData.PlaytimeMinutes = int16(mappedValue["playtime_minutes"].(float64))
// 		playerData.VeterancyRank = int16(mappedValue["veterancy_rank"].(float64))
// 		playerData.Deaths = int16(mappedValue["deaths"].(float64))
// 		playerData.Totalkills = int16(mappedValue["total_kills"].(float64))
// 		playerData.OffensiveKills = int16(mappedValue["offensive_kills"].(float64))
// 		playerData.NeutralKills = int16(mappedValue["neutral_kills"].(float64))
// 		playerData.DefensiveKills = int16(mappedValue["defensive_kills"].(float64))
// 		playerData.BeaconKills = int16(mappedValue["beacon_kills"].(float64))
// 		playerData.MineKills = int16(mappedValue["mine_Kills"].(float64))
// 		playerData.TotalAssists = int16(mappedValue["total_assists"].(float64))
// 		playerData.OffensiveAssists = int16(mappedValue["offensive_assists"].(float64))
// 		playerData.DefensiveAssists = int16(mappedValue["defensive_assists"].(float64))
// 		playerData.NeutralAssists = int16(mappedValue["neutral_assists"].(float64))
// 		playerData.TotalVehicleKills = int16(mappedValue["total_vehicle_kills"].(float64))
// 		playerData.OffensiveVehicleKills = int16(mappedValue["offensive_vehicle_kills"].(float64))
// 		playerData.DefensiveVehicleKills = int16(mappedValue["defensive_vehicle_kills"].(float64))
// 		playerData.NeutralVehicleKills = int16(mappedValue["neutral_vehicle_kills"].(float64))
// 		playerData.TotalVehicleassists = int16(mappedValue["total_vehicle_assists"].(float64))
// 		playerData.OffensiveVehicleAssists = int16(mappedValue["offensive_vehicle_assists"].(float64))
// 		playerData.DefensiveVehicleAssists = int16(mappedValue["defensive_vehicle_assists"].(float64))
// 		playerData.NeutralVehicleAssists = int16(mappedValue["neutral_vehicle_assists"].(float64))
// 		playerData.VehicleDamage = int32(mappedValue["vehicle_damage"].(float64))
// 		playerData.VehicleRepairs = int32(mappedValue["vehicle_repairs"].(float64))
// 		playerData.BuildingRepairs = int32(mappedValue["building_repairs"].(float64))
// 		playerData.InfantryRepairs = int32(mappedValue["infantry_repairs"].(float64))
// 		playerData.BeaconDamage = int32(mappedValue["beacon_damage"].(float64))
// 		playerData.BuildingDamage = int32(mappedValue["building_damage"].(float64))
// 		playerData.BuildingArmourDamage = int32(mappedValue["building_armour_damage"].(float64))
// 		playerData.TechCaptures = int16(mappedValue["tech_captures"].(float64))
// 		playerData.VehicleEmps = int16(mappedValue["vehicle_emps"].(float64))
// 		fmt.Println(playerData)
// 		playerDataArray = append(playerDataArray, playerData)
// 	}
// 	return playerDataArray
// }

//End of inbetween types
