// Code generated from journey.go by gpa. DO NOT EDIT.

package test_mysql_interface

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/liupangzi/gpa/test/mysql/another_model"
	model "github.com/liupangzi/gpa/test/mysql/model"
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
	return db.DB.NamedExecContext(ctx, fmt.Sprintf("INSERT INTO %s (`card_id`, `order`, `id`, `member_id`, `email`, `name`, `address`, `create_time`, `update_time`) VALUES (:card_id, :order, :id, :member_id, :email, :name, :address, :create_time, :update_time)", c.TableName()), &c)
}

func (tx *TxJourneyImpl) AddCustomer(ctx context.Context, c model.Customer) (sql.Result, error) {
	return tx.TX.NamedExecContext(ctx, fmt.Sprintf("INSERT INTO %s (`card_id`, `order`, `id`, `member_id`, `email`, `name`, `address`, `create_time`, `update_time`) VALUES (:card_id, :order, :id, :member_id, :email, :name, :address, :create_time, :update_time)", c.TableName()), &c)
}

func (db *JourneyImpl) AddTicket(ctx context.Context, t *another_model.Ticket) (sql.Result, error) {
	return db.DB.NamedExecContext(ctx, fmt.Sprintf("INSERT INTO %s (`from`, `duration`, `customer_id`, `id`, `price`, `meal`, `to`, `operation`, `secret_key`, `climate`, `tea`, `token`, `create_time`, `update_time`) VALUES (:from, :duration, :customer_id, :id, :price, :meal, :to, :operation, :secret_key, :climate, :tea, :token, :create_time, :update_time)", t.TableName()), t)
}

func (tx *TxJourneyImpl) AddTicket(ctx context.Context, t *another_model.Ticket) (sql.Result, error) {
	return tx.TX.NamedExecContext(ctx, fmt.Sprintf("INSERT INTO %s (`from`, `duration`, `customer_id`, `id`, `price`, `meal`, `to`, `operation`, `secret_key`, `climate`, `tea`, `token`, `create_time`, `update_time`) VALUES (:from, :duration, :customer_id, :id, :price, :meal, :to, :operation, :secret_key, :climate, :tea, :token, :create_time, :update_time)", t.TableName()), t)
}

func (db *JourneyImpl) UpdateCustomerNameByEmail(ctx context.Context, name, email string) (sql.Result, error) {
	return db.DB.ExecContext(ctx, "UPDATE customer SET name = ? WHERE email = ?", name, email)
}

func (tx *TxJourneyImpl) UpdateCustomerNameByEmail(ctx context.Context, name, email string) (sql.Result, error) {
	return tx.TX.ExecContext(ctx, "UPDATE customer SET name = ? WHERE email = ?", name, email)
}

func (db *JourneyImpl) UpdateCustomerNameByIDList(ctx context.Context, name string, idList []uint64) (sql.Result, error) {
	query, args, inErr := sqlx.In("UPDATE customer SET name = ? WHERE id in (?)", name, idList)
	if inErr != nil {
		return nil, inErr
	}

	return db.DB.ExecContext(ctx, db.DB.Rebind(query), args...)
}

func (tx *TxJourneyImpl) UpdateCustomerNameByIDList(ctx context.Context, name string, idList []uint64) (sql.Result, error) {
	query, args, inErr := sqlx.In("UPDATE customer SET name = ? WHERE id in (?)", name, idList)
	if inErr != nil {
		return nil, inErr
	}

	return tx.TX.ExecContext(ctx, tx.TX.Rebind(query), args...)
}

func (db *JourneyImpl) DeleteCustomerByID(ctx context.Context, ID uint64) (sql.Result, error) {
	return db.DB.ExecContext(ctx, "DELETE FROM customer WHERE id = ?", ID)
}

func (tx *TxJourneyImpl) DeleteCustomerByID(ctx context.Context, ID uint64) (sql.Result, error) {
	return tx.TX.ExecContext(ctx, "DELETE FROM customer WHERE id = ?", ID)
}

func (db *JourneyImpl) DeleteCustomerByIDList(ctx context.Context, IDList []uint64) (sql.Result, error) {
	query, args, inErr := sqlx.In("DELETE FROM customer WHERE id In (?)", IDList)
	if inErr != nil {
		return nil, inErr
	}

	return db.DB.ExecContext(ctx, db.DB.Rebind(query), args...)
}

func (tx *TxJourneyImpl) DeleteCustomerByIDList(ctx context.Context, IDList []uint64) (sql.Result, error) {
	query, args, inErr := sqlx.In("DELETE FROM customer WHERE id In (?)", IDList)
	if inErr != nil {
		return nil, inErr
	}

	return tx.TX.ExecContext(ctx, tx.TX.Rebind(query), args...)
}

func (db *JourneyImpl) GetCustomerByID(ctx context.Context, ID uint64) (model.Customer, error) {
	var result model.Customer
	err := db.DB.GetContext(ctx, &result, "SELECT * FROM customer WHERE id = ?", ID)
	return result, err
}

func (tx *TxJourneyImpl) GetCustomerByID(ctx context.Context, ID uint64) (model.Customer, error) {
	var result model.Customer
	err := tx.TX.GetContext(ctx, &result, "SELECT * FROM customer WHERE id = ?", ID)
	return result, err
}

func (db *JourneyImpl) GetTicketByID(ctx context.Context, ID uint64) (*another_model.Ticket, error) {
	var result another_model.Ticket
	err := db.DB.GetContext(ctx, &result, "SELECT * FROM ticket WHERE id = ?", ID)
	return &result, err
}

func (tx *TxJourneyImpl) GetTicketByID(ctx context.Context, ID uint64) (*another_model.Ticket, error) {
	var result another_model.Ticket
	err := tx.TX.GetContext(ctx, &result, "SELECT * FROM ticket WHERE id = ?", ID)
	return &result, err
}

func (db *JourneyImpl) GetTicketsByCustomerID(ctx context.Context, customerID uint64) ([]*another_model.Ticket, error) {
	result := make([]*another_model.Ticket, 0, 10)
	err := db.DB.SelectContext(ctx, &result, "SELECT * FROM ticket WHERE customer_id = ? ORDER BY id DESC LIMIT 10", customerID)
	return result, err
}

func (tx *TxJourneyImpl) GetTicketsByCustomerID(ctx context.Context, customerID uint64) ([]*another_model.Ticket, error) {
	result := make([]*another_model.Ticket, 0, 10)
	err := tx.TX.SelectContext(ctx, &result, "SELECT * FROM ticket WHERE customer_id = ? ORDER BY id DESC LIMIT 10", customerID)
	return result, err
}

func (db *JourneyImpl) GetCustomersByNameAndEmail(ctx context.Context, name, email string, offset, limit uint64) ([]model.Customer, error) {
	result := make([]model.Customer, 0, limit)
	err := db.DB.SelectContext(ctx, &result, "SELECT * FROM customer WHERE name = ? AND email = ? LIMIT ?, ?", name, email, offset, limit)
	return result, err
}

func (tx *TxJourneyImpl) GetCustomersByNameAndEmail(ctx context.Context, name, email string, offset, limit uint64) ([]model.Customer, error) {
	result := make([]model.Customer, 0, limit)
	err := tx.TX.SelectContext(ctx, &result, "SELECT * FROM customer WHERE name = ? AND email = ? LIMIT ?, ?", name, email, offset, limit)
	return result, err
}

func (db *JourneyImpl) GetCustomerEmailsByName(ctx context.Context, name string) ([]string, error) {
	result := make([]string, 0, 3)
	err := db.DB.SelectContext(ctx, &result, "SELECT `email` FROM customer WHERE name = ? LIMIT 0, 3", name)
	return result, err
}

func (tx *TxJourneyImpl) GetCustomerEmailsByName(ctx context.Context, name string) ([]string, error) {
	result := make([]string, 0, 3)
	err := tx.TX.SelectContext(ctx, &result, "SELECT `email` FROM customer WHERE name = ? LIMIT 0, 3", name)
	return result, err
}

func (db *JourneyImpl) GetCustomerNameByEmail(ctx context.Context, email string) (string, error) {
	var result string
	err := db.DB.GetContext(ctx, &result, "SELECT name FROM `customer` WHERE email = ? LIMIT 1", email)
	return result, err
}

func (tx *TxJourneyImpl) GetCustomerNameByEmail(ctx context.Context, email string) (string, error) {
	var result string
	err := tx.TX.GetContext(ctx, &result, "SELECT name FROM `customer` WHERE email = ? LIMIT 1", email)
	return result, err
}

func (db *JourneyImpl) GetFirstFreeTicket(ctx context.Context) (*another_model.Ticket, error) {
	var result another_model.Ticket
	err := db.DB.GetContext(ctx, &result, "SELECT * FROM ticket WHERE id IN (1, 2) AND price = 0 limit 1")
	return &result, err
}

func (tx *TxJourneyImpl) GetFirstFreeTicket(ctx context.Context) (*another_model.Ticket, error) {
	var result another_model.Ticket
	err := tx.TX.GetContext(ctx, &result, "SELECT * FROM ticket WHERE id IN (1, 2) AND price = 0 limit 1")
	return &result, err
}

func (db *JourneyImpl) GetTicketsByIDListAndPrice(ctx context.Context, IDList []uint64, price float64) ([]*another_model.Ticket, error) {
	result := make([]*another_model.Ticket, 0, 12)

	query, args, inErr := sqlx.In("SELECT * FROM ticket WHERE id IN (?) AND `price` > ? LIMIT 12", IDList, price)
	if inErr != nil {
		return result, inErr
	}

	err := db.DB.SelectContext(ctx, &result, db.DB.Rebind(query), args...)
	return result, err
}

func (tx *TxJourneyImpl) GetTicketsByIDListAndPrice(ctx context.Context, IDList []uint64, price float64) ([]*another_model.Ticket, error) {
	result := make([]*another_model.Ticket, 0, 12)

	query, args, inErr := sqlx.In("SELECT * FROM ticket WHERE id IN (?) AND `price` > ? LIMIT 12", IDList, price)
	if inErr != nil {
		return result, inErr
	}

	err := tx.TX.SelectContext(ctx, &result, tx.TX.Rebind(query), args...)
	return result, err
}

func (db *JourneyImpl) GetTicketsLikeTea(ctx context.Context, payload map[string]interface{}) ([]another_model.Ticket, error) {
	var result []another_model.Ticket

	query, args, namedErr := sqlx.Named("SELECT * FROM ticket WHERE tea like :tea AND id IN (:idList) limit :limit", payload)
	if namedErr != nil {
		return result, namedErr
	}
	query, args, inErr := sqlx.In(query, args...)
	if inErr != nil {
		return result, inErr
	}

	err := db.DB.SelectContext(ctx, &result, db.DB.Rebind(query), args...)
	return result, err
}

func (tx *TxJourneyImpl) GetTicketsLikeTea(ctx context.Context, payload map[string]interface{}) ([]another_model.Ticket, error) {
	var result []another_model.Ticket

	query, args, namedErr := sqlx.Named("SELECT * FROM ticket WHERE tea like :tea AND id IN (:idList) limit :limit", payload)
	if namedErr != nil {
		return result, namedErr
	}
	query, args, inErr := sqlx.In(query, args...)
	if inErr != nil {
		return result, inErr
	}

	err := tx.TX.SelectContext(ctx, &result, tx.TX.Rebind(query), args...)
	return result, err
}

func (db *JourneyImpl) GetTicketsByOffset(ctx context.Context, payload *another_model.Ticket) ([]*another_model.Ticket, error) {
	result := make([]*another_model.Ticket, 0, 2)

	query, args, namedErr := sqlx.Named("SELECT * FROM ticket WHERE id in (1, 2, 3, 4, 5) limit :customer_id, 2", payload)
	if namedErr != nil {
		return result, namedErr
	}

	err := db.DB.SelectContext(ctx, &result, db.DB.Rebind(query), args...)
	return result, err
}

func (tx *TxJourneyImpl) GetTicketsByOffset(ctx context.Context, payload *another_model.Ticket) ([]*another_model.Ticket, error) {
	result := make([]*another_model.Ticket, 0, 2)

	query, args, namedErr := sqlx.Named("SELECT * FROM ticket WHERE id in (1, 2, 3, 4, 5) limit :customer_id, 2", payload)
	if namedErr != nil {
		return result, namedErr
	}

	err := tx.TX.SelectContext(ctx, &result, tx.TX.Rebind(query), args...)
	return result, err
}

func (db *JourneyImpl) GetTicketsByFrom(ctx context.Context, payload another_model.Ticket) ([]*another_model.Ticket, error) {
	result := make([]*another_model.Ticket, 0, payload.CustomerId)

	query, args, namedErr := sqlx.Named("SELECT * FROM ticket WHERE `from`=:from limit 2, :customer_id", payload)
	if namedErr != nil {
		return result, namedErr
	}

	err := db.DB.SelectContext(ctx, &result, db.DB.Rebind(query), args...)
	return result, err
}

func (tx *TxJourneyImpl) GetTicketsByFrom(ctx context.Context, payload another_model.Ticket) ([]*another_model.Ticket, error) {
	result := make([]*another_model.Ticket, 0, payload.CustomerId)

	query, args, namedErr := sqlx.Named("SELECT * FROM ticket WHERE `from`=:from limit 2, :customer_id", payload)
	if namedErr != nil {
		return result, namedErr
	}

	err := tx.TX.SelectContext(ctx, &result, tx.TX.Rebind(query), args...)
	return result, err
}

func (db *JourneyImpl) GetTicketWhichTokenIsNotEqualTo(ctx context.Context, payload map[string]interface{}) (*another_model.Ticket, error) {
	var result another_model.Ticket
	query, args, namedErr := sqlx.Named("SELECT * FROM ticket WHERE token!= :order AND id IN (:from) AND price > 0 limit 2, :limit", payload)
	if namedErr != nil {
		return &result, namedErr
	}
	query, args, inErr := sqlx.In(query, args...)
	if inErr != nil {
		return &result, inErr
	}

	err := db.DB.GetContext(ctx, &result, db.DB.Rebind(query), args...)
	return &result, err
}

func (tx *TxJourneyImpl) GetTicketWhichTokenIsNotEqualTo(ctx context.Context, payload map[string]interface{}) (*another_model.Ticket, error) {
	var result another_model.Ticket
	query, args, namedErr := sqlx.Named("SELECT * FROM ticket WHERE token!= :order AND id IN (:from) AND price > 0 limit 2, :limit", payload)
	if namedErr != nil {
		return &result, namedErr
	}
	query, args, inErr := sqlx.In(query, args...)
	if inErr != nil {
		return &result, inErr
	}

	err := tx.TX.GetContext(ctx, &result, tx.TX.Rebind(query), args...)
	return &result, err
}

func (db *JourneyImpl) GetTicketGreaterThanPrice(ctx context.Context, payload *another_model.Ticket) (another_model.Ticket, error) {
	var result another_model.Ticket
	query, args, namedErr := sqlx.Named("SELECT * FROM ticket WHERE price>:price AND id in (1, 2)", payload)
	if namedErr != nil {
		return result, namedErr
	}

	err := db.DB.GetContext(ctx, &result, db.DB.Rebind(query), args...)
	return result, err
}

func (tx *TxJourneyImpl) GetTicketGreaterThanPrice(ctx context.Context, payload *another_model.Ticket) (another_model.Ticket, error) {
	var result another_model.Ticket
	query, args, namedErr := sqlx.Named("SELECT * FROM ticket WHERE price>:price AND id in (1, 2)", payload)
	if namedErr != nil {
		return result, namedErr
	}

	err := tx.TX.GetContext(ctx, &result, tx.TX.Rebind(query), args...)
	return result, err
}

func (db *JourneyImpl) UpdateCustomerMemberIDByID(ctx context.Context, c *model.Customer) (sql.Result, error) {
	return db.DB.NamedExecContext(ctx, "Update customer sET member_id= :member_id WHERE id = 1", c)
}

func (tx *TxJourneyImpl) UpdateCustomerMemberIDByID(ctx context.Context, c *model.Customer) (sql.Result, error) {
	return tx.TX.NamedExecContext(ctx, "Update customer sET member_id= :member_id WHERE id = 1", c)
}

func (db *JourneyImpl) UpdateCustomerMemberIDByIDList(ctx context.Context, c *model.Customer) (sql.Result, error) {
	return db.DB.NamedExecContext(ctx, "Update customer sET member_id= :member_id WHERE id in (1, 2)", c)
}

func (tx *TxJourneyImpl) UpdateCustomerMemberIDByIDList(ctx context.Context, c *model.Customer) (sql.Result, error) {
	return tx.TX.NamedExecContext(ctx, "Update customer sET member_id= :member_id WHERE id in (1, 2)", c)
}

func (db *JourneyImpl) DeleteCustomerByEmail(ctx context.Context, x map[string]interface{}) (sql.Result, error) {
	return db.DB.NamedExecContext(ctx, "Delete FroM customer where email= :email", x)
}

func (tx *TxJourneyImpl) DeleteCustomerByEmail(ctx context.Context, x map[string]interface{}) (sql.Result, error) {
	return tx.TX.NamedExecContext(ctx, "Delete FroM customer where email= :email", x)
}

func (db *JourneyImpl) DeleteCustomerByEmailOrIDList(ctx context.Context, x map[string]interface{}) (sql.Result, error) {
	query, args, namedErr := sqlx.Named("Delete FroM customer where email= :email oR id In (:id_list)", x)
	if namedErr != nil {
		return nil, namedErr
	}

	query, args, inErr := sqlx.In(query, args...)
	if inErr != nil {
		return nil, inErr
	}

	return db.DB.ExecContext(ctx, db.DB.Rebind(query), args...)
}

func (tx *TxJourneyImpl) DeleteCustomerByEmailOrIDList(ctx context.Context, x map[string]interface{}) (sql.Result, error) {
	query, args, namedErr := sqlx.Named("Delete FroM customer where email= :email oR id In (:id_list)", x)
	if namedErr != nil {
		return nil, namedErr
	}

	query, args, inErr := sqlx.In(query, args...)
	if inErr != nil {
		return nil, inErr
	}

	return tx.TX.ExecContext(ctx, tx.TX.Rebind(query), args...)
}
