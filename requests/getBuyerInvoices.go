package requests

import (
	"bytes"
	"encoding/xml"
	"github.com/gogotchuri/go-rs-ge/soap"
)

var _ soap.XMLRequest = &GetBuyerInvoicesRequest{}

type GetBuyerInvoicesRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	XMLName xml.Name `xml:"get_buyer_invoices"`
}

func (gbir *GetBuyerInvoicesRequest) GetXML() (*bytes.Buffer, error) {
	gbir.Initialize(CheckCredentials)
	return soap.GetXMLBytes(gbir)
}
