package responses

import (
	"github.com/gogotchuri/go-rs-ge/models"
)

type GetBuyerInvoicesResponse struct {
	XMLResponse
	Invoices []*models.Invoice `xml:"Body>get_buyer_invoicesResponse>get_buyer_invoicesResult>diffgram>DocumentElement>invoices"`
}
