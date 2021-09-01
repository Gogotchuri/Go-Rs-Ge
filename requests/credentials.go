package requests

import (
	"bytes"
	"encoding/xml"

	"github.com/gogotchuri/go-rs-ge/soap"
)

var _ soap.XMLRequest = &GetUniqueIDRequest{}
var _ soap.XMLRequest = &CredentialCheckRequest{}

//GetUniqueIDRequest for operation GetUnIDFromUserID
type GetUniqueIDRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	XMLName xml.Name `xml:"get_un_id_from_user_id"`
}

func (guir *GetUniqueIDRequest) GetXML() (*bytes.Buffer, error) {
	guir.Initialize(GetUniqueID)
	return soap.GetXMLBytes(guir)
}

// CredentialCheckRequest .
type CredentialCheckRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	XMLName xml.Name `xml:"chek"`
}

func (ccr *CredentialCheckRequest) GetXML() (*bytes.Buffer, error) {
	ccr.Initialize(CheckCredentials)
	return soap.GetXMLBytes(ccr)
}
