// Code generated from customer.sql by gpa. DO NOT EDIT.

package test_mysql_model

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

func (Customer) TableName() string {
	return "customer"
}

type Customer struct {
	CardId     int16          `db:"card_id" json:"card_id"` // card id
	Order      int32          `db:"order" json:"order"`     // order by
	Id         uint64         `db:"id" json:"id"`
	MemberId   sql.NullInt64  `db:"member_id" json:"member_id"` // member id
	Email      string         `db:"email" json:"email"`
	Name       string         `db:"name" json:"name"`
	Address    sql.NullString `db:"address" json:"address"`
	CreateTime time.Time      `db:"create_time" json:"create_time"`
	UpdateTime mysql.NullTime `db:"update_time" json:"update_time"`
}

func NewCustomer() *Customer {
	return &Customer{}
}
