package routes

import (
	"net/http"
	"example.com/project/models"
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
