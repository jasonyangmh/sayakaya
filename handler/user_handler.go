package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jasonyangmh/sayakaya/dto"
	"github.com/jasonyangmh/sayakaya/model"
	"github.com/jasonyangmh/sayakaya/shared"
	"github.com/jasonyangmh/sayakaya/usecase"
	"gorm.io/gorm"
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
		ID:         user.Model.ID,
		Email:      user.Email,
		Birthday:   user.Birthday.Format(shared.TimeLayout),
		IsVerified: user.IsVerified,
		IsBirthday: isBirthday,
	}

	c.JSON(http.StatusCreated, dto.JSONResponse{Data: res})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.userUsecase.FindUsers(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	res := []dto.UserResponse{}

	for _, user := range users {
		isBirthday := h.userUsecase.CheckUserBirthday(ctx, &user)

		res = append(res, dto.UserResponse{
			ID:         user.Model.ID,
			Email:      user.Email,
			Birthday:   user.Birthday.Format(shared.TimeLayout),
			IsVerified: user.IsVerified,
			IsBirthday: isBirthday,
		})
	}

	c.JSON(http.StatusCreated, dto.JSONResponse{Data: res})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	str, match := c.Params.Get(shared.ID)
	id, err := strconv.Atoi(str)

	if !match || err != nil {
		c.Error(shared.ErrInvalidID)
		return
	}

	user := &model.User{
		Model: gorm.Model{ID: uint(id)},
	}

	user, err = h.userUsecase.FindUserByID(ctx, user)
	if err != nil {
		c.Error(err)
		return
	}

	isBirthday := h.userUsecase.CheckUserBirthday(ctx, user)

	res := dto.UserResponse{
		ID:         user.Model.ID,
		Email:      user.Email,
		Birthday:   user.Birthday.Format(shared.TimeLayout),
		IsVerified: user.IsVerified,
		IsBirthday: isBirthday,
	}

	c.JSON(http.StatusOK, dto.JSONResponse{Data: res})
}

func (h *UserHandler) PutUser(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.UserUpdateRequest{}

	str, match := c.Params.Get(shared.ID)
	id, err := strconv.Atoi(str)

	if !match || err != nil {
		c.Error(shared.ErrInvalidID)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	user := &model.User{
		Model:      gorm.Model{ID: uint(id)},
		IsVerified: req.IsVerified,
	}

	user, err = h.userUsecase.UpdateUser(ctx, user)
	if err != nil {
		c.Error(err)
		return
	}

	user, err = h.userUsecase.FindUserByID(ctx, user)
	if err != nil {
		c.Error(err)
		return
	}

	isBirthday := h.userUsecase.CheckUserBirthday(ctx, user)

	res := dto.UserResponse{
		ID:         user.Model.ID,
		Email:      user.Email,
		Birthday:   user.Birthday.Format(shared.TimeLayout),
		IsVerified: user.IsVerified,
		IsBirthday: isBirthday,
	}

	c.JSON(http.StatusOK, dto.JSONResponse{Data: res})
}
