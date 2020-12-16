package repositories

import (
	"errors"

	"github.com/jasongauvin/GAUVIN_JASON/api/database"
	"github.com/jasongauvin/GAUVIN_JASON/api/models"
	"github.com/jinzhu/gorm"
)

// SaveQuestion persist a question in the db
func SaveQuestion(question *models.Question) error {
	err := database.DB.Debug().Create(&question).Error
	if err != nil {
		return err
	}
	return nil
}

// SaveSlider persist a slider in the db
func SaveSlider(slider *models.Slider) error {
	err := database.DB.Debug().Create(&slider).Error
	if err != nil {
		return err
	}
	return nil
}

// SaveChoice persist a choice in the db
func SaveChoice(choice *models.MultiChoice) error {
	err := database.DB.Debug().Create(&choice).Error
	if err != nil {
		return err
	}
	return nil
}
// FindQuestionByID finds a question in the db
func FindQuestionByID(question *models.Question) error {
	err := database.DB.Debug().Where("id = ?", question.ID).First(&question).Error
	if err != nil {
		return  err
	}
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Question Not Found")
	}
	return err
}