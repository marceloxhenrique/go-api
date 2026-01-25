# Go Products API

A simple RESTful API built with **Go** and **Gin** for managing products.  
This project follows a layered architecture using **controller → usecase → repository → database**.

---

## 🚀 Features

- Ping health check
- List all products
- Get product by ID
- Create a new product
- Delete a product by ID

---

## 🧱 Project Structure

```text
go-products/
│
├── controller/    # HTTP handlers (Gin controllers)
├── usecase/       # Business logic layer
├── repository/   # Database access layer
├── db/           # Database connection
├── model/        # Domain models
├── main.go       # Application entry point
```
