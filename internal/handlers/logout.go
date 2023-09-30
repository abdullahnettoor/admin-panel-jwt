package handlers

import (
	"net/http"

	"github.com/abdullahnettoor/admin-panel-jwt/internal/utils"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	// Clear Cache
	utils.ClearCache(c)

	// Delete existing cookie
	utils.DeleteCookie(c)

	c.Set("message", "Logged out succesfully")
	c.Redirect(http.StatusSeeOther, "/login")
}
