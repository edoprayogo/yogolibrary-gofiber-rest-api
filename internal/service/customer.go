package service

import (
	"context"
	"yogolibrary-gofiber-rest-api/domain"
	"yogolibrary-gofiber-rest-api/dto"
)

type customerService struct {
	customerRepostory domain.CutomerRepository
}

// Index implements domain.CustomerService.
func (c *customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepostory.FindAll(ctx)

	if err != nil {
		return nil, err
	}
	var customerData []dto.CustomerData
	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}
	return customerData, nil
}

func NewCustomer(customerRepostory domain.CutomerRepository) domain.CustomerService {
	return &customerService{
		customerRepostory: customerRepostory,
	}
}
