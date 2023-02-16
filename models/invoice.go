package models

import (
	"github.com/adrien3d/jump-technical-test/helpers"
	"github.com/adrien3d/jump-technical-test/store"
	mgobson "github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Invoice struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
	Label  string  `json:"label"`
}

// CreateInvoice checks if invoice already exists, and if not, creates it
func CreateInvoice(c *store.Context, invoice *Invoice) error {
	/*var existingInvoices []*Group
	err := c.Store.FindAll(c, bson.M{"name": invoice.Amount}, &existingInvoices)
	if err != nil {
		return err
	}

	if len(existingInvoices) > 0 {
		return helpers.NewError(http.StatusConflict, "group_already_exists", "Group already exists", err)
	}*/

	if len(invoice.ID) == 0 {
		invoice.ID = mgobson.NewObjectId().Hex()
	}
	err := c.Store.Create(c, "invoices", invoice)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, "invoice_creation_failed", "Failed to insert the invoice in the database", err)
	}

	return nil
}

// GetInvoice allows to retrieve an invoice by its characteristics
func GetInvoice(c *store.Context, filter bson.M) (*Invoice, error) {
	var invoice Invoice
	err := c.Store.Find(c, filter, &invoice)
	if err != nil {
		return nil, helpers.NewError(http.StatusNotFound, "invoice_not_found", "Invoice not found", err)
	}

	return &invoice, err
}
