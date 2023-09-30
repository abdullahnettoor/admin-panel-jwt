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

	// Check if already logged in
	if utils.ContainValidToken(c) {
		if uName, ok := c.Get("username"); ok {
			fmt.Println(uName)
			c.HTML(http.StatusOK, "index.html", uName)
		}
	} else {
		c.Set("message", "Login to see home")
		c.Redirect(http.StatusSeeOther, "/login")
	}

}
