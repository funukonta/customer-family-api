package services

import (
	"customer-data-api/internal/models"
	"customer-data-api/internal/repositories"
)

type CustomerService interface {
	Create(data *models.CreateCustomerReq) (*models.CreateCustomerRes, error)
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
