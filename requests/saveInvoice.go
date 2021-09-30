package requests

import (
	"bytes"
	"encoding/xml"

	"github.com/gogotchuri/go-rs-ge/models"
	"github.com/gogotchuri/go-rs-ge/soap"
)

var _ soap.XMLRequest = &GetInvoiceRequest{}

type SaveInvoiceRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	models.SaveInvoiceT
	XMLName xml.Name `xml:"save_invoice"`
}

func (sir *SaveInvoiceRequest) GetXML() (*bytes.Buffer, error) {
	sir.Initialize(GetInvoice)
	return soap.GetXMLBytes(sir)
}
