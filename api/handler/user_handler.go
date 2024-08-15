package handler

import (
	"log/slog"
	"net/http"
	"todo-app/api/request"
	"todo-app/application/usecase"

	"github.com/gin-gonic/gin"
)

func NewUserHandler(usecase usecase.UserRecord) *UserHandler {
	return &UserHandler{usecase: usecase}
}

type UserHandler struct {
	usecase usecase.UserRecord
}

func (u UserHandler) GetUsers(c *gin.Context) {
	users, err := u.usecase.FindAll()
	if err != nil {
		slog.ErrorContext(c, err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u UserHandler) GetUser(c *gin.Context) {
	userId := c.Param("id")
	user, err := u.usecase.FindUser(userId)
	if err != nil {
		slog.ErrorContext(c, err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u UserHandler) PostUser(c *gin.Context) {
	var body request.PostUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := body.Validate(); err != nil {
		slog.ErrorContext(c, err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := u.usecase.CreateUser(body.ToDto())
	if err != nil {
		slog.ErrorContext(c, err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (u UserHandler) PutUser(c *gin.Context) {
	userId := c.Param("id")
	var body request.PutUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := body.Validate(); err != nil {
		slog.ErrorContext(c, err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	dto := body.ToDto()
	dto.Id = userId
	err := u.usecase.UpdateUser(dto)
	if err != nil {
		slog.ErrorContext(c, err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (u UserHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	err := u.usecase.DeleteUser(userId)
	if err != nil {
		slog.ErrorContext(c, err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
