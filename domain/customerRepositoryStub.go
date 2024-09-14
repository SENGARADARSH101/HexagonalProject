package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Adarsh", City: "Ahmedabad", ZipCode: "1234", DateOfBirth: "1999-01-01", Status: "active"},
		{Id: "1002", Name: "Dipendra", City: "Ahmedabad", ZipCode: "1234", DateOfBirth: "1999-01-01", Status: "active"},
	}
	return CustomerRepositoryStub{customers: customers}
}
