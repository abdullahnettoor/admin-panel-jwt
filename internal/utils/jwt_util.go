package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/abdullahnettoor/admin-panel-jwt/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("KEY"))

// Create struct for custom claims
type CustomClaims struct {
	Email   string
	Name    string
	IsAdmin bool
	jwt.RegisteredClaims
}

func CreateToken(c *gin.Context, u models.User) {

	// Create the Custom Claims
	claims := &CustomClaims{
		u.Email,
		u.Name,
		u.IsAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			Issuer:    "iStore",
		},
	}

	// Generate token based on claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Retrieving token string
	ss, err := token.SignedString(secretKey)
	fmt.Printf("%v %v", ss, err)
	if err != nil {
		fmt.Println("Error occured while creating token:", err)
		return
	}

	// Set cookie from token
	c.SetCookie("Authorization", ss, 3600, "", "", false, true)

	// Set user values to gin context
	c.Set("username", claims.Name)
	c.Set("role", claims.IsAdmin)
}

// Validate Token
func ContainValidToken(c *gin.Context) bool {

	// Check if cookie is available and retrieve the token string
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Println("Error occured while validating Cookie:", err)
		return false
	}
	if tokenString == "" {
		fmt.Println("Cookie not found")
		return false
	}

	// Parse jwt token with custom claims
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}, jwt.WithLeeway(5*time.Second))

	// Check if token is valid
	if err != nil || !token.Valid {
		fmt.Println("Error occured whilr fetching token")
		return false
	}

	// Assign parsed data from token to calims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

		// Check if token is expired
		if claims.ExpiresAt.Before(time.Now()) {
			fmt.Println("Session Expired")
			return false
		}

		// Retrieve the role of user from claim
		var role string
		if claims.IsAdmin {
			role = "admin"
		} else {
			role = "user"
		}

		// Set user values to gin context
		c.Set("userEmail", claims.Email)
		c.Set("username", claims.Name)
		c.Set("role", role)

		return true

	} else {
		fmt.Println("Error occured while parsing token:", err)
		return false
	}
}

// Delete cookie
func DeleteCookie(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", false, true)
	fmt.Println("Cookie Deleted")
}
