package service

import (
    "errors"

    "github.com/MoChikayo/PBKK-FP/pkg/domain"
    "github.com/MoChikayo/PBKK-FP/pkg/repository"
)

type BookService interface {
    ListBooks() ([]domain.Book, error)
    GetBook(id uint) (*domain.Book, error)
    CreateBook(input domain.Book) (*domain.Book, error)
    UpdateBook(id uint, input domain.Book) (*domain.Book, error)
    DeleteBook(id uint) error
}

type bookService struct {
    repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
    return &bookService{repo: repo}
}

func (s *bookService) ListBooks() ([]domain.Book, error) {
    return s.repo.FindAll()
}

func (s *bookService) GetBook(id uint) (*domain.Book, error) {
    return s.repo.FindByID(id)
}

func (s *bookService) CreateBook(input domain.Book) (*domain.Book, error) {

    if input.Name == "" {
        return nil, errors.New("book name is required")
    }
    if input.Author == "" {
        return nil, errors.New("author is required")
    }

    if err := s.repo.Create(&input); err != nil {
        return nil, err
    }

    return &input, nil
}

func (s *bookService) UpdateBook(id uint, input domain.Book) (*domain.Book, error) {
    existing, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }

    existing.Name = input.Name
    existing.Author = input.Author
    existing.Publication = input.Publication

    if err := s.repo.Update(existing); err != nil {
        return nil, err
    }

    return existing, nil
}

func (s *bookService) DeleteBook(id uint) error {
    return s.repo.Delete(id)
}
