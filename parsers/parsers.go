package parsers

import (
	"gamestatsticregistry/mrgrease.com/models"
)

func GetAppParsedType(appName string) models.GameStats {
	switch appName {
	case "RenegadeX":
		return new(models.RenegadeXStats)
	}

	panic("Undefined Type")
}
