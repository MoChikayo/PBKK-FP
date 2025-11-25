package repository

import (
    "github.com/MoChikayo/PBKK-FP/pkg/config"
    "github.com/MoChikayo/PBKK-FP/pkg/domain"
)

type TransactionRepository interface {
    FindAll() ([]domain.Transaction, error)
    FindByID(id uint) (*domain.Transaction, error)
    Create(tx *domain.Transaction) error
    Update(tx *domain.Transaction) error
    Delete(id uint) error
}

type transactionRepository struct{}

func NewTransactionRepository() TransactionRepository {
    return &transactionRepository{}
}

func (r *transactionRepository) FindAll() ([]domain.Transaction, error) {
    db := config.GetDB()
    var txs []domain.Transaction

    err := db.
        Preload("Customer").
        Preload("Book").
        Find(&txs).
        Error

    if err != nil {
        return nil, err
    }
    return txs, nil
}

func (r *transactionRepository) FindByID(id uint) (*domain.Transaction, error) {
    db := config.GetDB()
    var tx domain.Transaction

    err := db.
        Preload("Customer").
        Preload("Book").
        First(&tx, id).
        Error

    if err != nil {
        return nil, err
    }
    return &tx, nil
}

func (r *transactionRepository) Create(tx *domain.Transaction) error {
    db := config.GetDB()
    return db.Create(tx).Error
}

func (r *transactionRepository) Update(tx *domain.Transaction) error {
    db := config.GetDB()
    return db.Save(tx).Error
}

func (r *transactionRepository) Delete(id uint) error {
    db := config.GetDB()
    return db.Delete(&domain.Transaction{}, id).Error
}
