package controllers

import (
	"github.com/adrien3d/jump-technical-test/helpers"
	"github.com/adrien3d/jump-technical-test/models"
	"github.com/adrien3d/jump-technical-test/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

// TransactionController holds all controller functions related to the transaction entity
type TransactionController struct {
	BaseController
}

// NewTransactionController instantiates the controller
func NewTransactionController() TransactionController {
	return TransactionController{}
}

// CreateTransaction to create a new transaction
func (tc TransactionController) CreateTransaction(c *gin.Context) {
	ctx := store.AuthContext(c)
	transactionInput := &models.Transaction{}

	if err := c.BindJSON(transactionInput); err != nil {
		tc.AbortWithError(c, helpers.ErrorInvalidInput(err))
		return
	}

	if _, userGroup, ok := tc.LoggedUser(c); ok {
		switch userGroup.GetRole() {
		case store.RoleGod:
			if invoice, err := models.GetInvoice(ctx, bson.M{"id": transactionInput.InvoiceID}); err != nil {
				//No invoice found => 404
				tc.AbortWithError(c, helpers.ErrorResourceNotFound(err))
			} else if invoice.Amount != transactionInput.Amount {
				//Amount NOK => 400
				tc.AbortWithError(c, helpers.ValidationError(err))
			} else {
				if transaction, err := models.GetTransaction(ctx, bson.M{"invoice_id": transactionInput.InvoiceID}); err == nil && len(transaction.ID) > 0 {
					//Invoice already paid => 422
					tc.AbortWithError(c, helpers.ErrorUnprocessableEntity(err))
				} else {
					//Credit user balance
					if user, err := models.GetUser(ctx, bson.M{"id": invoice.UserID}); err != nil {
						tc.AbortWithError(c, helpers.ErrorResourceNotFound(err))
					} else {
						user.Balance += invoice.Amount
						if err := models.UpdateUser(ctx, user.ID, user); err != nil {
							tc.AbortWithError(c, helpers.ErrorInternal(err))
						} else {
							c.JSON(http.StatusNoContent, transaction)
						}
					}
				}
			}

		case store.RoleAdmin, store.RoleUser, store.RoleCustomer:
			tc.AbortWithError(c, helpers.ErrorUserUnauthorized)
		}
	}
}
