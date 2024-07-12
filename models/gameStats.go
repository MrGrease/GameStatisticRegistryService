package models

type GameStats interface {
	ParseJsonData(rawData string)
}
