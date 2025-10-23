
# 🛒 Ecommerce Platform

A modern e-commerce backend built with Go, following Clean Architecture principles and Domain-Driven Design patterns for scalability and maintainability.

## 📌 Project Status
🚧 **Active Development** — Core authentication ,product management features, shopping cart functionalities are implemented with ongoing enhancements.

## ✨ Current Features

### Implemented ✅
- **Authentication System:** User registration and login with email/password
- **Custom JWT Implementation:** From-scratch JWT token generation and validation
- **Product Management:** Full CRUD operations for product catalog with PostgreSQL persistence
- **Domain Models:** Product, Cart, Order, Payment, Review, and User entities
- **RESTful API:** HTTP endpoints with proper routing
- **Middleware Stack:** CORS, logging, and custom middleware chain
- **PostgreSQL Storage:** Database-backed repositories with migrations
- **Clean Architecture:** Separation of concerns across layers
- **Error Handling:** Custom AppError type with HTTP status mapping
- **JWT AUTH:** JWT token integration in authentication flow, Protected route middleware
- **Shopping Cart:** Shopping cart functionalities ( add,edit and delete from cart )

### In Progress 🚧
- Shopping cart functionality
- Order management system

### Planned 📋
- Payment processing
- Product reviews and ratings
- Input validation and sanitization
- Unit and integration tests
- Google Oauth2
- Rate limiting using redis

## 🚀 Getting Started

### Prerequisites
- Go 1.24.2 or higher
- PostgreSQL 12+ 
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
3. Run migrations to create tables:
```bash
# Create tables
psql -U postgres -d ecommerce -f migrations/001_init_schema.up.sql
psql -U postgres -d ecommerce -f migrations/002_init_schema.up.sql

# To rollback (if needed)
psql -U postgres -d ecommerce -f migrations/001_init_schema.down.sql
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
JWT_SECRET=your jwt secret
```

### Running the Application
```bash
go run main.go
```

The server will start on the configured port (default: 9090).

## 🔌 API Endpoints

### Authentication Endpoints
| Method | Endpoint | Description | Status |
|--------|----------|-------------|--------|
| POST | `/auth/register` | Register new user with email/password | ✅ Implemented |
| POST | `/auth/login` | Authenticate user and get token | ✅ Implemented |

### Product Management Endpoints  
| Method | Endpoint | Description | Status |
|--------|----------|-------------|--------|
| GET | `/products` | Get all products | ✅ Implemented |
| GET | `/products/{productId}` | Get product by ID | ✅ Implemented |
| POST | `/products` | Create new product | ✅ Implemented |
| PUT | `/products/{productId}` | Update existing product | ✅ Implemented |
| DELETE | `/products/{productId}` | Delete product | ✅ Implemented |

### Category Management Endpoints  
| Method | Endpoint | Description | Status |
|--------|----------|-------------|--------|
| POST | `/category` | Create new category | ✅ Implemented |

### Shopping Cart Endpoints  
| Method | Endpoint | Description | Status |
|--------|----------|-------------|--------|
| GET |  | `/cart/{userId}`                      | Get cart with all cart items               | ✅ Implemented |
| POST   | `/cart/{user_id}/add`                 | Add a product to the user's cart           | ✅ Implemented |
| PATCH  | `/cart/{user_id}/update`              | Update product quantity in the user's cart | ✅ Implemented |
| DELETE | `/cart/{user_id}/remove/{product_id}` | Remove a product from the user's cart      | ✅ Implemented |



### Health Check
| Method | Endpoint | Description | Status |
|--------|----------|-------------|--------|
| GET | `/health` | System health check | ✅ Implemented |

### Example API Requests

#### Register User
```bash
curl -X POST http://localhost:9090/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword",
    "is_varified":false,
    "is_shop_owner":true
  }'
```

#### Login User
```bash
curl -X POST http://localhost:9090/auth/login \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <access_token>" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

#### Create Product
```bash
curl -X POST http://localhost:9090/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "description": "High-performance laptop",
    "type": "Electronics",
    "price": 999.99,
    "img_url": "https://example.com/laptop.jpg",
    "category_id": 1
  }'
```

## 🏗️ Architecture

The project follows Clean Architecture with clear separation of concerns:

- **Domain Layer:** Core business entities and rules (Product, User, Cart, Order, Payment, Review)
- **Use Case Layer:** Application-specific business logic (services and use cases)
- **Adapter Layer:** Interface adapters for HTTP handlers and route registration
- **Infrastructure Layer:** External frameworks, database implementations, and drivers

### Key Architectural Patterns
- **Repository Pattern:** Clear interfaces with PostgreSQL implementations
- **Dependency Injection:** Constructor-based DI for loose coupling
- **Middleware Chain:** Composable middleware with manager pattern
- **Custom Error Handling:** AppError type with HTTP status mapping
- **Interface Segregation:** Domain interfaces separate from implementations

## 📂 Project Structure

```
ecommerce/
├── adapter/                      # Interface adapters
│   ├── handlers/
│   │   ├── auth/                # Authentication handlers
│   │   │   ├── handler.go      # Auth handler interface
│   │   │   ├── login.go        # Login endpoint
│   │   │   └── register.go     # Registration endpoint
│   │   ├── health_handler/      # Health check handler
│   │   │   └── health.go
│   │   └── product_handlers/    # Product HTTP handlers
│   │       ├── create_product.go
│   │       ├── delete_product.go
│   │       ├── get_product.go
│   │       ├── get_product_by_id.go
│   │       ├── handler.go
│   │       └── update_product.go
│   └── routes/                  # Route interfaces
│       └── global_routes.go     # Route registration interface
├── cmd/                         # Application commands
│   ├── router.go                # HTTP route registration
│   └── serve.go                 # Server initialization with DB
├── config/                      # Configuration management
│   ├── config.go                # Configuration structure
│   └── loadenv.go               # Environment loader
├── domain/                      # Business entities & interfaces
│   ├── cart/
│   │   └── cart.go              # Cart entity (model only)
│   ├── health/
│   │   └── health.go            # Health entity
│   ├── order/
│   │   └── order.go             # Order entity (model only)
│   ├── payment/
│   │   └── payment.go           # Payment entity (model only)
│   ├── product/
│   │   └── product.go           # Product & Category entities
│   ├── review/
│   │   └── review.go            # Review entity (model only)
│   └── user/
│       └── user.go              # User entity & interfaces
├── drivers/                     # External drivers & utilities
│   ├── hash.go                  # SHA-256 password hashing
│   └── jwt.go                   # Custom JWT implementation
├── infra/                       # Infrastructure layer
│   ├── memory/                  # In-memory repositories
│   │   └── health_repo.go      # Health repository
│   └── postgresql/              # PostgreSQL integration
│       ├── db.go                # Database connection (singleton)
│       ├── product_repo.go     # Product repository
│       └── user_repo.go        # User repository
├── middlewares/                 # HTTP middlewares
│   ├── cors.go                  # CORS middleware
│   ├── logger.go                # Request logging middleware
│   ├── manager.go               # Middleware chain manager
│   ├── test1.go                 # Test middleware 1
│   └── test2.go                 # Test middleware 2
├── migrations/                  # Database migrations
│   ├── 001_init_schema.down.sql # Initial schema rollback
│   └── 001_init_schema.up.sql   # Initial schema setup
├── usecase/                     # Business logic services
│   ├── auth_service.go         # Authentication service
│   ├── health_service.go       # Health service
│   ├── product_service.go      # Product service
│   └── user_service.go         # User service
├── utils/                       # Utility functions
│   └── error.go                 # Custom error types
├── .env                         # Environment variables (not in repo)
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies lock file
├── main.go                      # Application entry point
└── README.md                    # This file
```

## 🛠️ Tech Stack

- **Language:** Go 1.24.2
- **Architecture:** Clean Architecture / Hexagonal Architecture
- **HTTP Server:** net/http (standard library)
- **Database:** PostgreSQL 12+ 
- **Database Driver:** lib/pq v1.10.9
- **Authentication:** Custom JWT implementation
- **Password Hashing:** SHA-256 (to be upgraded to bcrypt)
- **Configuration:** Environment-based with godotenv v1.5.1
- **Middleware:** Custom middleware chain manager

## 🔄 Development Roadmap

### Phase 1: Foundation ✅ Complete
- [x] Project structure setup with Clean Architecture
- [x] Domain models definition (User, Product, Cart, Order, Payment, Review)
- [x] Full product CRUD operations
- [x] Middleware implementation (CORS, Logger, Chain Manager)
- [x] Health check endpoint
- [x] Route registration interface
- [x] PostgreSQL database integration
- [x] Environment configuration
- [x] Database migrations 

### Phase 2: Authentication & Authorization 🚧 In Progress
- [x] User registration with email/password
- [x] User login endpoint
- [x] Custom JWT implementation
- [x] Password hashing (SHA-256)
- [x] JWT token return in login response
- [x] Protected route middleware
- [x] Role-based access control (RBAC)
- [ ] Password reset functionality
- [ ] Email verification

### Phase 3: Core E-commerce Features 📋 Planned
- [x] Shopping cart functionality
  - [x] Add/remove items
  - [x] Update quantities
  - [x] Cart persistence
- [ ] Order management system
  - [ ] Create orders from cart
  - [ ] Order status tracking
  - [ ] Order history
- [ ] Payment integration
  - [ ] Payment gateway integration
  - [ ] Transaction handling
  - [ ] Invoice generation
- [ ] Product features
  - [x] Categories management
  - [ ] Product search and filtering
  - [x] Inventory tracking
  - [ ] Product images handling
- [ ] Reviews and ratings system

### Phase 4: Quality & Testing 📋 Planned
- [ ] Unit tests for all services
- [ ] Integration tests for handlers
- [ ] Repository tests
- [ ] Middleware tests
- [ ] Input validation and sanitization
- [ ] Error handling improvements
- [ ] Request/Response validation

### Phase 5: Production Ready 📋 Planned
- [ ] API documentation (OpenAPI/Swagger)
- [ ] Docker containerization
- [ ] CI/CD pipeline (GitHub Actions)
- [ ] Performance optimization
  - [ ] Database query optimization
  - [ ] Caching layer (Redis)
  - [ ] Rate limiting
- [ ] Security hardening
  - [ ] Upgrade to bcrypt for passwords
  - [ ] JWT secret from environment
  - [ ] HTTPS support
  - [ ] Security headers
- [ ] Monitoring and logging
  - [ ] Structured logging
  - [ ] Metrics collection
  - [ ] Health monitoring dashboard

## 🔒 Security Considerations

### Current Security Measures
- ✅ Parameterized SQL queries (SQL injection protection)
- ✅ Password hashing before storage
- ✅ Custom error types to avoid exposing internals
- ✅ CORS middleware configured

### Security Improvements Needed
- ⚠️ **Password Hashing:** Using SHA-256, should upgrade to bcrypt
- ⚠️ **CORS:** Currently allows all origins (`*`)
- ⚠️ **Rate Limiting:** Not implemented
- ⚠️ **HTTPS:** No TLS/SSL configuration

## 🧪 Testing

Currently, the project has no test coverage. Testing implementation is planned for Phase 4.

### Planned Test Coverage
- Unit tests for all business logic services
- Integration tests for HTTP handlers
- Repository layer tests
- Middleware tests
- End-to-end API tests

To run tests (once implemented):
```bash
go test ./...
```

## 📊 Database Schema

### Current Tables

#### Users Table
```sql
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Category Table
```sql
CREATE TABLE category (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
```

#### Product Table
```sql
CREATE TABLE product (
    product_id SERIAL PRIMARY KEY,
    category_id INT REFERENCES category(category_id),
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    type VARCHAR(100),
    price FLOAT NOT NULL,
    img_url VARCHAR(255)
);
```
#### Cart Table

```sql
create TABLE cart(
    cart_id SERIAL PRIMARY KEY,
    user_id INT,
    created_at timestamp DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
       ON DELETE CASCADE
);
```
#### Cart_Items table
```sql
create table cart_item(
    cart_item_id SERIAL PRIMARY KEY,
    cart_id INT,
    product_id INT,
    quantity INT,
    created_at timestamp DEFAULT NOW(),
    FOREIGN KEY(cart_id) REFERENCES cart(cart_id)
       ON DELETE CASCADE,
    FOREIGN KEY(product_id) REFERENCES product(product_id)
       ON DELETE CASCADE
);
```
#### Inventory table 
```sql
CREATE TABLE inventory (
    inventory_id SERIAL PRIMARY KEY,
    product_id INT,
    category_id INT,
    quantity INT NOT NULL DEFAULT 0,
    reserved INT NOT NULL DEFAULT 0,
    updated TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (category_id) REFERENCES category(category_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product(product_id) ON DELETE CASCADE
);
```
#### Indexes
```sql
create index users_mail on users(email);
create index category_id on category(category_id);
create index categoryId_productId on product(product_id,category_id);
create index cart_user on cart(cart_id,user_id);
create index cart_items on cart_item(cart_item_id,cart_id);
create index inventory_query_product_id on inventory(product_id,quantity); 
```

```

## 👥 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Development Guidelines
- Follow Go best practices and conventions
- Maintain Clean Architecture principles
- Add appropriate error handling
- Consider adding tests for new functionality

## 📄 License

This project is currently under development. License will be added upon first release.

## 📞 Contact

For questions or support, please open an issue in the GitHub repository.
