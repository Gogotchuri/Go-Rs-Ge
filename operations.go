package gorsge

import "fmt"

func (rs *RsGeClient) GetInvoice(invoiceID int) (*Invoice, error) {
	gir := &GetInvoiceRequest{
		InvoiceID: invoiceID,
	}
	gir.ServiceUser = rs.ServiceUser
	gir.ServicePassword = rs.ServicePassword
	gir.UserID = rs.UserID
	req, err := generateSOAPRequest(gir)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return nil, err
	}
	ir := &invoiceResponse{}
	_, err = soapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return nil, err
	}
	ir.Invoice.InvoiceID = invoiceID
	return &ir.Invoice, nil
}

func (rs *RsGeClient) CheckCredentials() (*CredentialCheck, error) {
	ccr := &CredentialCheckRequest{}
	ccr.ServiceUser = rs.ServiceUser
	ccr.ServicePassword = rs.ServicePassword
	ccr.UserID = rs.UserID
	req, err := generateSOAPRequest(ccr)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return nil, err
	}
	ir := &checkResponse{}
	_, err = soapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return nil, err
	}
	rs.UserID = ir.CredentialCheck.UserID
	return &ir.CredentialCheck, nil
}
