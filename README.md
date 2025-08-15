# 🛒 Ecommerce Platform (In Development)

A fully-featured e-commerce backend currently under active development.  
This project is being built with **Clean Architecture** and **Software Engineering Best Practices** to ensure long-term scalability, maintainability, and testability.

---

## 📌 Project Status
🚧 **Work in Progress** — This repository contains the ongoing implementation of the backend.  
Features, APIs, and modules are being gradually developed and tested.

---

## 🎯 Goals & Best Practices

- **Clean Architecture Layers:** Domain, Use Case, Interface Adapters, Frameworks & Drivers  
- **Domain-Driven Design (DDD):** Business logic separated from infrastructure concerns  
- **REST API Standards:** Consistent HTTP methods, status codes, and payload structures  
- **GitHub Workflow:** Commit conventions, PR reviews, and feature branches  
- **CI/CD Ready:** Plan for automated testing and deployment pipelines  
- **PostgreSQL Integration:** Scalable relational database backend  
- **Middleware Plan:** Logging, authentication, and CORS handling  

---

## 📂 Planned Project Structure
ecommerce/
│
├── cmd/
│ └── api/
│ └── main.go # App entry point: will initialize services, DB, handlers
│
├── internal/
│ ├── config/
│ │ └── db.go # Database connection setup
│ │
│ ├── domain/
│ │ ├── product.go # Product entity and validation rules
│ │ └── order.go # Order entity and business logic
│ │
│ ├── port/
│ │ └── repository.go # Repository interfaces
│ │
│ ├── usecase/
│ │ ├── product_service.go # Product-related use cases
│ │ └── order_service.go # Order-related use cases
│ │
│ ├── adapter/
│ │ ├── http/
│ │ │ ├── router.go # Route registration
│ │ │ ├── handler/
│ │ │ │ ├── product_handler.go
│ │ │ │ └── order_handler.go
│ │ │ └── middleware/ # Logging, auth, CORS
│ │ │
│ │ └── repo/
│ │ └── postgres/
│ │ ├── product_repo.go
│ │ └── order_repo.go
│
├── go.mod
└── README.md
---

## 🏗️ Clean Architecture Vision

- **Domain Layer** → Business entities and rules (independent of frameworks)  
- **Use Case Layer** → Application-specific logic that uses domain models  
- **Interface Adapters** → Bridges between domain and frameworks (HTTP, DB)  
- **Frameworks & Drivers** → External services like PostgreSQL, HTTP server  

---

## 📅 Development Roadmap

- [ ] Set up basic project structure with Go modules  
- [ ] Implement PostgreSQL connection and repository interfaces  
- [ ] Add product & order domain models  
- [ ] Create HTTP handlers and routing  
- [ ] Implement authentication middleware  
- [ ] Write unit tests and integration tests  
- [ ] Add CI/CD pipelines  
- [ ] Prepare production deployment scripts  

---

## 📜 License
This project will be licensed under the MIT License after its first release.
