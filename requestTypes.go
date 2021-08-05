package gorsge

import (
	"bytes"
	"encoding/xml"
)

//XMLRequest interface
type XMLRequest interface {
	getXML() (*bytes.Buffer, error)
	getEndpointURL() string
}

var _ XMLRequest = &GetInvoiceRequest{}

type BaseRequest struct {
	ServiceURL string    `xml:"-"`
	Operation  Operation `xml:"-"`
	XMLns      string    `xml:"xmlns,attr"`
}

func (br *BaseRequest) getEndpointURL() string {
	return br.ServiceURL + string(br.Operation)
}

func (br *BaseRequest) Initialize(op Operation) {
	br.XMLns = "http://tempuri.org/"
	br.Operation = op
	br.ServiceURL = ServiceAddress
}

//BaseAuthorizedRequest contains base properties to be authorized by rs.ge
type BaseAuthorizedRequest struct {
	UserID          int    `xml:"user_id"`
	UniqueID        int    `xml:"un_id"`
	ServiceUser     string `xml:"su"`
	ServicePassword string `xml:"sp"`
}

// ------------ Invoice --------------
type GetInvoiceRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	XMLName   xml.Name `xml:"get_invoice"`
	InvoiceID int      `xml:"invois_id"`
}

func (gir *GetInvoiceRequest) getXML() (*bytes.Buffer, error) {
	gir.Initialize(GetInvoice)
	return getXMLBytes(gir)
}

type GetBuyerInvoicesRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	XMLName xml.Name `xml:"get_buyer_invoices"`
}

// ------------  Users  -----------------
type CredentialCheckRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	XMLName xml.Name `xml:"chek"`
}

func (ccr *CredentialCheckRequest) getXML() (*bytes.Buffer, error) {
	ccr.Initialize(GetInvoice)
	return getXMLBytes(ccr)
}
