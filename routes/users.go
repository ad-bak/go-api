package routes

import (
	"net/http"

	"example.com/books/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	var err error

	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	if err = user.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})

}
