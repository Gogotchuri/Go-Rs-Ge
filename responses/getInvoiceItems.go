package responses

import "github.com/gogotchuri/go-rs-ge/models"

type GetInvoiceItemsResponse struct {
	XMLResponse
	InvoiceItems []models.InvoiceItem `xml:"Body>get_invoice_descResponse>get_invoice_descResult>diffgram>DocumentElement>invoices_descs"`
}
