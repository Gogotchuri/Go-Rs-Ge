package responses

type ChangeInvoiceStatusResponse struct {
	XMLResponse
	Result bool `xml:"Body>change_invoice_statusResponse>change_invoice_statusResult"`
}
