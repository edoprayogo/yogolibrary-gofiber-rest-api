package repository

import (
	"context"
	"database/sql"
	"time"
	"yogolibrary-gofiber-rest-api/domain"

	"github.com/doug-martin/goqu/v9"
)

type customerRepostory struct {
	db *goqu.Database
}

func NewCustomer(conn *sql.DB) domain.CutomerRepository {
	return &customerRepostory{
		db: goqu.New("default", conn),
	}
}

// FindAll implements domain.CutomerRepository.
func (c *customerRepostory) FindAll(ctx context.Context) (result []domain.Customer, err error) {
	//panic("unimplemented")

	dataset := c.db.From("customers").Where(goqu.C("deleted_at").IsNull())

	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// FindById implements domain.CutomerRepository.
func (c *customerRepostory) FindById(ctx context.Context, id string) (result domain.Customer, err error) {
	//panic("unimplemented")
	dataset := c.db.From("customers").
		Where(goqu.C("deleted_at").IsNull(),
			goqu.C("id").Eq(id))

	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// Save implements domain.CutomerRepository.
func (c *customerRepostory) Save(ctx context.Context, cust *domain.Customer) error {

	excutor := c.db.Insert("customers").
		Rows(cust).
		Executor()
	_, err := excutor.ExecContext(ctx)
	return err
}

// Update implements domain.CutomerRepository.
func (c *customerRepostory) Update(ctx context.Context, cust *domain.Customer) error {
	//panic("unimplemented")
	excutor := c.db.Update("customers").
		Where(goqu.C("id").Eq(cust.ID)).
		Set(c).
		Executor()

	_, err := excutor.ExecContext(ctx)
	return err
}

// Delete implements domain.CutomerRepository.
func (c *customerRepostory) Delete(ctx context.Context, id string) error {
	//panic("unimplemented")
	excutor := c.db.Update("customers").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{
			"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).
		Executor()

	_, err := excutor.ExecContext(ctx)
	return err
}
