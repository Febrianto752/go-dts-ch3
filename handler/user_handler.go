package handler

import (
	"net/http"

	"github.com/Febrianto752/go-dts-ch3/entity"
	"github.com/Febrianto752/go-dts-ch3/helper"
	"github.com/Febrianto752/go-dts-ch3/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UserRegisterHandler(ctx *gin.Context)
	UserLoginHandler(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func (h *userHandler) UserLoginHandler(ctx *gin.Context) {
	var payload entity.UserLogin

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {

		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
	}

	loggedInUser, err := h.userService.Login(payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	token := helper.GenerateToken(loggedInUser.Id, loggedInUser.Email, loggedInUser.Role)
	helper.SuccessResponse(ctx, http.StatusOK, gin.H{
		"access_token": token,
	})
}

func (h *userHandler) UserRegisterHandler(ctx *gin.Context) {
	var payload entity.UserRequest

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := h.userService.Register(payload)
	if err != nil {

		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, gin.H{
		"id":        newUser.Id,
		"email":     newUser.Email,
		"full_name": newUser.FullName,
	})
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService: userService}
}
