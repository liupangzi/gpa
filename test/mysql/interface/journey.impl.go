// Code generated from journey.go by gpa. DO NOT EDIT.

package test_mysql_interface

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	model "github.com/liupangzi/gpa/test/mysql/model"
	"reflect"
	"strings"
)

type JourneyImpl struct {
	DB *sqlx.DB
}

type TxJourneyImpl struct {
	TX *sqlx.Tx
}

func NewJourney(db *sqlx.DB) *JourneyImpl {
	return &JourneyImpl{
		DB: db,
	}
}

func (db *JourneyImpl) Begin(ctx context.Context, options *sql.TxOptions) *TxJourneyImpl {
	return &TxJourneyImpl{
		TX: db.DB.MustBeginTx(ctx, options),
	}
}

func (tx *TxJourneyImpl) Commit(ctx context.Context) error {
	return tx.TX.Commit()
}

func (tx *TxJourneyImpl) Rollback(ctx context.Context) error {
	return tx.TX.Rollback()
}

func (db *JourneyImpl) AddCustomer(ctx context.Context, c model.Customer) (sql.Result, error) {
	columnNames, values := make([]string, 0), make([]string, 0)
	for i := 0; i < reflect.TypeOf(c).NumField(); i++ {
		columnNames = append(columnNames, "`"+reflect.TypeOf(c).Field(i).Tag.Get("db")+"`")
		values = append(values, ":"+reflect.TypeOf(c).Field(i).Tag.Get("db"))
	}
	return db.DB.NamedExec(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", c.TableName(), strings.Join(columnNames, ", "), strings.Join(values, ", ")), &c)
}

func (tx *TxJourneyImpl) AddCustomer(ctx context.Context, c model.Customer) (sql.Result, error) {
	columnNames, values := make([]string, 0), make([]string, 0)
	for i := 0; i < reflect.TypeOf(c).NumField(); i++ {
		columnNames = append(columnNames, "`"+reflect.TypeOf(c).Field(i).Tag.Get("db")+"`")
		values = append(values, ":"+reflect.TypeOf(c).Field(i).Tag.Get("db"))
	}
	return tx.TX.NamedExec(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", c.TableName(), strings.Join(columnNames, ", "), strings.Join(values, ", ")), &c)
}

func (db *JourneyImpl) AddTicket(ctx context.Context, t *model.Ticket) (sql.Result, error) {
	columnNames, values := make([]string, 0), make([]string, 0)
	for i := 0; i < reflect.TypeOf(*t).NumField(); i++ {
		columnNames = append(columnNames, "`"+reflect.TypeOf(*t).Field(i).Tag.Get("db")+"`")
		values = append(values, ":"+reflect.TypeOf(*t).Field(i).Tag.Get("db"))
	}
	return db.DB.NamedExec(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", t.TableName(), strings.Join(columnNames, ", "), strings.Join(values, ", ")), t)
}

func (tx *TxJourneyImpl) AddTicket(ctx context.Context, t *model.Ticket) (sql.Result, error) {
	columnNames, values := make([]string, 0), make([]string, 0)
	for i := 0; i < reflect.TypeOf(*t).NumField(); i++ {
		columnNames = append(columnNames, "`"+reflect.TypeOf(*t).Field(i).Tag.Get("db")+"`")
		values = append(values, ":"+reflect.TypeOf(*t).Field(i).Tag.Get("db"))
	}
	return tx.TX.NamedExec(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", t.TableName(), strings.Join(columnNames, ", "), strings.Join(values, ", ")), t)
}

func (db *JourneyImpl) UpdateCustomerNameByEmail(ctx context.Context, name, email string) (sql.Result, error) {
	return db.DB.ExecContext(ctx, "UPDATE customer SET name = ? WHERE email = ?", name, email)
}

func (tx *TxJourneyImpl) UpdateCustomerNameByEmail(ctx context.Context, name, email string) (sql.Result, error) {
	return tx.TX.ExecContext(ctx, "UPDATE customer SET name = ? WHERE email = ?", name, email)
}

func (db *JourneyImpl) DeleteCustomerByID(ctx context.Context, ID uint64) (sql.Result, error) {
	return db.DB.ExecContext(ctx, "DELETE FROM customer WHERE id = ?", ID)
}

func (tx *TxJourneyImpl) DeleteCustomerByID(ctx context.Context, ID uint64) (sql.Result, error) {
	return tx.TX.ExecContext(ctx, "DELETE FROM customer WHERE id = ?", ID)
}

func (db *JourneyImpl) GetCustomerByID(ctx context.Context, ID uint64) (model.Customer, error) {
	var result model.Customer
	err := db.DB.Get(&result, "SELECT * FROM customer WHERE id = ?", ID)
	return result, err
}

func (tx *TxJourneyImpl) GetCustomerByID(ctx context.Context, ID uint64) (model.Customer, error) {
	var result model.Customer
	err := tx.TX.Get(&result, "SELECT * FROM customer WHERE id = ?", ID)
	return result, err
}

func (db *JourneyImpl) GetTicketByID(ctx context.Context, ID uint64) (*model.Ticket, error) {
	var result model.Ticket
	err := db.DB.Get(&result, "SELECT * FROM ticket WHERE id = ?", ID)
	return &result, err
}

func (tx *TxJourneyImpl) GetTicketByID(ctx context.Context, ID uint64) (*model.Ticket, error) {
	var result model.Ticket
	err := tx.TX.Get(&result, "SELECT * FROM ticket WHERE id = ?", ID)
	return &result, err
}

func (db *JourneyImpl) GetTicketsByCustomerID(ctx context.Context, customerID uint64) ([]*model.Ticket, error) {
	var result []*model.Ticket
	err := db.DB.Select(&result, "SELECT * FROM ticket WHERE customer_id = ? ORDER BY id DESC LIMIT 10", customerID)
	return result, err
}

func (tx *TxJourneyImpl) GetTicketsByCustomerID(ctx context.Context, customerID uint64) ([]*model.Ticket, error) {
	var result []*model.Ticket
	err := tx.TX.Select(&result, "SELECT * FROM ticket WHERE customer_id = ? ORDER BY id DESC LIMIT 10", customerID)
	return result, err
}

func (db *JourneyImpl) GetCustomersByNameAndEmail(ctx context.Context, name, email string, offset, limit uint64) ([]model.Customer, error) {
	var result []model.Customer
	err := db.DB.Select(&result, "SELECT * FROM customer WHERE name = ? AND email = ? LIMIT ?, ?", name, email, offset, limit)
	return result, err
}

func (tx *TxJourneyImpl) GetCustomersByNameAndEmail(ctx context.Context, name, email string, offset, limit uint64) ([]model.Customer, error) {
	var result []model.Customer
	err := tx.TX.Select(&result, "SELECT * FROM customer WHERE name = ? AND email = ? LIMIT ?, ?", name, email, offset, limit)
	return result, err
}
