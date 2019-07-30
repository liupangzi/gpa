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
}
