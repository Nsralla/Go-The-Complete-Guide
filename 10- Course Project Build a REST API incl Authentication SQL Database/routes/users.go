package routes

import (
	"fmt"
	"net/http"

	"example.com/project/models"
	"example.com/project/utils"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var newUser models.User
	err := context.ShouldBind(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//Save the user to DB
	newUser, err = newUser.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": models.User{ID: newUser.ID, Email: newUser.Email}})

}

func login(context *gin.Context) {
	// Read data sent by user
	var user models.User
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	user, err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}
	fmt.Println("user after validating credentials", user)
	// All good?  Generate JWT token
	token, err := utils.GenerateJWT(user.Email, int(user.ID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func getAllUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, users)
}
