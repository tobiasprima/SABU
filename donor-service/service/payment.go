package service

import (
	"context"
	"donor-service/models"

	xendit "github.com/xendit/xendit-go/v6"
	invoice "github.com/xendit/xendit-go/v6/invoice"
)

type PaymentService interface {
	CreateInvoice(externalID string, amount float64) (*models.Invoice, error)
}

type XenditClient struct {
	Client *xendit.APIClient
}

func NewPaymentService(key string) PaymentService {
	return &XenditClient{Client: xendit.NewClient(key)}
}

func (xc *XenditClient) CreateInvoice(externalID string, amount float64) (*models.Invoice, error) {
	createInvoiceRequest := *invoice.NewCreateInvoiceRequest(externalID, amount)

	resp, _, err := xc.Client.InvoiceApi.CreateInvoice(context.Background()).
		CreateInvoiceRequest(createInvoiceRequest).
		Execute()
	if err != nil {
		return nil, err
	}

	invoiceData := &models.Invoice{
		InvoiceID:  *resp.Id,
		InvoiceUrl: resp.InvoiceUrl,
		Amount:     resp.Amount,
		Method:     "PENDING",
		Status:     resp.Status.String(),
	}

	return invoiceData, nil
}
