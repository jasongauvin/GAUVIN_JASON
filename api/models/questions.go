package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

// Question is the struct send on database
type Question struct {
	ID     uint64 `gorm:"primary_key"`
	FormID uint64
	Form   Form
	Name   string `gorm:"size:255"`
	Type   string `gorm:"size:255"`
}

// QuestionForm is the struct send by customer for use form
type QuestionForm struct {
	FormID uint64
	Name   string `gorm:"size:255"`
	Type   string `gorm:"size:255"`
	MultiChoiceForm []MultiChoiceForm
	SliderForm SliderForm

}

// MultiChoice is the struct send on database
type MultiChoice struct {
	ID         uint64 `gorm:"primary_key"`
	QuestionID uint64
	Question   Question
	Name       string `gorm:"size:255"`
}

// MultiChoiceForm is the struct send by customer for use form
type MultiChoiceForm struct {
	Name string `gorm:"size:255"`
}

// Slider is the struct send on database
type Slider struct {
	ID         uint64 `gorm:"primary_key"`
	QuestionID uint64
	Question   Question
	Name       string `gorm:"size:255"`
	Min        int
	Max        int
}

// SliderForm is the struct send by customer for use form
type SliderForm struct {
	Name string `gorm:"size:255"`
	Min  int
	Max  int
}

// ValidateQuestionForm check if its properties are valid
func ValidateQuestionForm(questionFrom *QuestionForm) error {
	_, err := govalidator.ValidateStruct(questionFrom)
	if err != nil {
		return err
	}
	if questionFrom.Name == "" {
		return errors.New("Name of question is empty")
	}

	return nil
}

// ValidateMultiChoiceForm check if its properties are valid
func ValidateMultiChoiceForm(multipleChoiceForm *MultiChoiceForm) error {
	_, err := govalidator.ValidateStruct(multipleChoiceForm)
	if err != nil {
		return err
	}
	if multipleChoiceForm.Name == "" {
		return errors.New("Name of choise is empty")
	}

	return nil
}

// ValidateSliderForm check if its properties are valid
func ValidateSliderForm(sliderForm *SliderForm) error {
	_, err := govalidator.ValidateStruct(sliderForm)
	if err != nil {
		return err
	}
	if sliderForm.Name == "" {
		return errors.New("Name of slider is empty")
	}

	return nil
}
