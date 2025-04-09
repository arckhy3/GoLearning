package routes

import (
	"net/http"
	"strconv"

	"example.com/event/models"
	"github.com/gin-gonic/gin"
)

func getUsers(context *gin.Context) {
	users, err := models.GetAllUser()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}

func getUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	user, err := models.FindUserByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}

func saveUser(context *gin.Context) {
	var user = models.User{}
	err := context.ShouldBind(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	user, err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, user)
}

func updateUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	var updatedUser = models.Event{}
	err = context.ShouldBind(&updatedUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	updatedUser.ID = id

	err = updatedUser.UpdateByID()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User update succes", "user": updatedUser})
}

func deleteUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	user, err := models.FindUserByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user", "error": err.Error()})
		return
	}

	err = user.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User delete success"})

}
