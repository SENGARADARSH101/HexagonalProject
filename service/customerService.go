package service

import (
	"HexagonalProject/domain"
	"HexagonalProject/domain/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
	GetAllActiveorInactiveCustomers(bool) ([]domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.FindById(id)
}

func (s DefaultCustomerService) GetAllActiveorInactiveCustomers(active bool) ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAllActiveOrInactive(active)
}
