package models

import (
	"errors"
	"fmt"
	"strconv"
)

// Inbetween types used to pass data around until we decide what DB to use
type RenegadeXPlayerData struct {
	UniqueId                string
	Name                    string
	PlaytimeMinutes         int16
	VeterancyRank           int16
	Deaths                  int16
	Totalkills              int16
	OffensiveKills          int16
	NeutralKills            int16
	DefensiveKills          int16
	BeaconKills             int16
	MineKills               int16
	TotalAssists            int16
	OffensiveAssists        int16
	DefensiveAssists        int16
	NeutralAssists          int16
	TotalVehicleKills       int16
	OffensiveVehicleKills   int16
	DefensiveVehicleKills   int16
	NeutralVehicleKills     int16
	TotalVehicleassists     int16
	OffensiveVehicleAssists int16
	DefensiveVehicleAssists int16
	NeutralVehicleAssists   int16
	VehicleDamage           int32
	VehicleRepairs          int32
	BuildingRepairs         int32
	InfantryRepairs         int32
	BeaconDamage            int32
	BuildingDamage          int32
	BuildingArmourDamage    int32
	TechCaptures            int16
	VehicleEmps             int16
}

type RenegadeXTeamData struct {
	Name           string
	Score          int64
	BuildingsAlive int16
	BuildingsMax   int16
	Players        []RenegadeXPlayerData
}

type RenegadeXStats struct {
	Winner    int16
	WinMethod string
	Team0     RenegadeXTeamData
	Team1     RenegadeXTeamData
}

func (statsData *RenegadeXStats) ParseJsonData(rawData map[string]interface{}) error {
	fmt.Println("Attempting to parse raw data")

	//Who won this game? integer represents teams
	winner, err := strconv.Atoi(rawData["winner"].(string))
	if err != nil {
		return err
	}
	statsData.Winner = int16(winner)
	//What was the win method? should be a string
	winMethod := rawData["winMethod"].(string)

	if winMethod == "" {
		return errors.New("win method should exist")
	}
	statsData.WinMethod = winMethod

	//Parse team data for team 0
	team0Data := rawData["team0"].(map[string]interface{})

	statsData.Team0.Name = team0Data["name"].(string)
	statsData.Team0.Score = int64(team0Data["score"].(float64))

	buildingsAlive := team0Data["buildingsAlive"].(float64)

	buildingsMax := team0Data["buildingsMax"].(float64)

	statsData.Team0.BuildingsAlive = int16(buildingsAlive)
	statsData.Team0.BuildingsMax = int16(buildingsMax)
	//parse each player
	team0Players := team0Data["players"].(map[string]interface{})

	statsData.Team0.Players = ParsePlayerJsonData(team0Players)
	//Parse team data for team 1
	team1Data := rawData["team1"].(map[string]interface{})

	statsData.Team1.Name = team1Data["name"].(string)
	statsData.Team1.Score = int64(team1Data["score"].(float64))

	buildingsAlive = team1Data["buildingsAlive"].(float64)

	buildingsMax = team1Data["buildingsMax"].(float64)

	statsData.Team1.BuildingsAlive = int16(buildingsAlive)
	statsData.Team1.BuildingsMax = int16(buildingsMax)
	//parse each player
	statsData.Team1.Players = ParsePlayerJsonData(team0Players)
	fmt.Println(statsData)

	return err
}

func ParsePlayerJsonData(rawData map[string]interface{}) []RenegadeXPlayerData {
	var playerDataArray []RenegadeXPlayerData

	for key, value := range rawData {
		fmt.Println("parsing key ", key)
		mappedValue := value.(map[string]interface{})
		var playerData RenegadeXPlayerData
		playerData.UniqueId = mappedValue["unique_id"].(string)
		playerData.Name = mappedValue["name"].(string)
		playerData.PlaytimeMinutes = int16(mappedValue["playtime_minutes"].(float64))
		playerData.VeterancyRank = int16(mappedValue["veterancy_rank"].(float64))
		playerData.Deaths = int16(mappedValue["deaths"].(float64))
		playerData.Totalkills = int16(mappedValue["total_kills"].(float64))
		playerData.OffensiveKills = int16(mappedValue["offensive_kills"].(float64))
		playerData.NeutralKills = int16(mappedValue["neutral_kills"].(float64))
		playerData.DefensiveKills = int16(mappedValue["defensive_kills"].(float64))
		playerData.BeaconKills = int16(mappedValue["beacon_kills"].(float64))
		playerData.MineKills = int16(mappedValue["mine_Kills"].(float64))
		playerData.TotalAssists = int16(mappedValue["total_assists"].(float64))
		playerData.OffensiveAssists = int16(mappedValue["offensive_assists"].(float64))
		playerData.DefensiveAssists = int16(mappedValue["defensive_assists"].(float64))
		playerData.NeutralAssists = int16(mappedValue["neutral_assists"].(float64))
		playerData.TotalVehicleKills = int16(mappedValue["total_vehicle_kills"].(float64))
		playerData.OffensiveVehicleKills = int16(mappedValue["offensive_vehicle_kills"].(float64))
		playerData.DefensiveVehicleKills = int16(mappedValue["defensive_vehicle_kills"].(float64))
		playerData.NeutralVehicleKills = int16(mappedValue["neutral_vehicle_kills"].(float64))
		playerData.TotalVehicleassists = int16(mappedValue["total_vehicle_assists"].(float64))
		playerData.OffensiveVehicleAssists = int16(mappedValue["offensive_vehicle_assists"].(float64))
		playerData.DefensiveVehicleAssists = int16(mappedValue["defensive_vehicle_assists"].(float64))
		playerData.NeutralVehicleAssists = int16(mappedValue["neutral_vehicle_assists"].(float64))
		playerData.VehicleDamage = int32(mappedValue["vehicle_damage"].(float64))
		playerData.VehicleRepairs = int32(mappedValue["vehicle_repairs"].(float64))
		playerData.BuildingRepairs = int32(mappedValue["building_repairs"].(float64))
		playerData.InfantryRepairs = int32(mappedValue["infantry_repairs"].(float64))
		playerData.BeaconDamage = int32(mappedValue["beacon_damage"].(float64))
		playerData.BuildingDamage = int32(mappedValue["building_damage"].(float64))
		playerData.BuildingArmourDamage = int32(mappedValue["building_armour_damage"].(float64))
		playerData.TechCaptures = int16(mappedValue["tech_captures"].(float64))
		playerData.VehicleEmps = int16(mappedValue["vehicle_emps"].(float64))
		fmt.Println(playerData)
		playerDataArray = append(playerDataArray, playerData)
	}
	return playerDataArray
}

//End of inbetween types
