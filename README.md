
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

```plaintext

          ┌───────────────────────┐
          │        main.go          │
          │ (cmd/api/main.go)       │
          │                         │
          │ - Loads config/env      │
          │ - Initializes DB Conn   │
          │ - Creates Repo (DB impl)│
          │ - Injects Repo into     │
          │   Usecase(Service)      │
          │ - Injects Service into  │
          │   HTTP Handler          │
          │ - Starts HTTP Router    │
          └───────────┬─────────────┘
                      │
                      ▼
          ┌───────────────────────────┐
          │  Frameworks / HTTP Router │
          │ (internal/frameworks/http)│
          │                           │
          │ - Registers routes        │
          │ - Attaches middleware     │
          │ - Delegates to Handlers   │
          └───────────┬───────────────┘
                      │
                      ▼
          ┌───────────────────────────┐
          │     Handler / Adapter     │
          │ (internal/adapter/http)   │
          │                           │
          │ - Receives HTTP requests  │
          │ - Maps Req → Usecase DTO  │
          │ - Calls Usecase methods   │
          │ - Returns response        │
          └───────────┬───────────────┘
                      │
                      ▼
          ┌───────────────────────────┐
          │      Usecase / Service    │
          │ (internal/usecase)        │
          │                           │
          │ - Business logic rules    │
          │ - Depends on Port         │
          │   (ProductRepository)     │
          └───────────┬───────────────┘
                      │
                      ▼
          ┌───────────────────────────┐
          │       Port / Interface    │
          │ (internal/port)           │
          │                           │
          │ - Defines contract:       │
          │   ProductRepository       │
          │ - e.g. GetAll(), GetByID()│
          └───────────┬───────────────┘
                      │
                      ▼
          ┌───────────────────────────┐
          │  Adapter + Framework Repo │
          │ (adapter/repo +           │
          │  frameworks/db/postgres)  │
          │                           │
          │ - Mapper (domain <-> row) │
          │ - Concrete DBRepo impl    │
          │ - Uses DB driver/sql/gorm │
          │ - Implements Port         │
          └───────────────────────────┘




ecommerce/
│
├── cmd/
│   └── api/
│       └── main.go               # Entry point: init config, DB, repos, usecases, router
│
├── internal/
│   ├── config/
│   │   ├── app.go                # App-level config (env, port, secrets)
│   │   └── db.go                 # DB connection setup (driver, pool)
│   │
│   ├── domain/
│   │   ├── product.go            # Product entity + validation
│   │   └── order.go              # Order entity + validation
│   │
│   ├── port/                     # Interfaces (Ports)
│   │   ├── product_repository.go # ProductRepository interface
│   │   └── order_repository.go   # OrderRepository interface
│   │
│   ├── usecase/                  # Application services (business logic)
│   │   ├── product_service.go
│   │   └── order_service.go
│   │
│   ├── adapter/                  # Interface Adapters (domain ↔ external)
│   │   ├── http/
│   │   │   ├── handler/
│   │   │   │   ├── product_handler.go
│   │   │   │   └── order_handler.go
│   │   │   └── middleware/       # Cross-cutting concerns (logging, auth)
│   │   │
│   │   └── repository/
│   │       └── mapper.go         # DB row ↔ domain entity mapping
│   │
│   └── frameworks/               # Frameworks & Drivers
│       ├── http/
│       │   └── router.go         # Gin/net/http router setup, middleware wiring
│       └── db/
│           ├── postgres_conn.go  # DB driver init, connection pool
│           ├── product_repo.go   # Concrete repo implementation
│           └── order_repo.go
│
├── pkg/                          # Optional shared utilities (logger, errors)
│
├── go.mod
└── README.md

```
---

## 🏗️ Clean Architecture Vision

- **Domain Layer** → Business entities and rules (independent of frameworks)  
- **Use Case Layer** → Application-specific logic that uses domain models  
- **Interface Adapters** → Bridges between domain and frameworks (HTTP, DB)  
- **Frameworks & Drivers** → External services like PostgreSQL, HTTP server  

---


## 📜 License
This project will be licensed under the MIT License after its first release.
