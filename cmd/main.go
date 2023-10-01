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
	r.GET("/admin/create/user", handlers.LoadCreateUser)
	r.POST("/admin/create/user", handlers.CreateUser)
	r.GET("/admin/update/user/:id", handlers.LoadUpdateUser)
	r.POST("/admin/update/user/:id", handlers.UpdateUser)
	r.POST("/admin/delete/user/:id", handlers.DeleteUser)

	r.Run()
}
