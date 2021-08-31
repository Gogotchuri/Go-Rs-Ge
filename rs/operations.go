package rs

import (
	"fmt"
	"github.com/gogotchuri/go-rs-ge/models"
	"github.com/gogotchuri/go-rs-ge/requests"
	"github.com/gogotchuri/go-rs-ge/responses"
	"github.com/gogotchuri/go-rs-ge/soap"
)

func (rs *Client) GetBuyerInvoices(invoiceID int) ([]*models.Invoice, error) {
	gir := &requests.GetBuyerInvoicesRequest{}
	gir.ServiceUser = rs.ServiceUser
	gir.ServicePassword = rs.ServicePassword
	gir.UserID = rs.UserID
	req, err := soap.GenerateSOAPRequest(gir)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return nil, err
	}
	ir := &responses.InvoiceResponse{}
	_, err = soap.SoapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return nil, err
	}
	ir.Invoice.InvoiceID = invoiceID
	return nil, nil
}

func (rs *Client) GetInvoice(invoiceID int) (*models.Invoice, error) {
	gir := &requests.GetInvoiceRequest{
		InvoiceID: invoiceID,
	}
	gir.ServiceUser = rs.ServiceUser
	gir.ServicePassword = rs.ServicePassword
	gir.UserID = rs.UserID
	req, err := soap.GenerateSOAPRequest(gir)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return nil, err
	}
	ir := &responses.InvoiceResponse{}
	_, err = soap.SoapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return nil, err
	}
	ir.Invoice.InvoiceID = invoiceID
	return &ir.Invoice, nil
}

func (rs *Client) CheckCredentials() (*models.CredentialCheck, error) {
	ccr := &requests.CredentialCheckRequest{}
	ccr.ServiceUser = rs.ServiceUser
	ccr.ServicePassword = rs.ServicePassword
	ccr.UserID = rs.UserID
	req, err := soap.GenerateSOAPRequest(ccr)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return nil, err
	}
	ir := &responses.CheckResponse{}
	_, err = soap.SoapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return nil, err
	}
	rs.UserID = ir.CredentialCheck.UserID
	return &ir.CredentialCheck, nil
}

func (rs *Client) GetUniqueID() (int, error) {
	gui := &requests.GetUniqueIDRequest{}
	gui.ServiceUser = rs.ServiceUser
	gui.ServicePassword = rs.ServicePassword
	if rs.UserID <= 0 {
		creds, err := rs.CheckCredentials()
		if err != nil {
			return 0, err
		}
		rs.UserID = creds.UserID
	}
	gui.UserID = rs.UserID
	req, err := soap.GenerateSOAPRequest(gui)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return 0, err
	}
	ir := &responses.GetUniqueIDResponse{}
	_, err = soap.SoapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return 0, err
	}
	rs.UniqueID = ir.UniqueID
	return ir.UniqueID, nil
}