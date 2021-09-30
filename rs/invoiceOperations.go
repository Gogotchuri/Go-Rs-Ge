package rs

import (
	"fmt"

	"github.com/gogotchuri/go-rs-ge/models"
	"github.com/gogotchuri/go-rs-ge/requests"
	"github.com/gogotchuri/go-rs-ge/responses"
	"github.com/gogotchuri/go-rs-ge/soap"
)

func (rs *Client) GetBuyerInvoices(filters *models.InvoiceSearchFilters) ([]models.Invoice, error) {
	gir := &requests.GetBuyerInvoicesRequest{}
	if filters != nil {
		gir.SetFilters(filters)
	}
	gir.ServiceUser = rs.ServiceUser
	gir.ServicePassword = rs.ServicePassword
	gir.UserID = rs.UserID
	gir.UniqueID = rs.UniqueID
	req, err := soap.GenerateSOAPRequest(gir)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return nil, err
	}
	ir := &responses.GetBuyerInvoicesResponse{}
	_, err = soap.SoapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return nil, err
	}
	return ir.Invoices, nil
}

func (rs *Client) GetSellerInvoices(filters *models.InvoiceSearchFilters) ([]models.Invoice, error) {
	gsir := &requests.GetSellerInvoicesRequest{}
	if filters != nil {
		gsir.SetFilters(filters)
	}
	gsir.ServiceUser = rs.ServiceUser
	gsir.ServicePassword = rs.ServicePassword
	gsir.UserID = rs.UserID
	gsir.UniqueID = rs.UniqueID
	req, err := soap.GenerateSOAPRequest(gsir)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return nil, err
	}
	ir := &responses.GetSellerInvoicesResponse{}
	_, err = soap.SoapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return nil, err
	}
	return ir.Invoices, nil
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
	ir.Invoice.ID = invoiceID
	return ir.Invoice.GetInvoice(), nil
}

func (rs *Client) GetInvoiceItems(invoiceID int) ([]models.InvoiceItem, error) {
	giir := &requests.GetInvoiceItemsRequest{
		InvoiceID: invoiceID,
	}
	giir.ServiceUser = rs.ServiceUser
	giir.ServicePassword = rs.ServicePassword
	giir.UserID = rs.UserID
	giir.UniqueID = rs.UniqueID
	req, err := soap.GenerateSOAPRequest(giir)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return nil, err
	}
	ir := &responses.GetInvoiceItemsResponse{}
	_, err = soap.SoapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return nil, err
	}
	return ir.InvoiceItems, nil
}

func (rs *Client) SaveInvoice(toSave models.SaveInvoiceT) (int, error) {
	sir := &requests.SaveInvoiceRequest{
		SaveInvoiceT: toSave,
	}
	sir.ServiceUser = rs.ServiceUser
	sir.ServicePassword = rs.ServicePassword
	sir.UserID = rs.UserID
	req, err := soap.GenerateSOAPRequest(sir)
	fmt.Println(req)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return 0, err
	}
	ir := &responses.SaveInvoiceResponse{}
	_, err = soap.SoapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return 0, err
	}
	if ir.InvoiceID <= 0 {
		return 0, fmt.Errorf("invoice id returned 0 (not created)")
	}

	return ir.InvoiceID, nil
}

func (rs *Client) SaveInvoiceItem(toSave models.SaveInvoiceItemT) (int, error) {
	siir := &requests.SaveInvoiceItemRequest{
		SaveInvoiceItemT: toSave,
	}
	siir.ServiceUser = rs.ServiceUser
	siir.ServicePassword = rs.ServicePassword
	siir.UserID = rs.UserID
	req, err := soap.GenerateSOAPRequest(siir)
	fmt.Println(req)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return 0, err
	}
	ir := &responses.SaveInvoiceItemResponse{}
	_, err = soap.SoapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return 0, err
	}
	if ir.ID <= 0 {
		return 0, fmt.Errorf("invoice item id returned 0 (not created)")
	}
	return ir.ID, nil
}
