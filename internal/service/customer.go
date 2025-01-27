package service

import (
	"context"
	"database/sql"
	"time"
	"yogolibrary-gofiber-rest-api/domain"
	"yogolibrary-gofiber-rest-api/dto"

	"github.com/google/uuid"
)

type customerService struct {
	customerRepostory domain.CutomerRepository
}

func NewCustomer(customerRepostory domain.CutomerRepository) domain.CustomerService {
	return &customerService{
		customerRepostory: customerRepostory,
	}
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

func (c customerService) Create(ctx context.Context, request dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:        uuid.New().String(),
		Code:      request.Code,
		Name:      request.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
		CreatedBy: sql.NullString{String: request.CreatedBy},
	}

	return c.customerRepostory.Save(ctx, &customer)
}
