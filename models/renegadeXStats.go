package models

import "fmt"

//Inbetween types used to pass data around until we decide what DB to use
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

func (statsData *RenegadeXStats) ParseJsonData(rawData map[string]interface{}) {
	fmt.Println("Attempting to parse raw data")
	//statsData.Winner = int16(rawData["winner"]) How do we handle this?
}

//End of inbetween types
