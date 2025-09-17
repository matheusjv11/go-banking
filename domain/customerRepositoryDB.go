package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/matheusjv11/go-banking/errs"
	"github.com/matheusjv11/go-banking/logger"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
		//rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
		//rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error querying customers table: " + err.Error())
		return nil, errs.NewUnexpectedError("Error querying customers table")
	}

	//err = sqlx.StructScan(rows, &customers)

	//if err != nil {
	//	logger.Error("Error scanning customers rows: " + err.Error())
	//	return nil, errs.NewUnexpectedError("Error scanning customers rows")
	//}

	//for rows.Next() {
	//	var c Customer
	//	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Dateofbirth, &c.Status)
	//	if err != nil {
	//		logger.Error("Error scanning customer row: " + err.Error())
	//		return nil, errs.NewUnexpectedError("Error scanning customer row")
	//	}
	//	customers = append(customers, c)
	//}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	// row := d.client.QueryRow(customerSql, id)

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Dateofbirth, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error scanning customer row: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{dbClient}
}
