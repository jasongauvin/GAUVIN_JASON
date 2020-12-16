package controllers

import (
	"fmt"
	"net/http"

	"github.com/JackMaarek/spiderMail/services"
	"github.com/jasongauvin/GAUVIN_JASON/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/GAUVIN_JASON/api/models"
	"github.com/jasongauvin/GAUVIN_JASON/api/repositories"
)

// CreateQuestion takes a json object and persists a form into the DB
func CreateQuestion(c *gin.Context) {
	var err error
	var questionForm models.QuestionForm

	form := models.Form{
		ID: services.ConvertStringToInt(c.Param("id")),
	}

	customer := middleware.GetCustomer(c)
	if customer.Role != "ADMIN_ROLE" {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := repositories.FindFormById(&form); err != nil {
		c.JSON(http.StatusBadRequest, "Not a number form for this question")
		return
	}

	if err := c.ShouldBindJSON(&questionForm); err != nil {
		c.JSON(http.StatusBadRequest, "Incorrect informations")
		return
	}

	// Check question
	err = models.ValidateQuestionForm(&questionForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	question := models.Question{
		Form: form,
		Name: questionForm.Name,
		Type: questionForm.Type,
	}

	if err = repositories.SaveQuestion(&question); err != nil {
		c.JSON(http.StatusInternalServerError, "Error while saving form informations in database")
		return
	}

	value := question.Type
	switch value {
	case "multichoice":
		fmt.Print("YEAH")
		choiceList := questionForm.MultiChoiceForm
		for _, choice := range choiceList {
			fmt.Print("CHOICE")
			choice := models.MultiChoice{
				Question: question,
				Name:     choice.Name,
			}

			if err = repositories.SaveChoice(&choice); err != nil {
				c.JSON(http.StatusInternalServerError, "Error while saving form informations in database")
				return
			}
			c.JSON(http.StatusOK, "Choice successfully created!")
		}

	case "slider":
		slider := models.Slider{
			Question: question,
			Min:      questionForm.SliderForm.Min,
			Max:      questionForm.SliderForm.Max,
		}

		if err = repositories.SaveSlider(&slider); err != nil {
			c.JSON(http.StatusInternalServerError, "Error while saving form informations in database")
			return
		}
		c.JSON(http.StatusOK, "Slider successfully created!")
	}

	c.JSON(http.StatusOK, "Question successfully created!")
}
