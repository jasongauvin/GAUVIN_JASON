package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

// Form is the struct send on database
type Form struct {
	ID         uint64 `gorm:"primary_key"`
	CustomerID uint64
	Customer   Customer
	Name       string `gorm:"size:255"`
}

// FormType is the struct send by customer for use form
type FormType struct {
	Name string `gorm:"size:255"`
}

// ValidateForm check if its properties are valid
func ValidateForm(formType *FormType) error {
	_, err := govalidator.ValidateStruct(formType)
	if err != nil {
		return err
	}
	if formType.Name == "" {
		return errors.New("Name of form is empty")
	}

	return nil
}
