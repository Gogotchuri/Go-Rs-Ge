package responses

type SaveInvoiceItemResponse struct {
	XMLResponse
	Result bool `xml:"Body>save_invoice_descResponse>save_invoice_descResult"`
	ID     int  `xml:"Body>save_invoice_descResponse>id"`
}
