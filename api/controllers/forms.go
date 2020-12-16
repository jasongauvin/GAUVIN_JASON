package controllers

import (
	"github.com/jasongauvin/GAUVIN_JASON/api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/GAUVIN_JASON/api/models"
	"github.com/jasongauvin/GAUVIN_JASON/api/repositories"
)

// CreateForm takes a json object and persists a form into the DB
func CreateForm(c *gin.Context) {
	var err error
	var formType models.FormType

	customer := middleware.GetCustomer(c)
	if customer.Role != "ADMIN_ROLE" {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := c.ShouldBindJSON(&formType); err != nil {
		c.JSON(http.StatusBadRequest, "Incorrect informations")
		return
	}

	// Check form
	err = models.ValidateForm(&formType)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}


	form := models.Form{
		Customer: customer,
		Name: formType.Name,
	}


	if err = repositories.SaveForm(&form); err != nil {
		c.JSON(http.StatusInternalServerError, "Error while saving form informations in database")
		return
	}

	c.JSON(http.StatusOK, "Form successfully created!")
}
