package models

import (
	"github.com/adrien3d/jump-technical-test/helpers"
	"github.com/adrien3d/jump-technical-test/store"
	mgobson "github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Transaction struct {
	ID        string  `json:"id"`
	InvoiceID string  `json:"invoice_id"`
	Amount    float64 `json:"amount"`
	Reference string  `json:"reference"`
}

// CreateTransaction checks if transaction already exists, and if not, creates it
func CreateTransaction(c *store.Context, transaction *Transaction) error {
	if len(transaction.ID) == 0 {
		transaction.ID = mgobson.NewObjectId().Hex()
	}
	err := c.Store.Create(c, "transactions", transaction)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, "invoice_creation_failed", "Failed to insert the invoice in the database", err)
	}

	return nil
}

// GetTransaction allows to retrieve a transaction by its characteristics
func GetTransaction(c *store.Context, filter bson.M) (*Transaction, error) {
	var transaction Transaction
	err := c.Store.Find(c, filter, &transaction)
	if err != nil {
		return nil, helpers.NewError(http.StatusNotFound, "transaction_not_found", "Transaction not found", err)
	}

	return &transaction, err
}
