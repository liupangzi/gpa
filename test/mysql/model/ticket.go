// Code generated from ticket.sql by gpa. DO NOT EDIT.

package test_mysql_model

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

func (Ticket) TableName() string {
	return "ticket"
}

type Ticket struct {
	From       uint32          `db:"from" json:"from"`
	Duration   float64         `db:"duration" json:"duration"`
	CustomerId uint64          `db:"customer_id" json:"customer_id"` // customer id
	Id         uint64          `db:"id" json:"id"`
	Price      sql.NullFloat64 `db:"price" json:"price"` // 1 dollar
	Meal       sql.NullInt64   `db:"meal" json:"meal"`
	To         sql.NullInt64   `db:"to" json:"to"`
	Operation  []byte          `db:"operation" json:"operation"`
	SecretKey  []byte          `db:"secret_key" json:"secret_key"` // Whats your secret
	Climate    sql.NullString  `db:"climate" json:"climate"`
	Tea        sql.NullString  `db:"tea" json:"tea"`
	Token      sql.NullString  `db:"token" json:"token"` // A long long text
	CreateTime time.Time       `db:"create_time" json:"create_time"`
	UpdateTime mysql.NullTime  `db:"update_time" json:"update_time"`
}

func NewTicket() *Ticket {
	return &Ticket{}
}
