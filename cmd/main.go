package main

import (
    "github.com/AbhijeetPundkar/ecommerce-backend/internals/db"
    "github.com/AbhijeetPundkar/ecommerce-backend/routes"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    // Initialize the database connection
    db.Connect()

    // Initialize the Gin router
    router := gin.Default()

    // Register routes
    routes.RegisterRoutes(router)

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}