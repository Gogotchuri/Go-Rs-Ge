package requests

import (
	"bytes"
	"encoding/xml"
	"github.com/gogotchuri/go-rs-ge/soap"
)

var _ soap.XMLRequest = &GetTINRequest{}

type GetTINRequest struct {
	BaseRequest
	BaseAuthorizedRequest
	XMLName xml.Name `xml:"get_tin_from_un_id"`
}

func (gtr *GetTINRequest) GetXML() (*bytes.Buffer, error) {
	gtr.Initialize(GetTINFromUniqueID)
	return soap.GetXMLBytes(gtr)
}