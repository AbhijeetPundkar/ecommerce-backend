package routes

import (
    "github.com/AbhijeetPundkar/ecommerce-backend/controllers"
    "github.com/AbhijeetPundkar/ecommerce-backend/middlewares"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
    router.POST("/signup", controllers.SignUp)
    router.POST("/login", controllers.Login)

    auth := router.Group("/")
    auth.Use(middlewares.JWTAuthMiddleware())
    {
        auth.GET("/protected", controllers.Protected)
    }
}