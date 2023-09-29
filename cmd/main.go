package main

import (
	"fmt"

	"github.com/abdullahnettoor/admin-panel-jwt/internal/handlers"
	"github.com/abdullahnettoor/admin-panel-jwt/internal/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	fmt.Println("Hello World!")

	r := gin.Default()
	r.LoadHTMLGlob("view/*.html")
	r.Static("/static", "./static")

	r.GET("/", handlers.Home)
	r.GET("/login", handlers.Login)
	r.GET("/signup", handlers.Signup)
	r.POST("/signup", handlers.SignupPost)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.Run()
}
