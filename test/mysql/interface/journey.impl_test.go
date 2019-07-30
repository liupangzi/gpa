package test_mysql_interface

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	model "github.com/liupangzi/gpa/test/mysql/model"
	. "github.com/smartystreets/goconvey/convey"
	"syreclabs.com/go/faker"
)

var j *JourneyImpl

func init() {
	db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/gpa?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	j = NewJourney(db)

	_ = db.MustExec("TRUNCATE TABLE customer;")
	_ = db.MustExec("TRUNCATE TABLE ticket;")
}

func TestCustomerImpl_AddCustomer(t *testing.T) {
	Convey("Test AddCustomer() w/ a pre-generated primary key", t, func() {
		now := time.Now()
		customer := model.Customer{
			CardId: int16(faker.RandomInt(-32768, 32767)),
			Order:  int32(faker.RandomInt(-2147483648, 2147483647)),
			Id:     uint64(faker.RandomInt(100, 9223372036854775807)),
			Email:  faker.Internet().Email(),
			Address: sql.NullString{
				String: faker.Address().String(),
				Valid:  true,
			},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
		}

		result, err := j.AddCustomer(context.TODO(), customer)
		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 1)
		So(err, ShouldBeNil)

		lastInsertID, err := result.LastInsertId()
		So(lastInsertID, ShouldBeGreaterThanOrEqualTo, 1)
		So(err, ShouldBeNil)

		getCustomer := model.Customer{}
		err = j.DB.Get(&getCustomer, "SELECT * FROM customer WHERE id = ? LIMIT 1", customer.Id)
		So(err, ShouldBeNil)
		So(customer, ShouldResemble, customer)
	})

	Convey("Test AddCustomer() w/o a primary key", t, func() {
		lastInsertID := int64(0)
		now := time.Now()

		customer := model.Customer{
			CardId:     int16(faker.RandomInt(-32768, 32767)),
			Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
			MemberId:   sql.NullInt64{},
			Email:      faker.Internet().Email(),
			Name:       faker.Name().Name(),
			Address:    sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: mysql.NullTime{
				Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				Valid: true,
			},
		}

		result, err := j.AddCustomer(context.TODO(), customer)
		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 1)
		So(err, ShouldBeNil)

		lastInsertID, err = result.LastInsertId()
		So(lastInsertID, ShouldBeGreaterThanOrEqualTo, 1)
		So(err, ShouldBeNil)

		getCustomer := model.Customer{}
		err = j.DB.Get(&getCustomer, "SELECT * FROM customer WHERE id = ?", lastInsertID)
		So(err, ShouldBeNil)

		customer.Id = uint64(lastInsertID)
		So(getCustomer, ShouldResemble, customer)
	})
}

func TestCustomerImpl_AddTicket(t *testing.T) {
	Convey("Test AddTicket() w/ a pre-generated primary key", t, func() {
		now := time.Now()
		ticket := &model.Ticket{
			From:       uint32(faker.RandomInt(1, 32767)),
			Duration:   float64(faker.RandomInt(1, 1024000)),
			CustomerId: uint64(faker.RandomInt(1, 32767)),
			Id:         uint64(faker.RandomInt(1, 32767)),
			Price:      sql.NullFloat64{},
			Meal: sql.NullInt64{
				Int64: 12,
				Valid: true,
			},
			To:        sql.NullInt64{},
			Operation: []byte{1, 2, 3, 4, 5},
			SecretKey: []byte{6, 7, 8, 9, 10},
			Climate: sql.NullString{
				String: "r",
				Valid:  true,
			},
			Tea:        sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: mysql.NullTime{
				Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				Valid: true,
			},
		}
		result, err := j.AddTicket(context.TODO(), ticket)
		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 1)
		So(err, ShouldBeNil)

		lastInsertID, err := result.LastInsertId()
		So(lastInsertID, ShouldBeGreaterThanOrEqualTo, 1)
		So(err, ShouldBeNil)

		getTicket := &model.Ticket{}
		err = j.DB.Get(getTicket, "SELECT * FROM ticket WHERE id = ? LIMIT 1", ticket.Id)
		So(err, ShouldBeNil)
		So(ticket, ShouldResemble, ticket)
	})

	Convey("Test AddTicket() w/o a primary key", t, func() {
		lastInsertID := int64(0)
		now := time.Now().AddDate(-1, -2, -3)

		ticket := &model.Ticket{
			From:       uint32(faker.RandomInt(1, 32767)),
			Duration:   float64(faker.RandomInt(1, 1024000)),
			CustomerId: uint64(faker.RandomInt(1, 32767)),
			Price: sql.NullFloat64{
				Float64: 1.023,
				Valid:   true,
			},
			Meal: sql.NullInt64{},
			To: sql.NullInt64{
				Int64: 1024,
				Valid: true,
			},
			Operation: []byte{10, 20, 30, 40, 50},
			SecretKey: []byte{60, 70, 80, 90, 100},
			Token: sql.NullString{
				String: "1234567890abcdefs",
				Valid:  true,
			},
			Climate:    sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: mysql.NullTime{
				Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				Valid: true,
			},
		}

		result, err := j.AddTicket(context.TODO(), ticket)
		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 1)
		So(err, ShouldBeNil)

		lastInsertID, err = result.LastInsertId()
		So(lastInsertID, ShouldBeGreaterThanOrEqualTo, 1)
		So(err, ShouldBeNil)

		getTicket := &model.Ticket{}
		err = j.DB.Get(getTicket, "SELECT * FROM ticket WHERE id = ?", lastInsertID)
		So(err, ShouldBeNil)

		ticket.Id = uint64(lastInsertID)
		So(getTicket, ShouldResemble, ticket)
	})
}

func TestJourneyImpl_UpdateCustomerNameByEmail(t *testing.T) {
	Convey("Add customer", t, func() {
		now := time.Now().AddDate(-1, 2, 3)
		customer := model.Customer{
			CardId:     int16(faker.RandomInt(-32768, 32767)),
			Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
			MemberId:   sql.NullInt64{},
			Email:      "i@liuchao.me",
			Name:       "a wrong name",
			Address:    sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: mysql.NullTime{
				Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				Valid: true,
			},
		}

		result, err := j.AddCustomer(context.TODO(), customer)
		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 1)
		So(err, ShouldBeNil)
	})

	Convey(`Update the customer with name "liupangzi"`, t, func() {
		result, err := j.UpdateCustomerNameByEmail(context.TODO(), "liupangzi", "i@liuchao.me")
		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 1)
		So(err, ShouldBeNil)

		getCustomer := &model.Customer{}
		err = j.DB.Get(getCustomer, "SELECT * FROM customer WHERE email = \"i@liuchao.me\" LIMIT 1")
		So(err, ShouldBeNil)
		So(getCustomer.Name, ShouldEqual, "liupangzi")
	})
}

func TestJourneyImpl_DeleteCustomerByID(t *testing.T) {
	ID := uint64(100)

	Convey("Add customer", t, func() {
		now := time.Now().AddDate(-1, 2, 3)
		customer := model.Customer{
			CardId: int16(faker.RandomInt(-32768, 32767)),
			Order:  int32(faker.RandomInt(-2147483648, 2147483647)),
			Id:     ID,
			MemberId: sql.NullInt64{
				Int64: faker.RandomInt64(1, 3344565),
				Valid: true,
			},
			Email: faker.Internet().Email(),
			Name:  faker.Name().Name(),
			Address: sql.NullString{
				String: faker.Address().String(),
				Valid:  true,
			},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: mysql.NullTime{
				Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				Valid: true,
			},
		}

		result, err := j.AddCustomer(context.TODO(), customer)
		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 1)
		So(err, ShouldBeNil)
	})

	Convey("Delete customer w/ ID == 100", t, func() {
		result, err := j.DeleteCustomerByID(context.TODO(), ID)
		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 1)
		So(err, ShouldBeNil)

		err = j.DB.Get(&model.Customer{}, "SELECT * FROM customer WHERE id = 100 LIMIT 1")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "sql: no rows in result set")
	})
}
