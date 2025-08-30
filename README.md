
# 🛒 Ecommerce Platform

A modern e-commerce backend built with Go, following Clean Architecture principles and Domain-Driven Design patterns for scalability and maintainability.

## 📌 Project Status
🚧 **Active Development** — Core features are implemented with ongoing enhancements.

## ✨ Current Features

- **Product Management:** CRUD operations for product catalog
- **Domain Models:** Product, Cart, Order, Payment, Review, and User entities
- **RESTful API:** HTTP endpoints with proper routing
- **Middleware Stack:** CORS, logging, and custom middleware chain
- **In-Memory Storage:** Initial implementation with memory repository
- **Clean Architecture:** Separation of concerns across layers

## 🚀 Getting Started

### Prerequisites
- Go 1.24.2 or higher

### Installation
```bash
git clone https://github.com/mazhar75/ecommerce.git
cd ecommerce
go mod download
```

### Running the Application
```bash
go run main.go
```

The server will start on the configured port (default: 8080).

## 🔌 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/products` | Get all products |
| GET | `/products/{productId}` | Get product by ID |
| POST | `/products` | Create new product |

## 🏗️ Architecture

The project follows Clean Architecture with clear separation of concerns:

- **Domain Layer:** Core business entities and rules
- **Use Case Layer:** Application-specific business logic
- **Adapter Layer:** Interface adapters for HTTP handlers
- **Infrastructure Layer:** External frameworks and drivers

## 📂 Current Project Structure

```
ecommerce/
├── adapter/                      # Interface adapters
│   └── handlers/
│       └── product_handlers/     # Product HTTP handlers
│           ├── create_product.go
│           ├── get_product.go
│           ├── get_product_by_id.go
│           └── handler.go
├── cmd/                          # Application commands
│   ├── router.go                 # HTTP route registration
│   └── serve.go                  # Server initialization
├── config/                       # Configuration (empty - planned)
├── domain/                       # Business entities
│   ├── cart/
│   │   └── cart.go              # Cart entity
│   ├── order/
│   │   └── order.go             # Order entity
│   ├── payment/
│   │   └── payment.go           # Payment entity
│   ├── product/
│   │   └── product.go           # Product entity
│   ├── review/
│   │   └── review.go            # Review entity
│   └── user/
│       └── user.go              # User entity
├── drivers/                      # External drivers (empty - planned)
├── infra/                        # Infrastructure layer
│   ├── db/                      # Database (empty - planned)
│   └── memory/
│       └── product_repo.go      # In-memory product repository
├── interfaces/                   # Port interfaces (empty - planned)
├── middlewares/                  # HTTP middlewares
│   ├── cors.go                  # CORS middleware
│   ├── logger.go                # Request logging middleware
│   ├── manager.go               # Middleware chain manager
│   ├── test1.go                 # Test middleware 1
│   └── test2.go                 # Test middleware 2
├── usecase/                      # Business logic
│   └── product_service.go       # Product service implementation
├── go.mod                        # Go module file
├── main.go                       # Application entry point
└── README.md                     # This file
```

## 🛠️ Tech Stack

- **Language:** Go 1.24.2
- **Architecture:** Clean Architecture / Hexagonal Architecture
- **HTTP Server:** net/http (standard library)
- **Storage:** In-memory (transitioning to PostgreSQL)

## 🔄 Development Roadmap

### Phase 1: Foundation ✅
- [x] Project structure setup
- [x] Domain models definition
- [x] Basic product CRUD operations
- [x] Middleware implementation
- [x] In-memory repository

### Phase 2: Core Features (In Progress)
- [ ] User authentication and authorization
- [ ] Shopping cart functionality
- [ ] Order management system
- [ ] Payment integration
- [ ] Product reviews and ratings

### Phase 3: Infrastructure
- [ ] PostgreSQL database integration
- [ ] Database migrations
- [ ] Environment configuration
- [ ] Error handling improvements
- [ ] Input validation

### Phase 4: Advanced Features
- [ ] Search and filtering
- [ ] Inventory management
- [ ] Email notifications
- [ ] Admin dashboard API
- [ ] Analytics and reporting

### Phase 5: Production Ready
- [ ] Unit and integration tests
- [ ] API documentation (OpenAPI/Swagger)
- [ ] Docker containerization
- [ ] CI/CD pipeline
- [ ] Performance optimization
- [ ] Security hardening

## 👥 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is currently under development. License will be added upon first release.
