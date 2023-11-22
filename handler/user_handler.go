package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jasonyangmh/sayakaya/dto"
	"github.com/jasonyangmh/sayakaya/model"
	"github.com/jasonyangmh/sayakaya/shared"
	"github.com/jasonyangmh/sayakaya/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: uu,
	}
}

func (h *UserHandler) PostUser(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.UserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	birthday, err := time.Parse(shared.TimeLayout, req.Birthday)
	if err != nil {
		c.Error(shared.ErrInvalidDateLayout)
		return
	}

	user := &model.User{
		Email:    req.Email,
		Birthday: birthday,
	}

	user, err = h.userUsecase.CreateUser(ctx, user)
	if err != nil {
		c.Error(err)
		return
	}

	isBirthday := h.userUsecase.CheckUserBirthday(ctx, user)

	res := dto.UserResponse{
		Email:      user.Email,
		Birthday:   user.Birthday.Format(shared.TimeLayout),
		IsVerified: user.IsVerified,
		IsBirthday: isBirthday,
	}

	c.JSON(http.StatusCreated, dto.JSONResponse{Data: res})
}
