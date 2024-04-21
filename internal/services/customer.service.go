package services

import (
	"customer-data-api/internal/models"
	"customer-data-api/internal/repositories"
)

type CustomerService interface {
	Create(data *models.CreateCustomerReq) (*models.CreateCustomerRes, error)
	GetAllCustomer() (*models.GetAllCustomerRes, error)
	GetCustomer(id int) (*models.GetCustomerRes, error)
	UpdateCustomer(data *models.UpdateCustomerReq, id int) (*models.UpdateCustomerRes, error)
	DeleteCustomer(data *models.DeleteCustomerReq, custID int) error
}

type customerService struct {
	repo repositories.CustomerRepo
}

func NewCustomerService(repo repositories.CustomerRepo) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) Create(data *models.CreateCustomerReq) (*models.CreateCustomerRes, error) {
	customer := models.Customer{
		Name:          data.Name,
		DOB:           data.DOB,
		PhoneNum:      data.PhoneNum,
		Email:         data.Email,
		NationalityID: data.NationalityID,
	}

	customerData := models.CustomerData{
		FamMember: data.FamMember,
		Customer:  customer,
	}

	err := s.repo.Create(&customerData)
	if err != nil {
		return nil, err
	}

	response := &models.CreateCustomerRes{
		ID:            customerData.Customer.ID,
		Name:          data.Name,
		DOB:           data.DOB,
		PhoneNum:      data.PhoneNum,
		Email:         data.Email,
		NationalityID: data.NationalityID,
	}

	return response, nil
}

func (s *customerService) GetAllCustomer() (*models.GetAllCustomerRes, error) {
	customer, err := s.repo.GetAllCustomer()
	if err != nil {
		return nil, err
	}

	res := &models.GetAllCustomerRes{
		Data: customer,
	}

	return res, err
}

func (s *customerService) GetCustomer(id int) (*models.GetCustomerRes, error) {
	custData, err := s.repo.GetCustomer(id)
	if err != nil {
		return nil, err
	}

	response := &models.GetCustomerRes{
		ID:            custData.Customer.ID,
		Name:          custData.Customer.Name,
		DOB:           custData.Customer.DOB,
		PhoneNum:      custData.Customer.PhoneNum,
		Email:         custData.Customer.Email,
		NationalityID: custData.Customer.NationalityID,
		FamMember:     custData.FamMember,
	}

	return response, err
}

func (s *customerService) UpdateCustomer(data *models.UpdateCustomerReq, id int) (*models.UpdateCustomerRes, error) {
	customer := models.Customer{
		ID:            id,
		Name:          data.Name,
		DOB:           data.DOB,
		PhoneNum:      data.PhoneNum,
		Email:         data.Email,
		NationalityID: data.NationalityID,
	}

	customerData := models.CustomerData{
		FamMember: data.FamMember,
		Customer:  customer,
	}

	err := s.repo.Create(&customerData)
	if err != nil {
		return nil, err
	}

	response := &models.UpdateCustomerRes{
		ID:            customerData.Customer.ID,
		Name:          data.Name,
		DOB:           data.DOB,
		PhoneNum:      data.PhoneNum,
		Email:         data.Email,
		NationalityID: data.NationalityID,
		FamMember:     customerData.FamMember,
	}

	return response, nil
}

func (s *customerService) DeleteCustomer(data *models.DeleteCustomerReq, custID int) error {
	return s.repo.DeleteCustomer(data, custID)
}
