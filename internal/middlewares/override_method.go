package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
This middleware is not used in this project yet.
I was trying out different things and something goes wrong.
I hope this code can be useful in future projects.
So, I'm keeping it in this project for future references.
*/

func OverrideMethod(c *gin.Context) {

	fmt.Println("⚠️ Initial Method is", c.Request.Method)

	if c.Request.Method == "POST" {

		method := c.Request.FormValue("_method")
		if method == "" {
			c.Next()
		}

		if method == "PUT" || method == "PATCH" || method == "DELETE" {
			c.Request.Method = method
		}

	}

	fmt.Println("⚠️ After Change: Method is", c.Request.Method)
	// Call the next handler
	c.Next()
}
