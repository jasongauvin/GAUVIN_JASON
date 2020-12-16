package repositories

import (
	"errors"

	"github.com/jasongauvin/GAUVIN_JASON/api/database"
	"github.com/jasongauvin/GAUVIN_JASON/api/models"
	"github.com/jinzhu/gorm"
)

func CreateCustomer(customer *models.Customer) error {
	err := database.DB.Debug().Create(&customer).Error
	if err != nil {
		return err
	}
	return nil
}

// FindCustomerByEmail finds a customer in the db
func FindCustomerByEmail(customer *models.Customer) error {
	err := database.DB.Debug().Where("email = ?", customer.Email).First(&customer).Error
	if err != nil {
		return  err
	}
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Customer Not Found")
	}
	return err
}
