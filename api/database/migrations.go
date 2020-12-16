package database

import (
	"github.com/jasongauvin/GAUVIN_JASON/api/models"
)

// MakeMigrations load sata structure on db
func MakeMigrations() {
	DB.AutoMigrate(&models.Customer{}, &models.Form{}, &models.Question{}, &models.MultiChoice{}, &models.Slider{})
}
