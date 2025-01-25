package domain

import (
	"context"
	"database/sql"
	"yogolibrary-gofiber-rest-api/dto"
)

type Customer struct {
	ID        string       `db:"id"`
	Code      string       `db:code`
	Name      string       `db:name`
	CreatedAt sql.NullTime `db:created_at`
	UpdateAt  sql.NullTime `db:updated_at`
	CreatedBy string       `db:created_by`
	UpdateBy  string       `db:updated_by`
	DeletedAt sql.NullTime `db:deleted_at`
}

type CutomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindById(ctx context.Context, id string) (Customer, error)
	Save(ctx context.Context, cust *Customer) error
	Update(ctx context.Context, cust *Customer) error
	Delete(ctx context.Context, id string) error
}

type CustomerService interface {
	Index(ctx context.Context) ([]dto.CustomerData, error)
}
