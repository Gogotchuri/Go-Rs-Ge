package requests

import (
	"bytes"
	"encoding/xml"

	"github.com/gogotchuri/go-rs-ge/models"
	"github.com/gogotchuri/go-rs-ge/soap"
)

var _ soap.XMLRequest = &ChangeInvoiceStatusRequest{}

type ChangeInvoiceStatusRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	XMLName   xml.Name            `xml:"change_invoice_status"`
	Status    models.InvoiceSatus `xml:"status"`
	InvoiceID int                 `xml:"inv_id"`
}

func (cisr *ChangeInvoiceStatusRequest) GetXML() (*bytes.Buffer, error) {
	cisr.Initialize(ChangeInvoiceStatus)
	return soap.GetXMLBytes(cisr)
}
