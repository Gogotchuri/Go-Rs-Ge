package requests

import (
	"bytes"
	"encoding/xml"

	"github.com/gogotchuri/go-rs-ge/soap"
)

var _ soap.XMLRequest = &GetInvoiceRequest{}

type GetInvoiceRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	XMLName   xml.Name `xml:"get_invoice"`
	InvoiceID int      `xml:"invois_id"`
}

func (gir *GetInvoiceRequest) GetXML() (*bytes.Buffer, error) {
	gir.Initialize(GetInvoice)
	return soap.GetXMLBytes(gir)
}
