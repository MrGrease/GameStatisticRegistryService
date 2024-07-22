package models

type GameStats interface {
	ParseJsonData(rawData map[string]interface{}) error
}
