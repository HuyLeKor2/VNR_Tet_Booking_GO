package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/huyle/go-ecommerce-backend-api/internal/controller"
	"github.com/huyle/go-ecommerce-backend-api/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before AA--->!")
		c.Next()
		fmt.Println("Alter AA--->!")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before BB--->!")
		c.Next()
		fmt.Println("Alter BB--->!")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before CC--->!")
	c.Next()
	fmt.Println("Alter CC--->!")
}

func NewRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	//use the middleware
	r.Use(middlewares.AuthMiddleware(), BB(), CC)

	// Define a simple GET en=dpoint
	r1 := r.Group("/v1/2024")
	{
		// Nested group for version 1 APIs
		r1.GET("/ping", c.NewPongController().Pong)
		r1.GET("/user/1", c.NewUserController().GetUserByID)
		// 	r1.PUT("/ping", pong)
		// 	r1.PATCH("/ping", pong)
		// 	r1.DELETE("/ping", pong)
		// 	r1.HEAD("/ping", pong)
		// 	r1.OPTIONS("/ping", pong)
		// }

		// r2 := r.Group("/v2/2024")
		// {
		// 	// Nested group for version 1 APIs
		// 	r2.GET("/ping", pong)
		// 	r2.PUT("/ping", pong)
		// 	r2.PATCH("/ping", pong)
		// 	r2.DELETE("/ping", pong)
		// 	r2.HEAD("/ping", pong)
		// 	r2.OPTIONS("/ping", pong)
		// }

		return r
	}
}
