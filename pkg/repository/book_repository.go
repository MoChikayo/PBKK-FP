package repository

import (
    "github.com/MoChikayo/PBKK-FP/pkg/config"
    "github.com/MoChikayo/PBKK-FP/pkg/domain"
)

type BookRepository interface {
    FindAll() ([]domain.Book, error)
    FindByID(id uint) (*domain.Book, error)
    Create(book *domain.Book) error
    Update(book *domain.Book) error
    Delete(id uint) error
}

type bookRepository struct{}

func NewBookRepository() BookRepository {
    return &bookRepository{}
}

func (r *bookRepository) FindAll() ([]domain.Book, error) {
    db := config.GetDB()
    var books []domain.Book
    if err := db.Find(&books).Error; err != nil {
        return nil, err
    }
    return books, nil
}

func (r *bookRepository) FindByID(id uint) (*domain.Book, error) {
    db := config.GetDB()
    var book domain.Book
    if err := db.First(&book, id).Error; err != nil {
        return nil, err
    }
    return &book, nil
}

func (r *bookRepository) Create(book *domain.Book) error {
    return config.GetDB().Create(book).Error
}

func (r *bookRepository) Update(book *domain.Book) error {
    return config.GetDB().Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
    return config.GetDB().Delete(&domain.Book{}, id).Error
}
