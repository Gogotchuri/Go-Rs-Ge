package requests

import (
	"bytes"
	"encoding/xml"

	"github.com/gogotchuri/go-rs-ge/soap"
)

type GetSellerInvoicesRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	GetInvoicesFilters
	XMLName xml.Name `xml:"get_seller_invoices"`
}

func (gsir *GetSellerInvoicesRequest) GetXML() (*bytes.Buffer, error) {
	gsir.Initialize(GetSellerInvoices)
	return soap.GetXMLBytes(gsir)
}
