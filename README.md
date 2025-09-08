
# 🛒 Ecommerce Platform

A modern e-commerce backend built with Go, following Clean Architecture principles and Domain-Driven Design patterns for scalability and maintainability.

## 📌 Project Status
🚧 **Active Development** — Core features are implemented with ongoing enhancements.

## ✨ Current Features

- **Product Management:** Full CRUD operations for product catalog with PostgreSQL persistence
- **Domain Models:** Product, Cart, Order, Payment, Review, and User entities
- **RESTful API:** HTTP endpoints with proper routing
- **Middleware Stack:** CORS, logging, and custom middleware chain
- **PostgreSQL Storage:** Database-backed product repository with migrations
- **Clean Architecture:** Separation of concerns across layers

## 🚀 Getting Started

### Prerequisites
- Go 1.24.2 or higher
- PostgreSQL 12+ (for database functionality)
- Environment file (.env) with required configurations

### Installation
```bash
git clone https://github.com/mazhar75/ecommerce.git
cd ecommerce
go mod download
```

### Database Setup
1. Ensure PostgreSQL is running on your system
2. Create the database:
```sql
CREATE DATABASE ecommerce;
```
3. Run migrations (manually for now):
```bash
psql -U postgres -d ecommerce -f migrations/001_init_schema.up.sql
```

### Configuration
Create a `.env` file in the project root with the following variables:
```env
SERVICE_NAME=ecommerce
VERSION=1.0.0
PORT=9090
DB_HOST=localhost
DB_PORT=5432
DB_NAME=ecommerce
DB_USER=postgres
DB_PASSWORD=your_password
```

### Running the Application
```bash
go run main.go
```

The server will start on the configured port (default: 9090).

## 🔌 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check endpoint |
| GET | `/products` | Get all products |
| GET | `/products/{productId}` | Get product by ID |
| POST | `/products` | Create new product |
| PUT | `/products/{productId}` | Update existing product |
| DELETE | `/products/{productId}` | Delete product |

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
│   ├── handlers/
│   │   ├── health_handler/       # Health check handler
│   │   │   └── health.go
│   │   └── product_handlers/     # Product HTTP handlers
│   │       ├── create_product.go
│   │       ├── delete_product.go
│   │       ├── get_product.go
│   │       ├── get_product_by_id.go
│   │       ├── handler.go
│   │       └── update_product.go
│   └── routes/                   # Route interfaces
│       └── global_routes.go      # Route registration interface
├── cmd/                          # Application commands
│   ├── router.go                 # HTTP route registration
│   └── serve.go                  # Server initialization with DB
├── config/                       # Configuration management
│   ├── config.go                 # Configuration structure
│   └── loadenv.go                # Environment loader
├── domain/                       # Business entities
│   ├── cart/
│   │   └── cart.go              # Cart entity
│   ├── health/
│   │   └── health.go            # Health entity
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
│   ├── memory/                   # In-memory repositories
│   │   └── health_repo.go       # Health repository
│   └── postgresql/               # PostgreSQL integration
│       ├── db.go                 # Database connection
│       └── product_repo.go      # PostgreSQL product repository
├── interfaces/                   # Port interfaces (empty - planned)
├── middlewares/                  # HTTP middlewares
│   ├── cors.go                  # CORS middleware
│   ├── logger.go                # Request logging middleware
│   ├── manager.go               # Middleware chain manager
│   ├── test1.go                 # Test middleware 1
│   └── test2.go                 # Test middleware 2
├── migrations/                   # Database migrations
│   ├── 001_init_schema.down.sql # Initial schema rollback
│   └── 001_init_schema.up.sql   # Initial schema setup
├── usecase/                      # Business logic
│   ├── health_service.go        # Health service implementation
│   └── product_service.go       # Product service implementation
├── .env                          # Environment variables (not in repo)
├── go.mod                        # Go module file
├── go.sum                        # Go dependencies lock file
├── main.go                       # Application entry point
└── README.md                     # This file
```

## 🛠️ Tech Stack

- **Language:** Go 1.24.2
- **Architecture:** Clean Architecture / Hexagonal Architecture
- **HTTP Server:** net/http (standard library)
- **Database:** PostgreSQL 12+ (with lib/pq driver)
- **Storage:** PostgreSQL repositories (migrated from in-memory)
- **Configuration:** Environment-based configuration

## 🔄 Development Roadmap

### Phase 1: Foundation ✅
- [x] Project structure setup
- [x] Domain models definition
- [x] Full product CRUD operations (Create, Read, Update, Delete)
- [x] Middleware implementation
- [x] In-memory repository
- [x] Health check endpoint
- [x] Route registration interface

### Phase 2: Core Features (In Progress)
- [ ] User authentication and authorization
- [ ] Shopping cart functionality
- [ ] Order management system
- [ ] Payment integration
- [ ] Product reviews and ratings

### Phase 3: Infrastructure ✅ (Partially Complete)
- [x] PostgreSQL database integration
- [x] Environment configuration
- [x] Database migrations (initial schema)
- [x] Product repository with PostgreSQL
- [ ] Error handling improvements
- [ ] Input validation
- [ ] Additional migrations for other entities

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
