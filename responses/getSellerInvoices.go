package responses

import "github.com/gogotchuri/go-rs-ge/models"

type GetSellerInvoicesResponse struct {
	XMLResponse
	Invoices []models.Invoice `xml:"Body>get_seller_invoicesResponse>get_seller_invoicesResult>diffgram>DocumentElement>invoices"`
}
