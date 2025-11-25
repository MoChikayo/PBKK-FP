package service

import (
    "errors"
    "time"

    "github.com/MoChikayo/PBKK-FP/pkg/domain"
    "github.com/MoChikayo/PBKK-FP/pkg/repository"
)

type TransactionService interface {
    ListTransactions() ([]domain.Transaction, error)
    GetTransaction(id uint) (*domain.Transaction, error)
    BorrowBook(input domain.Transaction) (*domain.Transaction, error)
    ReturnBook(id uint) (*domain.Transaction, error)
    DeleteTransaction(id uint) error
}

type transactionService struct {
    repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
    return &transactionService{repo: repo}
}

func (s *transactionService) ListTransactions() ([]domain.Transaction, error) {
    return s.repo.FindAll()
}

func (s *transactionService) GetTransaction(id uint) (*domain.Transaction, error) {
    return s.repo.FindByID(id)
}

func (s *transactionService) BorrowBook(input domain.Transaction) (*domain.Transaction, error) {

    // Validation
    if input.CustomerID == 0 {
        return nil, errors.New("customer_id is required")
    }
    if input.BookID == 0 {
        return nil, errors.New("book_id is required")
    }

    // Borrow logic
    input.BorrowDate = time.Now()
    input.Status = domain.StatusBorrowed

    if err := s.repo.Create(&input); err != nil {
        return nil, err
    }

    return &input, nil
}

func (s *transactionService) ReturnBook(id uint) (*domain.Transaction, error) {
    tx, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }

    if tx.Status != domain.StatusBorrowed {
        return nil, errors.New("book is not currently borrowed")
    }

    tx.ReturnDate = time.Now()
    tx.Status = domain.StatusReturned

    if err := s.repo.Update(tx); err != nil {
        return nil, err
    }

    return tx, nil
}

func (s *transactionService) DeleteTransaction(id uint) error {
    return s.repo.Delete(id)
}
