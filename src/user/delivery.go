package user

import (
	"errors"
	"go-project-management-book/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (dlv *UserHandler) RegisterUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		helper.SendResponseError(c, helper.ErrBadRequest)
		return
	}

	if err := dlv.service.RegisterUser(&user); err != nil {
		resp := helper.ErrorStruct{
			Code:        helper.RCGeneralError,
			HTTPCode:    http.StatusInternalServerError,
			Description: helper.DescErrorGeneral,
			Message:     err.Error(),
		}
		helper.SendResponseError(c, resp)
		return
	}
	resp := helper.Response{
		ResponseCode: helper.RCSuccess,
		Description:  helper.DescriptionSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data: map[string]interface{}{
			"username": user.Username,
		},
	}

	helper.SendResponseSuccess(c, resp)
}

func (dlv *UserHandler) LoginUser(ctx *gin.Context) {
	var req ReqUserLogin

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendResponseError(ctx, helper.ErrBadRequest)
		return
	}

	user, err := dlv.service.LoginUser(req.Username, req.Password)
	if err != nil {
		helper.SendResponseError(ctx, helper.ErrLoginFailed)
		return
	}

	resp := helper.Response{
		ResponseCode: helper.RCSuccess,
		Description:  helper.DescriptionSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         user,
	}

	helper.SendResponseSuccess(ctx, resp)
}

func (dlv *UserHandler) GetProfile(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := dlv.service.GetProfile(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.SendResponseError(ctx, helper.ErrNotFound)
			return
		}

		resp := helper.ErrorStruct{
			Code:        helper.RCGeneralError,
			HTTPCode:    http.StatusInternalServerError,
			Description: helper.DescErrorGeneral,
			Message:     err.Error(),
		}
		helper.SendResponseError(ctx, resp)
		return
	}
	resp := helper.Response{
		ResponseCode: helper.RCSuccess,
		Description:  helper.DescriptionSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         user,
	}
	helper.SendResponseSuccess(ctx, resp)
}
