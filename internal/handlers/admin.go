package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/abdullahnettoor/admin-panel-jwt/internal/initializers"
	"github.com/abdullahnettoor/admin-panel-jwt/internal/models"
	"github.com/abdullahnettoor/admin-panel-jwt/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var dataMap = make(map[string]any)

func AdminDashboard(c *gin.Context) {
	// Clear cache from the browser
	utils.ClearCache(c)

	// Checks if Token is Valid
	if utils.ContainValidToken(c) {

		fmt.Println("Is Admin", (c.GetString("role") == "admin"))

		// Check if it is admin
		if c.GetString("role") != "admin" {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}

		// Get Users list
		GetUsers(c)
		return
	}

	// Redirect to login because valid token not found
	c.HTML(http.StatusOK, "login.html", nil)
}

func GetUsers(c *gin.Context) {
	var users []models.User

	result := initializers.DB.Order("created_at DESC").Where("is_admin = false").Find(&users)
	if result.Error != nil {
		fmt.Println("Error fetching users:", result.Error)
		c.HTML(http.StatusOK, "admin-panel.html", users)
		return
	}
	dataMap["UsersList"] = users
	dataMap["Admin"] = c.GetString("username")

	c.HTML(http.StatusOK, "admin-panel.html", dataMap)
}

func LoadCreateUser(c *gin.Context) {
	// Clear cache from the browser
	utils.ClearCache(c)

	// Checks if Token is Valid
	if utils.ContainValidToken(c) {

		fmt.Println("Is Admin", (c.GetString("role") == "admin"))

		// Check if it is admin
		if c.GetString("role") != "admin" {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
	}

	// Get map data
	dataMap["Admin"] = c.GetString("username")
	dataMap["Message"] = ""

	c.HTML(http.StatusOK, "create-user.html", dataMap)
}

func CreateUser(c *gin.Context) {
	// Clear cache from the browser
	utils.ClearCache(c)

	// Checks if Token is Valid
	if utils.ContainValidToken(c) {

		fmt.Println("Is Admin", (c.GetString("role") == "admin"))

		// Check if it is admin
		if c.GetString("role") != "admin" {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
	}

	// Retrieve form data
	user := models.User{
		Name:     strings.TrimSpace(c.Request.FormValue("name")),
		Email:    strings.TrimSpace(c.Request.FormValue("email")),
		Password: strings.TrimSpace(c.Request.FormValue("password")),
	}

	// Check user with email exists
	result := initializers.DB.Exec("SELECT email FROM users WHERE email = ?", user.Email)
	if result.Error != nil {
		dataMap["Message"] = "Error creating user"
		c.HTML(http.StatusOK, "create-user.html", dataMap)
		return
	}
	if result.RowsAffected > 0 {
		dataMap["Message"] = "User with email already exist"
		c.HTML(http.StatusOK, "create-user.html", dataMap)
		return
	}

	// Hash Password
	user.Password = utils.HashPassword(user.Password)

	// Create New user
	result = initializers.DB.Create(&user)
	if result.Error != nil {
		dataMap["Message"] = "Error creating user"
		c.HTML(http.StatusOK, "create-user.html", dataMap)
		return
	}

	// Pass Success msg
	dataMap["Message"] = "User created succesfully"
	c.HTML(http.StatusOK, "create-user.html", dataMap)
}

func LoadUpdateUser(c *gin.Context) {
	// Clear cache from the browser
	utils.ClearCache(c)

	// Checks if Token is Valid
	if utils.ContainValidToken(c) {

		fmt.Println("Is Admin", (c.GetString("role") == "admin"))

		// Check if it is admin
		if c.GetString("role") != "admin" {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
	}

	// Get map datas
	dataMap["Admin"] = c.GetString("username")
	dataMap["Message"] = ""

	// Retrieve id from url
	id := c.Request.FormValue("id")

	// Fetch user from the db
	var user models.User
	res := initializers.DB.Raw(`SELECT * FROM users WHERE id = ?`, id).Scan(&user)
	if res.Error != nil {
		dataMap["Message"] = "Error fetching user"
		fmt.Println("Error is ", res.Error)
		c.HTML(http.StatusOK, "update-user.html", dataMap)
		return
	}

	// Pass the data to show in update page
	dataMap["User"] = user
	c.HTML(http.StatusOK, "update-user.html", dataMap)
}

func UpdateUser(c *gin.Context) {

	// Checks if Token is Valid
	if utils.ContainValidToken(c) {

		fmt.Println("Is Admin", (c.GetString("role") == "admin"))

		// Check if it is admin
		if c.GetString("role") != "admin" {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
	}

	// Get map datas
	dataMap["Admin"] = c.GetString("username")
	dataMap["Message"] = ""

	// Get id from request
	id := c.Request.FormValue("id")

	// Retrieve form data from form
	user := models.User{
		Name:  strings.TrimSpace(c.Request.FormValue("name")),
		Email: strings.TrimSpace(c.Request.FormValue("email")),
		Model: gorm.Model{UpdatedAt: time.Now()},
	}

	// Check user with email exists
	var currentUser string
	result := initializers.DB.Raw("SELECT email FROM users WHERE email = ?", user.Email).Scan(&currentUser)
	if result.Error != nil {
		dataMap["Message"] = "Error updating user"
		c.HTML(http.StatusOK, "update-user.html", dataMap)
		return
	}
	if result.RowsAffected > 0 && user.Email != currentUser {
		dataMap["Message"] = "Already user exist with this email"
		c.HTML(http.StatusOK, "update-user.html", dataMap)
		return
	}

	// Update user
	result = initializers.DB.Exec(`UPDATE users SET email = ?, name = ? WHERE id = ?`, user.Email, user.Name, id)
	if result.Error != nil {
		dataMap["Message"] = "Error updating user"
		c.HTML(http.StatusBadRequest, "update-user.html", dataMap)
		return
	}

	// Pass success Message
	dataMap["Message"] = "User updated succesfully"
	c.Redirect(http.StatusSeeOther, "/admin")

}

func DeleteUser(c *gin.Context) {
	// Clear cache from the browser
	utils.ClearCache(c)

	// Checks if Token is Valid
	if utils.ContainValidToken(c) {

		fmt.Println("Is Admin", (c.GetString("role") == "admin"))

		// Check if it is admin
		if c.GetString("role") != "admin" {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
	}

	// Get map datas
	dataMap["Admin"] = c.GetString("username")
	dataMap["Message"] = ""

	// Get id from request
	id := c.Request.FormValue("id")

	// Deleting user with the id
	result := initializers.DB.Exec(`DELETE FROM users WHERE id = ?`, id)
	if result.Error != nil {
		fmt.Println("Error while deleting user", result.Error)
		return
	}

	// Redirecting to admin panel
	c.Redirect(http.StatusSeeOther, "/admin")
}
