package database

import (
	"github.com/jasongauvin/GAUVIN_JASON/api/models"
)

func MakeMigrations() {
	DB.AutoMigrate(&models.Customer{}, &models.Form{})
}
