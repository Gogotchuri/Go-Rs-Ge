package requests

import (
	"bytes"
	"encoding/xml"

	"github.com/gogotchuri/go-rs-ge/models"
	"github.com/gogotchuri/go-rs-ge/soap"
)

var _ soap.XMLRequest = &GetBuyerInvoicesRequest{}
var _ soap.XMLRequest = &GetSellerInvoicesRequest{}

type GetInvoicesFilters struct {
	StartDate          string `xml:"s_dt"`
	EndDate            string `xml:"e_dt"`
	OperationStartDate string `xml:"op_s_dt"`
	OperationEndDate   string `xml:"op_e_dt"`

	InvoiceNo  string `xml:"invoice_no"`
	CompanyTIN string `xml:"sa_ident_no"`
	Desc       string `xml:"desc"`
}

func (gif *GetInvoicesFilters) SetFilters(filters *models.InvoiceSearchFilters) {
	gif.StartDate = filters.StartDate.String()[0:10]
	gif.EndDate = filters.EndDate.String()[0:10]
	gif.OperationStartDate = filters.OperationStartDate.String()[0:10]
	gif.OperationEndDate = filters.OperationEndDate.String()[0:10]

	gif.InvoiceNo = filters.InvoiceNo
	gif.CompanyTIN = filters.CompanyTIN
	gif.Desc = filters.Desc
}

type GetBuyerInvoicesRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	GetInvoicesFilters
	XMLName xml.Name `xml:"get_buyer_invoices"`
}

func (gbir *GetBuyerInvoicesRequest) GetXML() (*bytes.Buffer, error) {
	gbir.Initialize(GetBuyerInvoices)
	return soap.GetXMLBytes(gbir)
}
