package handlers

import (
	"net/http"
	"strings"

	"github.com/abdullahnettoor/admin-panel-jwt/internal/initializers"
	"github.com/abdullahnettoor/admin-panel-jwt/internal/models"
	"github.com/abdullahnettoor/admin-panel-jwt/internal/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginPost(c *gin.Context) {
	// Clear Cache of browser
	utils.ClearCache(c)

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

	c.HTML(http.StatusOK, "index.html", dbUser)
}
