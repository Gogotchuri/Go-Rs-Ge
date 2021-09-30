package requests

import (
	"bytes"
	"encoding/xml"

	"github.com/gogotchuri/go-rs-ge/models"
	"github.com/gogotchuri/go-rs-ge/soap"
)

var _ soap.XMLRequest = &GetInvoiceRequest{}

type SaveInvoiceItemRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	models.SaveInvoiceItemT
	XMLName xml.Name `xml:"save_invoice_desc"`
}

func (siir *SaveInvoiceItemRequest) GetXML() (*bytes.Buffer, error) {
	siir.Initialize(SaveInvoiceItem)
	return soap.GetXMLBytes(siir)
}
