package requests

import (
	"bytes"
	"encoding/xml"
	"github.com/gogotchuri/go-rs-ge/soap"
)

var _ soap.XMLRequest = &GetInvoiceItemsRequest{}

type GetInvoiceItemsRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	XMLName xml.Name `xml:"get_invoice_desc"`
	InvoiceID int `xml:"invois_id"`
}

func (giir *GetInvoiceItemsRequest) GetXML() (*bytes.Buffer, error) {
	giir.Initialize(GetInvoiceItems)
	return soap.GetXMLBytes(giir)
}


