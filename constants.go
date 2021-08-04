package gorsge

const ServiceAddress = "https://www.revenue.mof.ge/ntosservice/ntosservice.asmx?op="

const RsGeXMLTemplate = `
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		{{.OperationBody}}
	</soap:Body>
	</soap:Envelope>
`

//Operations
const (
	WhatIsMyIP Operation = "what_is_my_ip"
	GetInvoice Operation = "get_invoice"
)
