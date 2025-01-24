package repository

import (
	"context"
	"database/sql"
	"yogolibrary-gofiber-rest-api/domain"

	"github.com/doug-martin/goqu/v9"
)

type customerRepostory struct {
	db *goqu.Database
}

// FindAll implements domain.CutomerRepository.
func (c *customerRepostory) FindAll(ctx context.Context) ([]domain.Customer, error) {
	panic("unimplemented")
}

// FindById implements domain.CutomerRepository.
func (c *customerRepostory) FindById(ctx context.Context, id string) (domain.Customer, error) {
	panic("unimplemented")
}

// Save implements domain.CutomerRepository.
func (c *customerRepostory) Save(ctx context.Context, cust *domain.Customer) error {
	panic("unimplemented")
}

// Update implements domain.CutomerRepository.
func (c *customerRepostory) Update(ctx context.Context, cust *domain.Customer) error {
	panic("unimplemented")
}

// Delete implements domain.CutomerRepository.
func (c *customerRepostory) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func NewCustomer(conn *sql.DB) domain.CutomerRepository {
	return &customerRepostory{
		db: goqu.New("default", conn),
	}
}
