package rs

import (
	"fmt"

	"github.com/gogotchuri/go-rs-ge/models"
	"github.com/gogotchuri/go-rs-ge/requests"
	"github.com/gogotchuri/go-rs-ge/responses"
	"github.com/gogotchuri/go-rs-ge/soap"
)

func (rs *Client) GetAllCustomers(filters *models.InvoiceSearchFilters) ([]models.Customer, error) {
	invoices, err := rs.GetSellerInvoices(filters)
	if err != nil {
		return nil, err
	}
	customersMap := make(map[string]*models.Customer, 0)

	for _, inv := range invoices {
		customer, err := rs.GetCustomer(inv.BuyerUID)
		if err != nil {
			continue //TODO consider crashing
		}
		customersMap[customer.TIN] = customer
	}
	customers := make([]models.Customer, 0, len(customersMap))
	for _, v := range customersMap {
		customers = append(customers, *v)
	}
	return customers, nil
}

func (rs *Client) GetAllProducts(filters *models.InvoiceSearchFilters) ([]models.InvoiceItem, error) {
	invoices, err := rs.GetSellerInvoices(filters)
	if err != nil {
		return nil, err
	}
	productsMap := make(map[string]models.InvoiceItem, 0)

	for _, inv := range invoices {
		items, err := rs.GetInvoiceItems(inv.ID)
		if err != nil {
			continue //TODO consider crashing
		}
		for _, item := range items {
			productsMap[item.Name] = item
		}
	}
	products := make([]models.InvoiceItem, 0, len(productsMap))
	for _, v := range productsMap {
		products = append(products, v)
	}
	return products, nil
}

func (rs *Client) GetCustomer(uniqueID int) (*models.Customer, error) {
	return rs.GetTINFromUniqueID(uniqueID)
}

func (rs *Client) GetTINFromUniqueID(uniqueID int) (*models.Customer, error) {
	gtr := &requests.GetTINRequest{}
	gtr.ServiceUser = rs.ServiceUser
	gtr.ServicePassword = rs.ServicePassword
	gtr.UserID = rs.UserID
	gtr.UniqueID = uniqueID
	req, err := soap.GenerateSOAPRequest(gtr)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return nil, err
	}
	ir := &responses.GetTINResponse{}
	_, err = soap.SoapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return nil, err
	}
	return &ir.Customer, nil
}

func (rs *Client) GetUniqueIDFromTIN(TIN string) (*models.CustomerM, error) {
	guift := &requests.GetUniqueIDFromTINRequest{}
	guift.ServiceUser = rs.ServiceUser
	guift.ServicePassword = rs.ServicePassword
	guift.UserID = rs.UserID
	guift.TIN = TIN
	req, err := soap.GenerateSOAPRequest(guift)
	if err != nil {
		fmt.Println("Generating request:", err.Error())
		return nil, err
	}
	ir := &responses.GetUniqueIDFromTINResponse{}
	_, err = soap.SoapCall(req, ir)
	if err != nil {
		fmt.Println("Making call", err.Error())
		return nil, err
	}

	return &models.CustomerM{
		TIN:      TIN,
		Name:     ir.Name,
		UniqueID: ir.UniqueID,
	}, nil
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

//GetUniqueID returns unique id of current rs user
//This automatically sets user_id and unique_id to rs client
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
	if rs.UniqueID <= 0 {
		return ir.UniqueID, fmt.Errorf("unique id is <= 0")
	}
	return ir.UniqueID, nil
}
