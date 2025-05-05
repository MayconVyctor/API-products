# API-products

# 🛒 Product Management API

This is a **RESTful API** built with **Go (Golang)** to manage products, allowing for creation, retrieval, updating, and deletion.

---

## 🚀 Technologies Used

- **Go (Golang)**
- **Gin** – HTTP framework
- **PostgreSQL**
- **Docker**
- **DBeaver**
- **Git & GitHub**

---

## 📌 Features

- ✅ Create a product (`POST /products`)
- ✅ List all products (`GET /products`)
- ✅ Get a product by ID (`GET /products/:id`)
- ✅ Update a product (`PUT /products/:id`)
- ✅ Update only stock (`PATCH /products/:id/stock`)
- ✅ Delete a product (`DELETE /products/:id`)

---

## 📦 How to Run the Project

### Prerequisites:
- Go installed
- Docker and Docker Compose
- Git

### Steps:

```bash
# Clone the repository
git clone https://github.com/MayconVyctor/API-products.git
cd API-products

# Start the database container
docker-compose up -d

# Run the application
go run main.go


🧪 Testing
You can test the routes using Postman. A collection file can be added later with request examples.


📚 Learning Outcomes
This project strengthens skills in:

Building REST APIs with Go

Connecting to PostgreSQL databases

Organizing layers (Controller, Usecase, Repository)

Using Docker containers

Best practices with Git & GitHub


📫 Contact
LinkedIn: Maycon Vyctor

GitHub: @MayconVyctor
