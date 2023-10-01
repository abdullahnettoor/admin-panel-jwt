package handlers

import (
	"fmt"
	"net/http"

	"github.com/abdullahnettoor/admin-panel-jwt/internal/utils"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	// Clear Cache in browser
	utils.ClearCache(c)
	fmt.Println("Entered Home handler")

	// Check if already logged in
	if utils.ContainValidToken(c) {
		fmt.Println("Contains valid token")

		fmt.Println("Checks Admin or User")
		// Check if it is admin
		if c.GetString("role") == "admin" {
			fmt.Println("Role is :", c.GetString("role"))
			c.Redirect(http.StatusFound, "/admin")
			fmt.Println("Going Admin Dashboard")
		} else {
			c.HTML(http.StatusOK, "index.html", c.GetString("username"))
			fmt.Println("Going Home")
		}

	}
	c.Set("msg", "Login to see home")
	c.Redirect(http.StatusSeeOther, "/login")
	fmt.Println("Redircting to Login")
}
