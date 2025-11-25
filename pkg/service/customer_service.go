package service

import (
    "errors"

    "github.com/MoChikayo/PBKK-FP/pkg/domain"
    "github.com/MoChikayo/PBKK-FP/pkg/repository"
)

type CustomerService interface {
    ListCustomers() ([]domain.Customer, error)
    GetCustomer(id uint) (*domain.Customer, error)
    CreateCustomer(input domain.Customer) (*domain.Customer, error)
    UpdateCustomer(id uint, input domain.Customer) (*domain.Customer, error)
    DeleteCustomer(id uint) error
}

type customerService struct {
    repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
    return &customerService{repo: repo}
}

func (s *customerService) ListCustomers() ([]domain.Customer, error) {
    return s.repo.FindAll()
}

func (s *customerService) GetCustomer(id uint) (*domain.Customer, error) {
    return s.repo.FindByID(id)
}

func (s *customerService) CreateCustomer(input domain.Customer) (*domain.Customer, error) {

    // Basic validation
    if input.Name == "" {
        return nil, errors.New("name is required")
    }
    if input.Email == "" {
        return nil, errors.New("email is required")
    }

    if err := s.repo.Create(&input); err != nil {
        return nil, err
    }
    return &input, nil
}

func (s *customerService) UpdateCustomer(id uint, input domain.Customer) (*domain.Customer, error) {
    existing, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }

    // Update allowed fields
    existing.Name = input.Name
    existing.Email = input.Email
    existing.PhoneNumber = input.PhoneNumber
    existing.Address = input.Address

    if err := s.repo.Update(existing); err != nil {
        return nil, err
    }

    return existing, nil
}

func (s *customerService) DeleteCustomer(id uint) error {
    return s.repo.Delete(id)
}
