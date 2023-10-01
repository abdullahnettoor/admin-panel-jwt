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
	r.POST("/login", handlers.LoginPost)
	r.GET("/signup", handlers.Signup)
	r.POST("/signup", handlers.SignupPost)
	r.GET("/logout", handlers.Logout)

	r.GET("/admin", handlers.AdminDashboard)
	r.POST("/admin/create/user", handlers.CreateUser)
	r.PUT("/admin/update/user:id", handlers.UpdateUser)
	r.DELETE("/admin/delete/user:id", handlers.DeleteUser)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.Run()
}
