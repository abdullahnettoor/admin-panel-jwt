package handlers

import (
	"fmt"
	"net/http"

	"github.com/abdullahnettoor/admin-panel-jwt/internal/initializers"
	"github.com/abdullahnettoor/admin-panel-jwt/internal/models"
	"github.com/abdullahnettoor/admin-panel-jwt/internal/utils"
	"github.com/gin-gonic/gin"
)

var dataMap = make(map[string]any)

func AdminDashboard(c *gin.Context) {
	utils.ClearCache(c)

	if utils.ContainValidToken(c) {

		fmt.Println("Is Admin", (c.GetString("role") == "admin"))

		// Check if it is admin
		if c.GetString("role") != "admin" {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
		GetUsers(c)
		return
	}
	c.HTML(http.StatusOK, "login.html", nil)
}

func GetUsers(c *gin.Context) {
	var users []models.User

	result := initializers.DB.Find(&users)
	if result.Error != nil {
		fmt.Println("Error fetching users:", result.Error)
		c.HTML(http.StatusOK, "admin-panel.html", users)
		return
	}
	dataMap["UsersList"] = users
	dataMap["Admin"] = c.GetString("username")

	c.HTML(http.StatusOK, "admin-panel.html", dataMap)
}

func CreateUser(c *gin.Context) {

}
func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {
	// initializers.DB.Exec()

}
