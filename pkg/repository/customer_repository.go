package repository

import (
    "github.com/MoChikayo/PBKK-FP/pkg/config"
    "github.com/MoChikayo/PBKK-FP/pkg/domain"
)

type CustomerRepository interface {
    FindAll() ([]domain.Customer, error)
    FindByID(id uint) (*domain.Customer, error)
    Create(customer *domain.Customer) error
    Update(customer *domain.Customer) error
    Delete(id uint) error
}

type customerRepository struct{}

func NewCustomerRepository() CustomerRepository {
    return &customerRepository{}
}

func (r *customerRepository) FindAll() ([]domain.Customer, error) {
    db := config.GetDB()
    var customers []domain.Customer
    if err := db.Find(&customers).Error; err != nil {
        return nil, err
    }
    return customers, nil
}

func (r *customerRepository) FindByID(id uint) (*domain.Customer, error) {
    db := config.GetDB()
    var customer domain.Customer
    if err := db.First(&customer, id).Error; err != nil {
        return nil, err
    }
    return &customer, nil
}

func (r *customerRepository) Create(customer *domain.Customer) error {
    db := config.GetDB()
    return db.Create(customer).Error
}

func (r *customerRepository) Update(customer *domain.Customer) error {
    db := config.GetDB()
    return db.Save(customer).Error
}

func (r *customerRepository) Delete(id uint) error {
    db := config.GetDB()
    return db.Delete(&domain.Customer{}, id).Error
}
