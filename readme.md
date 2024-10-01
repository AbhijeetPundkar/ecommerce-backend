# E-commerce Backend

This is the backend for our e-commerce application.

## Project Structure

```
/ecom-backend
│
├── /cmd                # Main application entry point (e.g., main.go)
│   └── main.go         # Starts the server
│
├── /config             # Configuration files
│   └── config.go       # Load environment variables and configurations
│
├── /controllers        # Handlers for routes
│   └── product.go      # Product controller (CRUD operations)
│   └── user.go         # User controller (registration, login)
│   └── order.go        # Order controller (order processing)
│
├── /routes             # API route definitions
│   └── routes.go       # Registers all the routes
│
├── /middlewares        # Middleware (JWT authentication, etc.)
│   └── auth.go         # JWT middleware
│
├── /models             # Database models
│   └── product.go      # Product model
│   └── user.go         # User model
│   └── order.go        # Order model
│
├── /pkg                # Reusable code (utility functions)
│   └── utils.go        # Helper functions (e.g., password hashing, JWT generation)
│
├── /internal           # Core application logic
│   └── db              # Database connection setup
│       └── connection.go
│   └── services        # Business logic services (interact with DB models)
│       └── product_service.go
│       └── user_service.go
│       └── order_service.go
│
├── /migrations         # Database migrations
│   └── initial_schema.sql  # SQL scripts for setting up database schema
│
├── /tests              # Unit and integration tests
│   └── product_test.go  # Tests for product controller
│   └── user_test.go     # Tests for user controller
│
├── .env                # Environment variables
├── .gitignore          # Ignored files (e.g., binaries, build artifacts)
├── go.mod              # Go modules file
└── go.sum              # Go modules dependencies
```

## Getting Started

## API Documentation

## Contributing

