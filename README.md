# PBKK-FP

### Group Member
|   Name        | NRP      | Class |
|-----------------|--------------|---------|
| Shafa Kirana Mulia      | 5025221078 | IUP |
| Moh. Azril Addia Ananda     | 5025221084 | IUP |

### Overview
This project is a web application for managing books, customers, and transactions in a library system. The application uses the Go programming language with the Gin framework and the GORM ORM for database management. This application allows admin to view, add, edit, and delete books and customers, and manage transactions such as borrowing and returning books.

### Features
1. **Book Management**:
    - **Create**: Add new books to the library with details such as name, author, and publication
    - **Read**: View a list of all books.
    - **Update**: Edit book details.
    - **Delete**: Remove books from the library.

2. **Customer Management**:
    - **Create**: Add new customers with details such as name, email, and phone number.
    - **Read**: View a list of all customers.
    - **Update**: Edit customer details.
    - **Delete**: Remove customers from the library system.

3. **Transaction Management**:
    - **Create**: Add new transactions for borrowing or returning books. Each transaction includes:
    - **Customer details**: Linked to a Customer entity.
    - **Book details**: Linked to a Book entity.
    - **Borrow date**: Date when the book was borrowed.
    - **Return date**: Date when the book is expected to be returned.
    - **Status**: Indicates the current state of the transaction (borrowed, returned, overdue).
    - **Read**: View a list of all transactions.
    - **Update**: Modify the status of a transaction.

### Technology Stack
- Backend: Go, Gin, GORM
- Database: SQLite (or any other GORM-supported database)
- Frontend: HTML/CSS (via Gin templates)

### Endpoints:

- `/books`: View, add, edit, or delete books.
- `/customers`: View, add, edit, or delete customers.
- `/transactions`: View, add, or edit transactions.

### Example Data
1. Book JSON
```json
{
  "name": "Book Title",
  "author": "Book Author",
  "publication": "Book Publication"
}

```

2. Customer JSON
```json
{
  "name": "Customer Name",
  "email": "customer@example.com",
  "phone": "123-456-7890"
}
```

3. Transaction JSON
```json
{
  "customer_id": 1,
  "book_id": 2,
  "borrow_date": "2024-12-11T10:00:00Z",
  "return_date": "2024-12-18T10:00:00Z",
  "status": "borrowed",
  "view_link": "/transactions/1"
}
```