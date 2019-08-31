package test_mysql_interface

import (
	"context"
	"database/sql"

	"github.com/liupangzi/gpa/test/mysql/another_model"
	model "github.com/liupangzi/gpa/test/mysql/model"
)

type Journey interface {
	// @Insert()
	AddCustomer(ctx context.Context, c model.Customer) (sql.Result, error)

	// @Insert(				)
	AddTicket(ctx context.Context, t *another_model.Ticket) (sql.Result, error)

	// @Update("UPDATE 		customer SET name = ? WHERE email = ?")
	UpdateCustomerNameByEmail(ctx context.Context, name, email string) (sql.Result, error)

	// @Update("UPDATE customer SET name = ? WHERE id in (?)")
	UpdateCustomerNameByIDList(ctx context.Context, name string, idList []uint64) (sql.Result, error)

	// @Delete("DELETE FROM  customer WHERE 	id   = ?")
	DeleteCustomerByID(ctx context.Context, ID uint64) (sql.Result, error)

	// @Delete("DELETE FROM  customer WHERE 	id   In     (?)")
	DeleteCustomerByIDList(ctx context.Context, IDList []uint64) (sql.Result, error)

	// @Select("SELECT 		* FROM customer WHERE id   = ?")
	GetCustomerByID(ctx context.Context, ID uint64) (model.Customer, error)

	/* @Select(
	"SELECT * FROM ticket WHERE id =  ?"
	)    */
	GetTicketByID(ctx context.Context, ID uint64) (*another_model.Ticket, error)

	/*
	  @Select("SELECT * FROM  ticket WHERE customer_id = ? ORDER BY id DESC LIMIT 10")
	*/
	GetTicketsByCustomerID(ctx context.Context, customerID uint64) ([]*another_model.Ticket, error)

	/* @Select("SELECT * FROM customer
	        WHERE name = ? AND
	email = ? LIMIT ?, ?") */
	GetCustomersByNameAndEmail(ctx context.Context, name, email string, offset, limit uint64) ([]model.Customer, error)

	// @Select("SELECT `email` 		FROM customer 		WHERE name = ? LIMIT 0, 3")
	GetCustomerEmailsByName(ctx context.Context, name string) ([]string, error)

	// @Select("SELECT name FROM `customer` WHERE email = 	? LIMIT 1")
	GetCustomerNameByEmail(ctx context.Context, email string) (string, error)

	// @Select("SELECT *	 FROM ticket WHERE id IN (1,	 2) AND price = 0 limit 1")
	GetFirstFreeTicket(ctx context.Context) (*another_model.Ticket, error)

	// @Select("SELECT * FROM ticket WHERE id IN (?) AND `price` > ? LIMIT 12")
	GetTicketsByIDListAndPrice(ctx context.Context, IDList []uint64, price float64) ([]*another_model.Ticket, error)

	// @Select("SELECT * FROM ticket WHERE tea like :tea AND id IN (:idList) limit :limit")
	GetTicketsLikeTea(ctx context.Context, payload map[string]interface{}) ([]another_model.Ticket, error)

	// @Select("SELECT * FROM ticket WHERE id in (1, 2, 3, 4, 5) limit :customer_id, 2")
	GetTicketsByOffset(ctx context.Context, payload *another_model.Ticket) ([]*another_model.Ticket, error)

	// @Select("SELECT * FROM ticket WHERE `from`=:from limit 2, :customer_id")
	GetTicketsByFrom(ctx context.Context, payload another_model.Ticket) ([]*another_model.Ticket, error)

	// @Select("SELECT * FROM ticket WHERE token!= :order AND id IN (:from) AND price > 0 limit 2, :limit")
	GetTicketWhichTokenIsNotEqualTo(ctx context.Context, payload map[string]interface{}) (*another_model.Ticket, error)

	// @Select("SELECT * FROM ticket WHERE price>:price AND id in (1, 2)")
	GetTicketGreaterThanPrice(ctx context.Context, payload *another_model.Ticket) (another_model.Ticket, error)

	// @Update("Update customer sET member_id= :member_id WHERE id = 1")
	UpdateCustomerMemberIDByID(ctx context.Context, c *model.Customer) (sql.Result, error)

	// @Update("Update customer sET member_id= :member_id WHERE id in (1, 2)")
	UpdateCustomerMemberIDByIDList(ctx context.Context, c *model.Customer) (sql.Result, error)

	// @Delete("Delete FroM customer where email= :email")
	DeleteCustomerByEmail(ctx context.Context, x map[string]interface{}) (sql.Result, error)

	// @Delete("Delete FroM customer where email= :email oR id In (:id_list)")
	DeleteCustomerByEmailOrIDList(ctx context.Context, x map[string]interface{}) (sql.Result, error)
}
