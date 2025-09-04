package middleware

import (
	"fmt"
	"net/http"

	"example.com/project/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
		context.Abort() // This stops the request from proceeding further
		return
	}

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"}) // or like this
		return
	}

	userID := int64(claims.ID)
	fmt.Println("FROM authintication function, Authenticated user ID:", userID)
	context.Set("userID", userID) // Store user ID in context
	context.Next() // Call the next handler in the chain
}
