package main

import (
    "github.com/gin-gonic/gin"
    "easyBackend/controller"
    "easyBackend/middleware"
)

func main() {
    r := gin.Default()
    r.Use(middleware.CORSConfig())

    r.POST("/login", controller.Login)
	r.POST("/register", controller.RegisterHandler)

    protected := r.Group("/api", middleware.JWTMiddleware())
	{
		protected.GET("/users", controller.GetUsers)
	    protected.GET("/users/:id", controller.GetUsersByID)
	}
    r.Run(":7777") 
}
