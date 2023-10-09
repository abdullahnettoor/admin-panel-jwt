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

	// Check if the request we are getting from html is POST
	if c.Request.Method == "POST" {

		// Retreive data from the input filed in html.
		// For Example: <input type="hidden" name="_method" value="PUT" />
		method := c.Request.FormValue("_method")
		if method == "" {
			c.Next()
		}

		// if the value is other than POST convert the request method in to Specified method
		if method == "PUT" || method == "PATCH" || method == "DELETE" {
			c.Request.Method = method
		}

	}

	fmt.Println("⚠️ After Change: Method is", c.Request.Method)
	// Call the next handler
	c.Next()
}
