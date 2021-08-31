package responses

import (
	"github.com/gogotchuri/go-rs-ge/models"
)

type InvoiceResponse struct {
	XMLResponse
	Invoice models.Invoice  `xml:"Body>get_invoiceResponse"`
}