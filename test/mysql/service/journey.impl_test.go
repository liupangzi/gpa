package test_mysql_interface

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/liupangzi/gpa/test/mysql/another_model"
	model "github.com/liupangzi/gpa/test/mysql/model"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	"syreclabs.com/go/faker"
)

var j *JourneyImpl

type JourneySuite struct {
	suite.Suite
}

func TestJourneySuite(t *testing.T) {
	suite.Run(t, new(JourneySuite))
}

func (s *JourneySuite) SetupTest() {
	db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/gpa?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	j = NewJourney(db)
}

func (s *JourneySuite) BeforeTest(suiteName, testName string) {
	_ = j.DB.MustExec("TRUNCATE TABLE customer;")
	_ = j.DB.MustExec("TRUNCATE TABLE ticket;")
}

func (s *JourneySuite) TestAddCustomer() {
	Convey("Test AddCustomer() w/ a pre-generated primary key", s.T(), func() {
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

	Convey("Test AddCustomer() w/o a primary key", s.T(), func() {
		lastInsertID := int64(0)
		now := time.Now()

		customer := model.Customer{
			CardId:     int16(faker.RandomInt(-32768, 32767)),
			Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
			MemberId:   sql.NullInt32{},
			Email:      faker.Internet().Email(),
			Name:       faker.Name().Name(),
			Address:    sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: sql.NullTime{
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

func (s *JourneySuite) TestAddTicket() {
	Convey("Test AddTicket() w/ a pre-generated primary key", s.T(), func() {
		now := time.Now()
		ticket := &another_model.Ticket{
			From:       uint32(faker.RandomInt(1, 32767)),
			Duration:   float64(faker.RandomInt(1, 1024000)),
			CustomerId: uint64(faker.RandomInt(1, 32767)),
			Id:         uint64(faker.RandomInt(1, 32767)),
			Price:      sql.NullFloat64{},
			Meal: sql.NullInt64{
				Int64: 12,
				Valid: true,
			},
			To:        sql.NullInt32{},
			Operation: []byte{1, 2, 3, 4, 5},
			SecretKey: []byte{6, 7, 8, 9, 10},
			Climate: sql.NullString{
				String: "r",
				Valid:  true,
			},
			Tea:        sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: sql.NullTime{
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

		getTicket := &another_model.Ticket{}
		err = j.DB.Get(getTicket, "SELECT * FROM ticket WHERE id = ? LIMIT 1", ticket.Id)
		So(err, ShouldBeNil)
		So(ticket, ShouldResemble, ticket)
	})

	Convey("Test AddTicket() w/o a primary key", s.T(), func() {
		lastInsertID := int64(0)
		now := time.Now().AddDate(-1, -2, -3)

		ticket := &another_model.Ticket{
			From:       uint32(faker.RandomInt(1, 32767)),
			Duration:   float64(faker.RandomInt(1, 1024000)),
			CustomerId: uint64(faker.RandomInt(1, 32767)),
			Price: sql.NullFloat64{
				Float64: 1.023,
				Valid:   true,
			},
			Meal: sql.NullInt64{},
			To: sql.NullInt32{
				Int32: 1024,
				Valid: true,
			},
			Operation: []byte{10, 20, 30, 40, 50},
			SecretKey: []byte{60, 70, 80, 90, 100},
			Token: sql.NullString{
				String: "1234567890abc",
				Valid:  true,
			},
			Climate:    sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: sql.NullTime{
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

		getTicket := &another_model.Ticket{}
		err = j.DB.Get(getTicket, "SELECT * FROM ticket WHERE id = ?", lastInsertID)
		So(err, ShouldBeNil)

		ticket.Id = uint64(lastInsertID)
		So(getTicket, ShouldResemble, ticket)
	})
}

func (s *JourneySuite) TestUpdateCustomerNameByEmail() {
	email := faker.Internet().Email()
	name := faker.Name().String()

	Convey("Add customer", s.T(), func() {
		now := time.Now().AddDate(-1, 2, 3)
		customer := model.Customer{
			CardId:     int16(faker.RandomInt(-32768, 32767)),
			Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
			MemberId:   sql.NullInt32{},
			Email:      email,
			Name:       faker.Name().String(),
			Address:    sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: sql.NullTime{
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

	Convey(`Update the customer's name by email`, s.T(), func() {
		result, err := j.UpdateCustomerNameByEmail(context.TODO(), name, email)
		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 1)
		So(err, ShouldBeNil)

		getCustomer := &model.Customer{}
		err = j.DB.Get(getCustomer, "SELECT * FROM customer WHERE email = ? LIMIT 1", email)
		So(err, ShouldBeNil)
		So(getCustomer.Name, ShouldEqual, name)
	})
}

func (s *JourneySuite) TestUpdateCustomerNameByIDList() {
	Convey("Add customers", s.T(), func() {
		Convey("Add first customer", func() {
			now := time.Now().AddDate(-1, 2, 3)
			customer := model.Customer{
				Id:         1,
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId:   sql.NullInt32{},
				Email:      "i@liuchao.me",
				Name:       "foo",
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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

		Convey("Add second customer", func() {
			now := time.Now().AddDate(-4, 5, 6)
			customer := model.Customer{
				Id:         2,
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId:   sql.NullInt32{},
				Email:      "me@example.com",
				Name:       "bar",
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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
	})

	Convey(`Update the customer's name by id list`, s.T(), func() {
		result, err := j.UpdateCustomerNameByIDList(context.TODO(), "liupangzi", []uint64{1, 2})
		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 2)
		So(err, ShouldBeNil)

		var getCustomers []*model.Customer
		err = j.DB.Select(&getCustomers, "SELECT * FROM customer WHERE id in (1, 2)")
		So(err, ShouldBeNil)
		So(getCustomers[0].Name, ShouldEqual, "liupangzi")
		So(getCustomers[0].Email, ShouldEqual, "i@liuchao.me")
		So(getCustomers[1].Name, ShouldEqual, "liupangzi")
		So(getCustomers[1].Email, ShouldEqual, "me@example.com")
	})
}

func (s *JourneySuite) TestDeleteCustomerByID() {
	ID := uint64(100)

	Convey("Add customer w/ id == 100", s.T(), func() {
		now := time.Now().AddDate(-1, 2, 3)
		customer := model.Customer{
			CardId: int16(faker.RandomInt(-32768, 32767)),
			Order:  int32(faker.RandomInt(-2147483648, 2147483647)),
			Id:     ID,
			MemberId: sql.NullInt32{
				Int32: int32(faker.RandomInt(1, 3344565)),
				Valid: true,
			},
			Email: faker.Internet().Email(),
			Name:  faker.Name().Name(),
			Address: sql.NullString{
				String: faker.Address().String(),
				Valid:  true,
			},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: sql.NullTime{
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

	Convey("Delete customer w/ id == 100", s.T(), func() {
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

func (s *JourneySuite) TestDeleteCustomerByIDList() {
	Convey("Add customers", s.T(), func() {
		Convey("Add first customer", func() {
			now := time.Now().AddDate(-1, 2, 3)
			customer := model.Customer{
				CardId: int16(faker.RandomInt(-32768, 32767)),
				Order:  int32(faker.RandomInt(-2147483648, 2147483647)),
				Id:     100,
				MemberId: sql.NullInt32{
					Int32: int32(faker.RandomInt(1, 3344565)),
					Valid: true,
				},
				Email: faker.Internet().Email(),
				Name:  faker.Name().Name(),
				Address: sql.NullString{
					String: faker.Address().String(),
					Valid:  true,
				},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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

		Convey("Add second customer", func() {
			now := time.Now().AddDate(-9, -6, 3)
			customer := model.Customer{
				CardId: int16(faker.RandomInt(-32768, 32767)),
				Order:  int32(faker.RandomInt(-2147483648, 2147483647)),
				Id:     101,
				MemberId: sql.NullInt32{
					Int32: int32(faker.RandomInt(1, 3344565)),
					Valid: true,
				},
				Email: faker.Internet().Email(),
				Name:  faker.Name().Name(),
				Address: sql.NullString{
					String: faker.Address().String(),
					Valid:  true,
				},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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

		Convey("Add third customer", func() {
			now := time.Now().AddDate(-2, 8, -3)
			customer := model.Customer{
				CardId: int16(faker.RandomInt(-32768, 32767)),
				Order:  int32(faker.RandomInt(-2147483648, 2147483647)),
				Id:     1,
				MemberId: sql.NullInt32{
					Int32: int32(faker.RandomInt(1, 3344565)),
					Valid: true,
				},
				Email: faker.Internet().Email(),
				Name:  "liupangzi",
				Address: sql.NullString{
					String: faker.Address().String(),
					Valid:  true,
				},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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
	})

	Convey("Delete customer id in (100, 101)", s.T(), func() {
		result, err := j.DeleteCustomerByIDList(context.TODO(), []uint64{100, 101})
		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 2)
		So(err, ShouldBeNil)

		customer := model.Customer{}
		err = j.DB.Get(&customer, "SELECT * FROM customer LIMIT 1")
		So(err, ShouldBeNil)
		So(customer.Id, ShouldEqual, 1)
		So(customer.Name, ShouldEqual, "liupangzi")
	})
}

func (s *JourneySuite) TestGetCustomerByID() {
	Convey("Test GetCustomerByID", s.T(), func() {
		lastInsertID := int64(0)
		now := time.Now()

		customer := model.Customer{
			CardId:     int16(faker.RandomInt(-32768, 32767)),
			Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
			MemberId:   sql.NullInt32{},
			Email:      faker.Internet().Email(),
			Name:       faker.Name().Name(),
			Address:    sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: sql.NullTime{
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

		getCustomer, err := j.GetCustomerByID(context.TODO(), uint64(lastInsertID))
		So(err, ShouldBeNil)

		customer.Id = uint64(lastInsertID)
		So(getCustomer, ShouldResemble, customer)
	})
}

func (s *JourneySuite) TestGetTicketByID() {
	Convey("Test GetTicketByID()", s.T(), func() {
		lastInsertID := int64(0)
		now := time.Now().AddDate(-4, -5, -6)

		ticket := &another_model.Ticket{
			From:       uint32(faker.RandomInt(1, 32767)),
			Duration:   float64(faker.RandomInt(1, 1024000)),
			CustomerId: uint64(faker.RandomInt(1, 32767)),
			Price: sql.NullFloat64{
				Float64: 1.023,
				Valid:   true,
			},
			Meal: sql.NullInt64{},
			To: sql.NullInt32{
				Int32: 1024,
				Valid: true,
			},
			Operation: []byte{100, 200, 102, 140, 150},
			SecretKey: []byte{26, 67, 48, 23, 120},
			Token: sql.NullString{
				String: "1234567890xyz",
				Valid:  true,
			},
			Climate:    sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: sql.NullTime{
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

		getTicket, err := j.GetTicketByID(context.TODO(), uint64(lastInsertID))
		So(err, ShouldBeNil)

		ticket.Id = uint64(lastInsertID)
		So(getTicket, ShouldResemble, ticket)
	})
}

func (s *JourneySuite) TestGetTicketsByCustomerID() {
	const customerID uint64 = 1234321

	Convey("Add tickets for customer", s.T(), func() {
		Convey("First ticket", func() {
			lastInsertID := int64(0)
			now := time.Now().AddDate(-4, -5, -6)

			ticket := &another_model.Ticket{
				From:       1024,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: customerID,
				Price: sql.NullFloat64{
					Float64: 1.023,
					Valid:   true,
				},
				Meal: sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "1234567890xyz",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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
		})

		Convey("Second ticket", func() {
			lastInsertID := int64(0)
			now := time.Now().AddDate(-4, -5, -6)

			ticket := &another_model.Ticket{
				From:       1025,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: customerID,
				Price: sql.NullFloat64{
					Float64: 1.023,
					Valid:   true,
				},
				Meal: sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "1234567890xyz",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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
		})
	})

	Convey("Test GetTicketsByCustomerID()", s.T(), func() {
		tickets, err := j.GetTicketsByCustomerID(context.Background(), customerID)
		So(err, ShouldBeNil)
		So(tickets, ShouldHaveLength, 2)
		So(tickets[0].From, ShouldEqual, 1025)
		So(tickets[1].From, ShouldEqual, 1024)
	})
}

func (s *JourneySuite) TestGetCustomersByNameAndEmail() {
	name := faker.Name().String()
	email := faker.Internet().Email()

	Convey("Add customers", s.T(), func() {
		Convey("First customer", func() {
			lastInsertID := int64(0)
			now := time.Now()

			customer := model.Customer{
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      -32,
				MemberId:   sql.NullInt32{},
				Email:      email,
				Name:       name,
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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
		})

		Convey("Second customer", func() {
			lastInsertID := int64(0)
			now := time.Now()

			customer := model.Customer{
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      -16,
				MemberId:   sql.NullInt32{},
				Email:      faker.Internet().Email(),
				Name:       faker.Name().String(),
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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
		})

		Convey("Third customer", func() {
			lastInsertID := int64(0)
			now := time.Now()

			customer := model.Customer{
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      -8,
				MemberId:   sql.NullInt32{},
				Email:      email,
				Name:       name,
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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
		})
	})

	Convey("Test GetCustomersByNameAndEmail()", s.T(), func() {
		result, err := j.GetCustomersByNameAndEmail(context.Background(), name, email, 0, 10)
		So(err, ShouldBeNil)
		So(result, ShouldHaveLength, 2)
		So(result[0].Order, ShouldEqual, -32)
		So(result[1].Order, ShouldEqual, -8)
	})
}

func (s *JourneySuite) TestGetCustomerEmailsByName() {
	name := faker.Name().String()
	email1, email2 := faker.Internet().Email(), faker.Internet().Email()

	Convey("Add customers", s.T(), func() {
		Convey("First customer", func() {
			lastInsertID := int64(0)
			now := time.Now()

			customer := model.Customer{
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId:   sql.NullInt32{},
				Email:      email1,
				Name:       name,
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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
		})

		Convey("Second customer", func() {
			lastInsertID := int64(0)
			now := time.Now()

			customer := model.Customer{
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId:   sql.NullInt32{},
				Email:      faker.Internet().Email(),
				Name:       faker.Name().String(),
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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
		})

		Convey("Third customer", func() {
			lastInsertID := int64(0)
			now := time.Now()

			customer := model.Customer{
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId:   sql.NullInt32{},
				Email:      email2,
				Name:       name,
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
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
		})
	})

	Convey("Test GetCustomerEmailsByName()", s.T(), func() {
		result, err := j.GetCustomerEmailsByName(context.Background(), name)
		So(err, ShouldBeNil)
		So(result, ShouldHaveLength, 2)
		So(result[0], ShouldEqual, email1)
		So(result[1], ShouldEqual, email2)
	})
}

func (s *JourneySuite) TestGetCustomerNameByEmail() {
	name := faker.Name().String()
	email := faker.Internet().Email()

	Convey("Add customer", s.T(), func() {
		lastInsertID := int64(0)
		now := time.Now()

		customer := model.Customer{
			CardId:     int16(faker.RandomInt(-32768, 32767)),
			Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
			MemberId:   sql.NullInt32{},
			Email:      email,
			Name:       name,
			Address:    sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: sql.NullTime{
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
	})

	Convey("Test GetCustomerNameByEmail()", s.T(), func() {
		result, err := j.GetCustomerNameByEmail(context.Background(), email)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, name)
	})
}

func (s *JourneySuite) TestGetFirstFreeTicket() {
	Convey("Add tickets for customer", s.T(), func() {
		Convey("First ticket", func() {
			now := time.Now().AddDate(-4, -5, -6)
			ticket := &another_model.Ticket{
				Id:         1,
				From:       1024,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Price: sql.NullFloat64{
					Float64: 0.0,
					Valid:   true,
				},
				Meal: sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "1234567890free",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Not a free ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         1024,
				From:       1026,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Price: sql.NullFloat64{
					Float64: 0.01,
					Valid:   true,
				},
				Meal: sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "1234567890lol",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Second ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         2,
				From:       1025,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "1234567890free",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
	})

	Convey("Test GetFirstFreeTicket()", s.T(), func() {
		result, err := j.GetFirstFreeTicket(context.Background())
		So(err, ShouldBeNil)
		So(result.Id, ShouldEqual, 1)
		So(result.Price, ShouldResemble, sql.NullFloat64{
			Float64: 0.0,
			Valid:   true,
		})
	})
}

func (s *JourneySuite) TestGetTicketsByIDListAndPrice() {
	Convey("Add tickets for customer", s.T(), func() {
		Convey("First ticket", func() {
			now := time.Now().AddDate(-4, -5, -6)
			ticket := &another_model.Ticket{
				Id:         1,
				From:       1024,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Price: sql.NullFloat64{
					Float64: 1.0,
					Valid:   true,
				},
				Meal: sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "1234567890free",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Second ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         2,
				From:       1025,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Price: sql.NullFloat64{
					Float64: 0.01,
					Valid:   true,
				},
				Meal: sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "1234567890lol",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Third ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         3,
				From:       1099,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "1234567890free",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
	})

	Convey("Test GetTicketsByIDListAndPrice()", s.T(), func() {
		result, err := j.GetTicketsByIDListAndPrice(context.Background(), []uint64{1, 2, 3}, 0.009)
		So(err, ShouldBeNil)
		So(result, ShouldHaveLength, 2)

		So(result[0].Id, ShouldEqual, 1)
		So(result[0].Price.Float64, ShouldAlmostEqual, 1.0)

		So(result[1].Id, ShouldEqual, 2)
		So(result[1].Price.Float64, ShouldAlmostEqual, 0.01)
	})
}

func (s *JourneySuite) TestGetTicketsLikeTea() {
	Convey("Add tickets for customer", s.T(), func() {
		Convey("First ticket", func() {
			now := time.Now().AddDate(-4, -5, -6)
			ticket := &another_model.Ticket{
				Id:         1,
				From:       1024,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "Black",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Second ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         2,
				From:       1025,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "Green",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Third ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         3,
				From:       1068,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "Oolong",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Fourth ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         4,
				From:       1099,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "White",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
	})

	Convey("Test GetTicketsLikeTea()", s.T(), func() {
		Convey("limit 1", func() {
			payload := make(map[string]interface{})
			payload["tea"] = "%e%"
			payload["idList"] = []int64{1, 2, 3, 4, 5, 6, 7}
			payload["limit"] = int64(1)

			result, err := j.GetTicketsLikeTea(context.Background(), payload)
			So(err, ShouldBeNil)
			So(result, ShouldHaveLength, 1)

			So(result[0].Id, ShouldEqual, 2)
			So(result[0].Tea.String, ShouldEqual, "Green")
		})

		Convey("limit 2", func() {
			payload := make(map[string]interface{})
			payload["tea"] = "%e%"
			payload["idList"] = []int64{1, 2, 3, 4, 5, 6, 7}
			payload["limit"] = int64(2)

			result, err := j.GetTicketsLikeTea(context.Background(), payload)
			So(err, ShouldBeNil)
			So(result, ShouldHaveLength, 2)

			So(result[0].Id, ShouldEqual, 2)
			So(result[0].Tea.String, ShouldEqual, "Green")
			So(result[1].Id, ShouldEqual, 4)
			So(result[1].Tea.String, ShouldEqual, "White")
		})

		Convey("limit 3", func() {
			payload := make(map[string]interface{})
			payload["tea"] = "%e%"
			payload["idList"] = []int64{1, 2, 3, 4, 5, 6, 7}
			payload["limit"] = int64(3)

			result, err := j.GetTicketsLikeTea(context.Background(), payload)
			So(err, ShouldBeNil)
			So(result, ShouldHaveLength, 2)

			So(result[0].Id, ShouldEqual, 2)
			So(result[0].Tea.String, ShouldEqual, "Green")
			So(result[1].Id, ShouldEqual, 4)
			So(result[1].Tea.String, ShouldEqual, "White")
		})
	})
}

func (s *JourneySuite) TestGetTicketsByOffset() {
	Convey("Add tickets for customer", s.T(), func() {
		Convey("First ticket", func() {
			now := time.Now().AddDate(-4, -5, -6)
			ticket := &another_model.Ticket{
				Id:         1,
				From:       1024,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "Black",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Second ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         2,
				From:       1025,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "Green",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Third ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         3,
				From:       1068,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "Oolong",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Fourth ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         4,
				From:       1099,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "White",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
	})

	Convey("Test GetTicketsByOffset()", s.T(), func() {
		Convey("limit 0, 2", func() {
			result, err := j.GetTicketsByOffset(context.Background(), &another_model.Ticket{
				CustomerId: 0,
			})
			So(err, ShouldBeNil)
			So(result, ShouldHaveLength, 2)

			So(result[0].Id, ShouldEqual, 1)
			So(result[0].Tea.String, ShouldEqual, "Black")
			So(result[1].Id, ShouldEqual, 2)
			So(result[1].Tea.String, ShouldEqual, "Green")
		})

		Convey("limit 1, 2", func() {
			result, err := j.GetTicketsByOffset(context.Background(), &another_model.Ticket{
				CustomerId: 1,
			})
			So(err, ShouldBeNil)
			So(result, ShouldHaveLength, 2)

			So(result[0].Id, ShouldEqual, 2)
			So(result[0].Tea.String, ShouldEqual, "Green")
			So(result[1].Id, ShouldEqual, 3)
			So(result[1].Tea.String, ShouldEqual, "Oolong")
		})

		Convey("limit 2, 2", func() {
			result, err := j.GetTicketsByOffset(context.Background(), &another_model.Ticket{
				CustomerId: 2,
			})
			So(err, ShouldBeNil)
			So(result, ShouldHaveLength, 2)

			So(result[0].Id, ShouldEqual, 3)
			So(result[0].Tea.String, ShouldEqual, "Oolong")
			So(result[1].Id, ShouldEqual, 4)
			So(result[1].Tea.String, ShouldEqual, "White")
		})

		Convey("limit 3, 2", func() {
			result, err := j.GetTicketsByOffset(context.Background(), &another_model.Ticket{
				CustomerId: 3,
			})
			So(err, ShouldBeNil)
			So(result, ShouldHaveLength, 1)

			So(result[0].Id, ShouldEqual, 4)
			So(result[0].Tea.String, ShouldEqual, "White")
		})
	})
}

func (s *JourneySuite) TestGetTicketsByFrom() {
	Convey("Add tickets for customer", s.T(), func() {
		Convey("First ticket", func() {
			now := time.Now().AddDate(-4, -5, -6)
			ticket := &another_model.Ticket{
				Id:         1,
				From:       1024,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "Black",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Second ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         2,
				From:       1025,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "Green",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Third ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         3,
				From:       1068,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "Oolong",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Fourth ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         4,
				From:       1099,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Tea: sql.NullString{
					String: "White",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
	})

	Convey("Test GetTicketsByFrom()", s.T(), func() {
		Convey("limit 2, 10", func() {
			result, err := j.GetTicketsByFrom(context.Background(), another_model.Ticket{
				From:       1024,
				CustomerId: 10,
			})
			So(err, ShouldBeNil)
			So(result, ShouldHaveLength, 0)
		})
	})
}

func (s *JourneySuite) TestGetTicketWhichTokenIsNotEqualTo() {
	Convey("Add tickets for customer", s.T(), func() {
		Convey("First ticket", func() {
			now := time.Now().AddDate(-4, -5, -6)
			ticket := &another_model.Ticket{
				Id:         1,
				From:       1024,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Price: sql.NullFloat64{
					Float64: 100,
					Valid:   true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "Black",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Second ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         2,
				From:       1025,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				Price: sql.NullFloat64{
					Float64: 90.21,
					Valid:   true,
				},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "Green",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Third ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         3,
				From:       1068,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				Price: sql.NullFloat64{
					Float64: 102.03,
					Valid:   true,
				},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "Oolong",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Fourth ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         4,
				From:       1099,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Price: sql.NullFloat64{
					Float64: 0.03,
					Valid:   true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "White",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
	})

	Convey("Test GetTicketWhichTokenIsNotEqualTo()", s.T(), func() {
		Convey("limit 2, 10", func() {
			result, err := j.GetTicketWhichTokenIsNotEqualTo(context.Background(), map[string]interface{}{
				"order": "???(ÒωÓױ)！",
				"from":  []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				"limit": 10,
			})
			So(err, ShouldBeNil)
			So(result.Id, ShouldEqual, 3)
			So(result.Token.String, ShouldEqual, "Oolong")
		})
	})
}

func (s *JourneySuite) TestGetTicketGreaterThanPrice() {
	Convey("Add tickets for customer", s.T(), func() {
		Convey("First ticket", func() {
			now := time.Now().AddDate(-4, -5, -6)
			ticket := &another_model.Ticket{
				Id:         1,
				From:       1024,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Price: sql.NullFloat64{
					Float64: 100,
					Valid:   true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "Black",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Second ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         2,
				From:       1025,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				Price: sql.NullFloat64{
					Float64: 90.21,
					Valid:   true,
				},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "Green",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Third ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         3,
				From:       1068,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				Price: sql.NullFloat64{
					Float64: 102.03,
					Valid:   true,
				},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "Oolong",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Fourth ticket", func() {
			now := time.Now().AddDate(-7, -8, -9)
			ticket := &another_model.Ticket{
				Id:         4,
				From:       1099,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: uint64(faker.RandomInt(1, 1024000)),
				Meal:       sql.NullInt64{},
				To: sql.NullInt32{
					Int32: 1024,
					Valid: true,
				},
				Price: sql.NullFloat64{
					Float64: 0.03,
					Valid:   true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "White",
					Valid:  true,
				},
				Climate:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddTicket(context.TODO(), ticket)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
	})

	Convey("Test GetTicketGreaterThanPrice()", s.T(), func() {
		Convey("price>10 AND id in (1, 2)", func() {
			result, err := j.GetTicketGreaterThanPrice(context.Background(), &another_model.Ticket{
				Price: sql.NullFloat64{
					Float64: 10.0,
					Valid:   true,
				},
			})
			So(err, ShouldBeNil)
			So(result.Id, ShouldEqual, 1)
			So(result.Token.String, ShouldEqual, "Black")
		})
	})
}

func (s *JourneySuite) TestUpdateCustomerMemberIDByID() {
	Convey("Add a customer", s.T(), func() {
		lastInsertID := int64(0)
		now := time.Now()

		customer := model.Customer{
			CardId:     int16(faker.RandomInt(-32768, 32767)),
			Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
			MemberId:   sql.NullInt32{},
			Email:      faker.Internet().Email(),
			Name:       faker.Name().Name(),
			Address:    sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: sql.NullTime{
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

	Convey("Test UpdateCustomerMemberIDByID()", s.T(), func() {
		result, err := j.UpdateCustomerMemberIDByID(context.Background(), &model.Customer{
			Id: 1,
			MemberId: sql.NullInt32{
				Int32: 100086,
				Valid: true,
			},
		})

		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 1)
		So(err, ShouldBeNil)

		getCustomer := model.Customer{}
		err = j.DB.Get(&getCustomer, "SELECT * FROM customer WHERE id = ?", 1)
		So(err, ShouldBeNil)
		So(getCustomer.Id, ShouldEqual, 1)
		So(getCustomer.MemberId.Int32, ShouldEqual, 100086)
	})
}

func (s *JourneySuite) TestUpdateCustomerMemberIDByIDList() {
	Convey("Add customers", s.T(), func() {
		Convey("Add first customer", func() {
			now := time.Now()

			customer := model.Customer{
				Id:     1,
				CardId: int16(faker.RandomInt(-32768, 32767)),
				Order:  int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId: sql.NullInt32{
					Int32: 1,
					Valid: true,
				},
				Email:      faker.Internet().Email(),
				Name:       faker.Name().Name(),
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddCustomer(context.TODO(), customer)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)

			getCustomer := model.Customer{}
			err = j.DB.Get(&getCustomer, "SELECT * FROM customer WHERE id = ?", 1)
			So(err, ShouldBeNil)

			customer.Id = 1
			So(getCustomer, ShouldResemble, customer)
		})

		Convey("Add second customer", func() {
			now := time.Now()

			customer := model.Customer{
				Id:     2,
				CardId: int16(faker.RandomInt(-32768, 32767)),
				Order:  int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId: sql.NullInt32{
					Int32: 2,
					Valid: true,
				},
				Email:      faker.Internet().Email(),
				Name:       faker.Name().Name(),
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddCustomer(context.TODO(), customer)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)

			getCustomer := model.Customer{}
			err = j.DB.Get(&getCustomer, "SELECT * FROM customer WHERE id = ?", 2)
			So(err, ShouldBeNil)

			customer.Id = 2
			So(getCustomer, ShouldResemble, customer)
		})

		Convey("Add third customer", func() {
			now := time.Now()

			customer := model.Customer{
				Id:     3,
				CardId: int16(faker.RandomInt(-32768, 32767)),
				Order:  int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId: sql.NullInt32{
					Int32: 3,
					Valid: true,
				},
				Email:      faker.Internet().Email(),
				Name:       "liupangzi",
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddCustomer(context.TODO(), customer)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)

			getCustomer := model.Customer{}
			err = j.DB.Get(&getCustomer, "SELECT * FROM customer WHERE id = ?", 3)
			So(err, ShouldBeNil)

			customer.Id = 3
			So(getCustomer, ShouldResemble, customer)
		})
	})

	Convey("Test UpdateCustomerMemberIDByIDList()", s.T(), func() {
		result, err := j.UpdateCustomerMemberIDByIDList(context.Background(), &model.Customer{
			MemberId: sql.NullInt32{
				Int32: 100086,
				Valid: true,
			},
		})

		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 2)
		So(err, ShouldBeNil)

		getCustomer := model.Customer{}
		err = j.DB.Get(&getCustomer, "SELECT * FROM customer WHERE member_id != 100086")
		So(err, ShouldBeNil)
		So(getCustomer.Id, ShouldEqual, 3)
		So(getCustomer.MemberId.Int32, ShouldEqual, 3)
	})
}

func (s *JourneySuite) TestDeleteCustomerByEmail() {
	email := "i@liuchao.me"

	Convey("Add a customer", s.T(), func() {
		lastInsertID := int64(0)
		now := time.Now()

		customer := model.Customer{
			CardId:     int16(faker.RandomInt(-32768, 32767)),
			Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
			MemberId:   sql.NullInt32{},
			Email:      email,
			Name:       faker.Name().Name(),
			Address:    sql.NullString{},
			CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
			UpdateTime: sql.NullTime{
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

	Convey("Test DeleteCustomerByEmail()", s.T(), func() {
		result, err := j.DeleteCustomerByEmail(context.Background(), map[string]interface{}{
			"email": email,
		})

		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 1)
		So(err, ShouldBeNil)

		getCustomer := model.Customer{}
		err = j.DB.Get(&getCustomer, "SELECT * FROM customer WHERE email = ?", email)
		So(err, ShouldEqual, sql.ErrNoRows)
	})
}

func (s *JourneySuite) TestDeleteCustomerByEmailOrIDList() {
	email1 := "i@liuchao.me"
	email2 := "i@example.com"
	email3 := "me@example.com"

	Convey("Add customers", s.T(), func() {
		Convey("Add first customer", func() {
			now := time.Now()

			customer := model.Customer{
				Id:         1,
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId:   sql.NullInt32{},
				Email:      email1,
				Name:       faker.Name().Name(),
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddCustomer(context.TODO(), customer)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)

			getCustomer := model.Customer{}
			err = j.DB.Get(&getCustomer, "SELECT * FROM customer WHERE id = ?", 1)
			So(err, ShouldBeNil)

			customer.Id = 1
			So(getCustomer, ShouldResemble, customer)
		})

		Convey("Add second customer", func() {
			now := time.Now()

			customer := model.Customer{
				Id:         2,
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId:   sql.NullInt32{},
				Email:      email2,
				Name:       faker.Name().Name(),
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddCustomer(context.TODO(), customer)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)

			getCustomer := model.Customer{}
			err = j.DB.Get(&getCustomer, "SELECT * FROM customer WHERE id = ?", 2)
			So(err, ShouldBeNil)

			customer.Id = 2
			So(getCustomer, ShouldResemble, customer)
		})

		Convey("Add third customer", func() {
			now := time.Now()

			customer := model.Customer{
				Id:         3,
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId:   sql.NullInt32{},
				Email:      email3,
				Name:       faker.Name().Name(),
				Address:    sql.NullString{},
				CreateTime: time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
				UpdateTime: sql.NullTime{
					Time:  time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location()),
					Valid: true,
				},
			}

			result, err := j.AddCustomer(context.TODO(), customer)
			So(err, ShouldBeNil)

			rowsAffected, err := result.RowsAffected()
			So(rowsAffected, ShouldEqual, 1)
			So(err, ShouldBeNil)

			getCustomer := model.Customer{}
			err = j.DB.Get(&getCustomer, "SELECT * FROM customer WHERE id = ?", 3)
			So(err, ShouldBeNil)

			customer.Id = 3
			So(getCustomer, ShouldResemble, customer)
		})
	})

	Convey("Test DeleteCustomerByEmailOrIDList()", s.T(), func() {
		result, err := j.DeleteCustomerByEmailOrIDList(context.Background(), map[string]interface{}{
			"email":   email2,
			"id_list": []uint64{1, 3},
		})

		So(err, ShouldBeNil)

		rowsAffected, err := result.RowsAffected()
		So(rowsAffected, ShouldEqual, 3)
		So(err, ShouldBeNil)

		err = j.DB.Get(&model.Customer{}, "SELECT * FROM customer")
		So(err, ShouldEqual, sql.ErrNoRows)
	})
}
