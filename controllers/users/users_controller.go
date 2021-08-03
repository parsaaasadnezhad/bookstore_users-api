package users

import (
	"github.com/gin-gonic/gin"
	"main/domain/users"
	"main/services"
	"main/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId , userErr := strconv.ParseInt(c.Param("user_id") , 10 , 64)
	if userErr != nil{
		err := errors.NewBadRequestError("user id should be number")
		c.JSON(err.Status , err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status , getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status , restError)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status , saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
