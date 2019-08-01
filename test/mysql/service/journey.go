package test_mysql_interface

import (
	"context"
	"database/sql"

	model "github.com/liupangzi/gpa/test/mysql/model"
)

type Journey interface {
	// @Insert()
	AddCustomer(ctx context.Context, c model.Customer) (sql.Result, error)

	// @Insert()
	AddTicket(ctx context.Context, t *model.Ticket) (sql.Result, error)

	// @Update("UPDATE customer SET name = ? WHERE email = ?")
	UpdateCustomerNameByEmail(ctx context.Context, name, email string) (sql.Result, error)

	// @Delete("DELETE FROM customer WHERE id = ?")
	DeleteCustomerByID(ctx context.Context, ID uint64) (sql.Result, error)

	// @Select("SELECT * FROM customer WHERE id = ?")
	GetCustomerByID(ctx context.Context, ID uint64) (model.Customer, error)

	// @Select("SELECT * FROM ticket WHERE id = ?")
	GetTicketByID(ctx context.Context, ID uint64) (*model.Ticket, error)

	// @Select("SELECT * FROM ticket WHERE customer_id = ? ORDER BY id DESC LIMIT 10")
	GetTicketsByCustomerID(ctx context.Context, customerID uint64) ([]*model.Ticket, error)

	// @Select("SELECT * FROM customer WHERE name = ? AND email = ? LIMIT ?, ?")
	GetCustomersByNameAndEmail(ctx context.Context, name, email string, offset, limit uint64) ([]model.Customer, error)

	// @Select("SELECT `email` FROM customer WHERE name = ? LIMIT 0, 3")
	GetCustomerEmailsByName(ctx context.Context, name string) ([]string, error)

	// @Select("SELECT name FROM `customer` WHERE email = ? LIMIT 1")
	GetCustomerNameByEmail(ctx context.Context, email string) (string, error)
}
