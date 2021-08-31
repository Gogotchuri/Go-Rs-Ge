package responses

import (
	"encoding/xml"
	"github.com/gogotchuri/go-rs-ge/models"
)

type CheckResponse struct {
	XMLName         xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	CredentialCheck models.CredentialCheck `xml:"Body>chekResponse"`
}

type GetUniqueIDResponse struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	UniqueID int  `xml:"Body>get_un_id_from_user_idResponse>get_un_id_from_user_idResult"`
}