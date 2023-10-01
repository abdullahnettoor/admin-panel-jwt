package middlewares

import "github.com/gin-gonic/gin"

func OverrideMethod(c *gin.Context) {

	if c.Request.Method == "POST" {

		method := c.Request.FormValue("_method")
		if method == "" {
			c.Next()
		}

		if method == "PUT" || method == "PATCH" || method == "DELETE" {
			c.Request.Method = method
		}

	}

	// Call the next handler
	c.Next()
}
