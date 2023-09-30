package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/abdullahnettoor/admin-panel-jwt/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secretKey = []byte(os.Getenv("KEY"))

type CustomClaims struct {
	UserID uuid.UUID
	Name   string
	jwt.RegisteredClaims
}

func CreateToken(c *gin.Context, u models.User) {

	// Create the Custom Claims
	claims := &CustomClaims{
		u.ID,
		u.Name,
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
	c.Set("username", u.Name)
	fmt.Println("Cookie Created")
}

// Validate Token
func ContainValidToken(c *gin.Context) bool {

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Println("Error occured while validating Cookie:", err)
		return false
	}
	if tokenString == "" {
		fmt.Println("Cookie not found")
		return false
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil || !token.Valid {
		fmt.Println("Error occured whilr fetching token")
		return false
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

		if claims.ExpiresAt.Before(time.Now()) {
			fmt.Println("Session Expired")
			return false
		}

		username := claims.Name

		fmt.Println("From JWT", username)
		c.Set("username", username)
		return true

	} else {
		fmt.Println("Error occured while parsing token:", err)
		return false
	}
}

func DeleteCookie(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", false, true)
	fmt.Println("Cookie Deleted")
}
