package responses

import (
	"github.com/gogotchuri/go-rs-ge/models"
)

type InvoiceResponse struct {
	XMLResponse
	Invoice models.InvoiceSingle `xml:"Body>get_invoiceResponse"`
}
