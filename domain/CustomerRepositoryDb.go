package domain

import (
	"HexagonalProject/domain/errs"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAllActiveOrInactive(active bool) ([]Customer, *errs.AppError) {
	findAllSql := "select * from customers where status =?"

	row, err := d.client.Query(findAllSql, active)
	if err != nil {
		log.Println("Error while scanning customer table " + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected Databse Error")
	}
	customers := make([]Customer, 0)

	for row.Next() {
		var customer Customer
		err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.DateOfBirth, &customer.Status, &customer.ZipCode)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("Customers Not Found")
			} else {
				log.Println("Error while scanning customer table " + err.Error())
				return nil, errs.NewUnExpectedError("Unexpected DataBase Error")
			}
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select * from customers"

	row, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while scanning customer table " + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected Databse Error")
	}
	customers := make([]Customer, 0)

	for row.Next() {
		var customer Customer
		err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.DateOfBirth, &customer.Status, &customer.ZipCode)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("Customers Not Found")
			} else {
				log.Println("Error while scanning customer table " + err.Error())
				return nil, errs.NewUnExpectedError("Unexpected DataBase Error")
			}
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(customerId string) (*Customer, *errs.AppError) {
	findByIdSql := "select * from customers where customer_id=?"

	row := d.client.QueryRow(findByIdSql, customerId)

	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.DateOfBirth, &customer.Status, &customer.ZipCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else {
			log.Println("Error while scanning customer data with id " + err.Error())
			return nil, errs.NewUnExpectedError("Unexpected DataBase Error")
		}
	}
	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}

}
