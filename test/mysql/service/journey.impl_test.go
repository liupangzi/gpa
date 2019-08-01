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

func (s *JourneySuite) TestAddTicket() {
	Convey("Test AddTicket() w/ a pre-generated primary key", s.T(), func() {
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

	Convey("Test AddTicket() w/o a primary key", s.T(), func() {
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

func (s *JourneySuite) TestUpdateCustomerNameByEmail() {
	email := faker.Internet().Email()
	name := faker.Name().String()

	Convey("Add customer", s.T(), func() {
		now := time.Now().AddDate(-1, 2, 3)
		customer := model.Customer{
			CardId:     int16(faker.RandomInt(-32768, 32767)),
			Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
			MemberId:   sql.NullInt64{},
			Email:      email,
			Name:       faker.Name().String(),
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

func (s *JourneySuite) TestDeleteCustomerByID() {
	ID := uint64(100)

	Convey("Add customer w/ id == 100", s.T(), func() {
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

func (s *JourneySuite) TestGetCustomerByID() {
	Convey("Test GetCustomerByID", s.T(), func() {
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
			Operation: []byte{100, 200, 102, 140, 150},
			SecretKey: []byte{26, 67, 48, 23, 120},
			Token: sql.NullString{
				String: "1234567890zxcvb",
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

			ticket := &model.Ticket{
				From:       1024,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: customerID,
				Price: sql.NullFloat64{
					Float64: 1.023,
					Valid:   true,
				},
				Meal: sql.NullInt64{},
				To: sql.NullInt64{
					Int64: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "1234567890zxcvb",
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
		})

		Convey("Second ticket", func() {
			lastInsertID := int64(0)
			now := time.Now().AddDate(-4, -5, -6)

			ticket := &model.Ticket{
				From:       1025,
				Duration:   float64(faker.RandomInt(1, 1024000)),
				CustomerId: customerID,
				Price: sql.NullFloat64{
					Float64: 1.023,
					Valid:   true,
				},
				Meal: sql.NullInt64{},
				To: sql.NullInt64{
					Int64: 1024,
					Valid: true,
				},
				Operation: []byte{100, 200, 102, 140, 150},
				SecretKey: []byte{26, 67, 48, 23, 120},
				Token: sql.NullString{
					String: "1234567890zxcvb",
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
				MemberId:   sql.NullInt64{},
				Email:      email,
				Name:       name,
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
		})

		Convey("Second customer", func() {
			lastInsertID := int64(0)
			now := time.Now()

			customer := model.Customer{
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      -16,
				MemberId:   sql.NullInt64{},
				Email:      faker.Internet().Email(),
				Name:       faker.Name().String(),
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
		})

		Convey("Third customer", func() {
			lastInsertID := int64(0)
			now := time.Now()

			customer := model.Customer{
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      -8,
				MemberId:   sql.NullInt64{},
				Email:      email,
				Name:       name,
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
				MemberId:   sql.NullInt64{},
				Email:      email1,
				Name:       name,
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
		})

		Convey("Second customer", func() {
			lastInsertID := int64(0)
			now := time.Now()

			customer := model.Customer{
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId:   sql.NullInt64{},
				Email:      faker.Internet().Email(),
				Name:       faker.Name().String(),
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
		})

		Convey("Third customer", func() {
			lastInsertID := int64(0)
			now := time.Now()

			customer := model.Customer{
				CardId:     int16(faker.RandomInt(-32768, 32767)),
				Order:      int32(faker.RandomInt(-2147483648, 2147483647)),
				MemberId:   sql.NullInt64{},
				Email:      email2,
				Name:       name,
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
			MemberId:   sql.NullInt64{},
			Email:      email,
			Name:       name,
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
	})

	Convey("Test GetCustomerNameByEmail()", s.T(), func() {
		result, err := j.GetCustomerNameByEmail(context.Background(), email)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, name)
	})
}
