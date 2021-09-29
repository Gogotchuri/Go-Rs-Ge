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
	TIN     int      `xml:"tin"`
	XMLName xml.Name `xml:"get_tin_from_un_id"`
}

func (uift *GetUniqueIDFromTINRequest) GetXML() (*bytes.Buffer, error) {
	uift.Initialize(GetUniqueIDFromTIN)
	return soap.GetXMLBytes(uift)
}
