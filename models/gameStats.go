package models

type GameStats interface {
	Parse(rawData []byte) error
}
