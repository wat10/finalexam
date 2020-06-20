package customer

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// driver
	_ "github.com/lib/pq"
	"github.com/wat10/finalexam/customer/errors"
)

var db *sql.DB

func init() {
	fmt.Println("data access init")
	var err error
	url := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
}

// insertCustomer is a function for insert new customer into database
func insertCustomer(customer *Customer) *errors.Error {
	insertStmt := `INSERT INTO customers (
		name,
		email,
		status
	) VALUES ($1, $2, $3) RETURNING id;`

	row := db.QueryRow(insertStmt, customer.Name, customer.Email, customer.Status)

	var id int
	scanerr := row.Scan(&id)
	if scanerr != nil {
		return &errors.Error{
			Code:        111,
			Message:     "insert error",
			OriginError: scanerr,
		}
	}

	customer.ID = id

	return nil
}

func getCustomers() ([]*Customer, *errors.Error) {
	queryStmt := `SELECT * FROM customers;`

	stmt, prepareerr := db.Prepare(queryStmt)
	if prepareerr != nil {
		return nil, &errors.Error{
			Code:        111,
			Message:     "select error",
			OriginError: prepareerr,
		}
	}

	rows, _ := stmt.Query()

	var id int
	var name, email, status string
	customers := []*Customer{}
	for rows.Next() {
		scanerr := rows.Scan(&id, &name, &email, &status)
		if scanerr != nil {
			return nil, &errors.Error{
				Code:        111,
				Message:     "select error",
				OriginError: scanerr,
			}
		}

		fmt.Printf("select success: id:%#v name:%#v email:%#v status:%#v\n", id, name, email, status)
		customers = append(customers, &Customer{
			ID: id, Name: name, Email: email, Status: status,
		})
	}
	return customers, nil
}

func getCustomerByID(id string) (*Customer, *errors.Error) {
	queryStmt := `SELECT * FROM customers WHERE id=$1;`

	stmt, prepareerr := db.Prepare(queryStmt)
	if prepareerr != nil {
		return nil, &errors.Error{
			Code:        111,
			Message:     "select error",
			OriginError: prepareerr,
		}
	}

	row := stmt.QueryRow(id)

	var cid int
	var name, email, status string
	scanerr := row.Scan(&cid, &name, &email, &status)
	if scanerr != nil {
		return nil, &errors.Error{
			Code:        111,
			Message:     "select error",
			OriginError: scanerr,
		}
	}

	return &Customer{
		ID:     cid,
		Name:   name,
		Email:  email,
		Status: status,
	}, nil
}

func updateCustomerByID(customer *Customer) *errors.Error {
	updateStmt := `UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1;`
	stmt, prepareerr := db.Prepare(updateStmt)
	if prepareerr != nil {
		return &errors.Error{
			Code:        111,
			Message:     "update error",
			OriginError: prepareerr,
		}
	}

	_, execerr := stmt.Exec(customer.ID, customer.Name, customer.Email, customer.Status)
	if execerr != nil {
		return &errors.Error{
			Code:        111,
			Message:     "update error",
			OriginError: execerr,
		}
	}

	return nil
}

func deleteCustomerByID(id string) *errors.Error {
	delStmt := `DELETE FROM customers WHERE id=$1;`
	stmt, prepareerr := db.Prepare(delStmt)
	if prepareerr != nil {
		return &errors.Error{
			Code:        111,
			Message:     "delete error",
			OriginError: prepareerr,
		}
	}

	_, execerr := stmt.Exec(id)
	if execerr != nil {
		return &errors.Error{
			Code:        111,
			Message:     "delete error",
			OriginError: execerr,
		}
	}
	return nil
}
