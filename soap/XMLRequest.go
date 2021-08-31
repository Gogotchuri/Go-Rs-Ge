package soap

import "bytes"

//XMLRequest interface
type XMLRequest interface {
	GetXML() (*bytes.Buffer, error)
	GetEndpointURL() string
}

const RsGeXMLTemplate = `
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		{{.OperationBody}}
	</soap:Body>
	</soap:Envelope>
`