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

## Environment Variables

The application uses the following environment variables, which should be defined in the `.env` file:

- `DB_HOST`: The hostname of the database server.
- `DB_PORT`: The port number on which the database server is running.
- `DB_USER`: The username for the database connection.
- `DB_PASSWORD`: The password for the database connection.
- `DB_NAME`: The name of the database to connect to.
- `JWT_SECRET`: The secret key used for signing JWT tokens. --> add to env
- `PORT`: The port number on which the server will listen.

Ensure that you have a `.env` file in the root directory with the appropriate values for these variables.

## TODO List

- [x] Set up the initial project structure
- [x] Implement user authentication and authorization
- [ ] Implement GoRoutine to refresh the token in background
- [ ] Create CRUD operations for products




