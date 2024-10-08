package domain

import "HexagonalProject/domain/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindById(id string) (*Customer, *errs.AppError)
	FindAllActiveOrInactive(bool) ([]Customer, *errs.AppError)
}
