package responses

type SaveInvoiceResponse struct {
	XMLResponse
	Result    bool `xml:"Body>save_invoiceResponse>save_invoiceResult"`
	InvoiceID int  `xml:"Body>save_invoiceResponse>invois_id"`
}
