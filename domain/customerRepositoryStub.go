package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "123", Name: "Matheus", City: "Porto Alegre", Zipcode: "12345", Dateofbirth: "1999-09-06", Status: "1"},
		{Id: "124", Name: "Jade", City: "Porto Alegre", Zipcode: "12345", Dateofbirth: "1999-09-06", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
