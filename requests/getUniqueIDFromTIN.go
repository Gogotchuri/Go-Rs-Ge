package requests

import (
	"bytes"
	"encoding/xml"

	"github.com/gogotchuri/go-rs-ge/soap"
)

var _ soap.XMLRequest = &GetUniqueIDFromTINRequest{}

type GetUniqueIDFromTINRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	TIN     string   `xml:"tin"`
	XMLName xml.Name `xml:"get_un_id_from_tin"`
}

func (uift *GetUniqueIDFromTINRequest) GetXML() (*bytes.Buffer, error) {
	uift.Initialize(GetUniqueIDFromTIN)
	return soap.GetXMLBytes(uift)
}
