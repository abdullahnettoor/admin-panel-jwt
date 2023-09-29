package initializers

import (
	"fmt"

	"github.com/abdullahnettoor/admin-panel-jwt/internal/models"
)

func SyncDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Error syncing database -->", err)
	}
}
