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

func Login(c *gin.Context) {
	// Clear Cache of browser
	utils.ClearCache(c)

	// Checks if User already logged in
	if utils.ContainValidToken(c) {

		if c.GetString("role") == "admin" {
			c.Redirect(http.StatusFound, "/admin")
			return
		}

		c.Redirect(http.StatusFound, "/")
		return
	}

	// Load Login
	msg := c.GetString("msg")
	fmt.Println("Msg is:", msg)
	c.HTML(http.StatusOK, "login.html", msg)
}

func LoginPost(c *gin.Context) {
	// Clear Cache of browser
	utils.ClearCache(c)

	// Check if user already logged in
	if utils.ContainValidToken(c) {
		c.Redirect(http.StatusFound, "/")
		return
	}

	// Collect form data
	user := models.User{
		Email:    strings.TrimSpace(c.Request.FormValue("email")),
		Password: strings.TrimSpace(c.Request.Form.Get("password")),
	}

	// Validate user
	var dbUser models.User
	result := initializers.DB.Raw(`SELECT * FROM users WHERE email = ?`, user.Email).Scan(&dbUser)
	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "login.html", "Something went wrong")
		return
	}
	if result.RowsAffected != 1 {
		c.HTML(http.StatusBadRequest, "login.html", "User doesn't exist")
		return
	}

	// Check Password
	if !utils.ComparePassword(user.Password, dbUser.Password) {
		c.HTML(http.StatusBadRequest, "login.html", "Password is wrong")
		return
	}

	// Generate Token and Create Cookie
	utils.CreateToken(c, dbUser)
	uName := c.GetString("username")
	fmt.Println(uName)

	if dbUser.IsAdmin {
		c.Redirect(http.StatusSeeOther, "/admin")
		return
	}

	// Load home
	c.HTML(http.StatusOK, "index.html", uName)
}
