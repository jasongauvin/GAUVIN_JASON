package repositories

import (
	"errors"

	"github.com/jasongauvin/GAUVIN_JASON/api/database"
	"github.com/jasongauvin/GAUVIN_JASON/api/models"
	"github.com/jinzhu/gorm"
)

func SaveForm(form *models.Form) error {
	err := database.DB.Debug().Create(&form).Error
	if err != nil {
		return err
	}
	return nil
}

// FindFormById finds a form in the db
func FindFormById(form *models.Form) error {
	err := database.DB.Debug().Where("id = ?", form.ID).First(&form).Error
	if err != nil {
		return  err
	}
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Form Not Found")
	}
	return err
}

// FindFormByUserId finds a form by User in the db
func FindFormByCustomerId(form *models.Form) error {
	err := database.DB.Debug().Where("id = ?", form.CustomerID).First(&form).Error
	if err != nil {
		return  err
	}
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Form by CustomerID Not Found")
	}
	return err
}