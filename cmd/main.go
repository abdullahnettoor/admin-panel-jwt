package main

import (
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

	// Create Server
	r := gin.Default()

	// Load templates and static files
	r.LoadHTMLGlob("view/*.html")
	r.Static("/static", "./static")

	// Handle User actions
	r.GET("/", handlers.Home)
	r.GET("/login", handlers.Login)
	r.POST("/login", handlers.LoginPost)
	r.GET("/signup", handlers.Signup)
	r.POST("/signup", handlers.SignupPost)
	r.GET("/logout", handlers.Logout)

	// Handle Admin actions
	r.GET("/admin", handlers.AdminDashboard)
	r.GET("/admin/create/user", handlers.LoadCreateUser)
	r.POST("/admin/create/user", handlers.CreateUser)
	r.POST("/admin/update", handlers.LoadUpdateUser)
	r.POST("/admin/update/user", handlers.UpdateUser)
	r.POST("/admin/delete/user", handlers.DeleteUser)

	// Run Server
	r.Run(":8080")
}
