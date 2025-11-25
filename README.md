# 📚 PBKK Final Project — Bookstore Management System  
A full-stack Bookstore Management System built with **Go**, **Gin Web Framework**, **Clean Architecture**, and **MySQL**, featuring both **REST APIs** and a **template-based UI** for managing books, customers, and transactions.

---

# 🚀 Features

### ✅ Books Management
- Add new books  
- Update book details  
- Delete books  
- View list of available books  

### ✅ Customer Management
- Create customer records  
- Edit details  
- Delete customers  
- View customer list  

### ✅ Transaction Management
- Borrow books  
- Return books  
- Track due dates  
- Automatic book availability handling  

### ✅ UI + API
- Fully functional UI using Go HTML templates  
- RESTful JSON API endpoints  
- Clean folder structure  

---

# 🏗️ Project Architecture (Clean Architecture)

```
PBKK-FP/
│
├── cmd/
│   └── server/
│       └── main.go                # Application entry point
│
├── pkg/
│   ├── config/                    # DB configuration
│   │   └── db.go
│   │
│   ├── domain/                    # Entities (Book, Customer, Transaction)
│   │   ├── book.go
│   │   ├── customer.go
│   │   └── transaction.go
│   │
│   ├── repository/                # Repository layer
│   │   ├── book_repository.go
│   │   ├── customer_repository.go
│   │   └── transaction_repository.go
│   │
│   ├── service/                   # Business logic
│   │   ├── book_service.go
│   │   ├── customer_service.go
│   │   └── transaction_service.go
│   │
│   ├── delivery/                  # HTTP handlers
│   │   └── http/
│   │       ├── book_handler.go
│   │       ├── customer_handler.go
│   │       └── transaction_handler.go
│   │
│   ├── routes/                    # Route definitions
│   │   └── router.go
│   │
│   └── utils/
│       └── response.go            # Standard API responses
│
├── templates/                     # HTML UI templates
│   ├── dashboard.html
│   ├── books.list.html
│   ├── books.create.html
│   ├── books.update.html
│   ├── customers.list.html
│   ├── customers.create.html
│   ├── transactions.list.html
│   ├── transactions.create.html
│   └── components/
│       ├── header.html
│       ├── footer.html
│       └── sidebar.html
│
├── go.mod
└── go.sum
```

---

# 🛢️ Database

This project uses **MySQL**.

### Example `.env` (if you use one later)
```
DB_USER=root
DB_PASS=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=bookstore
```

---

# ⚙️ Installation & Setup

### **1. Clone the repository**
```bash
git clone https://github.com/<your-username>/PBKK-FP.git
cd PBKK-FP
```

### **2. Install dependencies**
```bash
go mod tidy
```

### **3. Create a MySQL database**
```sql
CREATE DATABASE bookstore;
```

### **4. Run the server**
```bash
go run ./cmd/server
```

Server will start at:

```
http://localhost:9010/
```

---

# 🌐 UI Routes

| Page | URL |
|------|-----|
| Dashboard | `/` |
| Books List | `/books` |
| Create Book | `/books/new` |
| Update Book | `/books/update/:id` |
| Customers | `/customers` |
| Transactions | `/transactions` |

---

# 🧩 API Endpoints

## 📘 Books API
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/api/books` | Get all books |
| GET    | `/api/books/:id` | Get book by ID |
| POST   | `/api/books` | Create new book |
| PUT    | `/api/books/:id` | Update book |
| DELETE | `/api/books/:id` | Delete book |

## 👤 Customers API
| Method | Endpoint |
|--------|----------|
| GET | `/api/customers` |
| GET | `/api/customers/:id` |
| POST | `/api/customers` |
| PUT | `/api/customers/:id` |
| DELETE | `/api/customers/:id` |

## 🔄 Transactions API
| Method | Endpoint |
|--------|----------|
| GET | `/api/transactions` |
| GET | `/api/transactions/:id` |
| POST | `/api/transactions/borrow` |
| PUT | `/api/transactions/return/:id` |
| DELETE | `/api/transactions/:id` |

---

# 📄 License
This project is open-source — use freely for learning and development.

