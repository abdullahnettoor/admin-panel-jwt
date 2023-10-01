package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/abdullahnettoor/admin-panel-jwt/internal/initializers"
	"github.com/abdullahnettoor/admin-panel-jwt/internal/models"
	"github.com/abdullahnettoor/admin-panel-jwt/internal/utils"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	// Clear Cache
	utils.ClearCache(c)

	// Check if user already logged in
	if utils.ContainValidToken(c) {

		// Check if it is admin
		if c.GetString("role") == "admin" {
			c.Redirect(http.StatusFound, "/admin")
			return
		}

		c.Redirect(http.StatusFound, "/")
		return
	}

	// Load signup page
	c.HTML(http.StatusOK, "signup.html", nil)
}

func SignupPost(c *gin.Context) {
	// Clear Cache
	utils.ClearCache(c)

	// Check if user already logged in
	if utils.ContainValidToken(c) {

		// Check if it is admin
		if c.GetString("role") == "admin" {
			c.Redirect(http.StatusFound, "/admin")
			return
		}

		c.Redirect(http.StatusFound, "/")
		return
	}

	// Recieve values from form
	newUser := models.User{
		Name:     strings.TrimSpace(c.Request.FormValue("name")),
		Email:    strings.TrimSpace(c.Request.FormValue("email")),
		Password: strings.TrimSpace(c.Request.FormValue("password")),
	}
	confirmPassword := c.Request.FormValue("confirm-password")

	// Check if user exists
	var users []models.User
	result := initializers.DB.Raw("SELECT email FROM users WHERE email = ?", newUser.Email).Scan(&users)
	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "signup.html", "Error occured while creating user. Try again")
		fmt.Println(result.Error)
		return
	}
	if result.RowsAffected > 0 {
		c.HTML(http.StatusConflict, "signup.html", "User already exists")
		return
	}

	// Check passwords match
	if newUser.Password != confirmPassword {
		c.HTML(http.StatusNotAcceptable, "signup.html", "Password must match")
		return
	}

	// Hash password
	newUser.Password = utils.HashPassword(newUser.Password)

	// Create user
	result = initializers.DB.Create(&newUser)
	if result.Error != nil {
		c.HTML(http.StatusNotAcceptable, "signup.html", "Something went wrong. Try again")
		return
	}

	fmt.Println("User created successfully with id:", newUser.ID)
	utils.CreateToken(c, newUser)

	// Redirect to Home page
	c.HTML(http.StatusOK, "index.html", newUser.Name)
}
