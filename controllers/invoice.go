package controllers

import (
	"github.com/adrien3d/jump-technical-test/helpers"
	"github.com/adrien3d/jump-technical-test/models"
	"github.com/adrien3d/jump-technical-test/store"
	"github.com/adrien3d/jump-technical-test/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InvoiceController holds all controller functions related to the invoice entity
type InvoiceController struct {
	BaseController
}

// NewInvoiceController instantiates the controller
func NewInvoiceController() InvoiceController {
	return InvoiceController{}
}

// CreateInvoice to create a new invoice
func (ic InvoiceController) CreateInvoice(c *gin.Context) {
	ctx := store.AuthContext(c)
	invoice := &models.Invoice{}

	if err := c.BindJSON(invoice); err != nil {
		ic.AbortWithError(c, helpers.ErrorInvalidInput(err))
		return
	}

	if _, userGroup, ok := ic.LoggedUser(c); ok {
		switch userGroup.GetRole() {
		case store.RoleGod:
			if err := models.CreateInvoice(ctx, invoice); err != nil {
				ic.AbortWithError(c, helpers.ErrorInternal(err))
				return
			}
			c.JSON(http.StatusNoContent, invoice)
		case store.RoleAdmin, store.RoleUser, store.RoleCustomer:
			ic.AbortWithError(c, helpers.ErrorUserUnauthorized)
		}
	} else {
		utils.Log(c, "warn", ok)
	}
}
