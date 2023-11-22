package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jasonyangmh/sayakaya/dto"
	"github.com/jasonyangmh/sayakaya/model"
	"github.com/jasonyangmh/sayakaya/shared"
	"github.com/jasonyangmh/sayakaya/usecase"
	"gorm.io/gorm"
)

type PromoHandler struct {
	promoUsecase usecase.PromoUsecase
}

func NewPromoHandler(pu usecase.PromoUsecase) *PromoHandler {
	return &PromoHandler{
		promoUsecase: pu,
	}
}

func (h *PromoHandler) PostPromo(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.PromoRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	startDate, err := time.Parse(shared.TimeLayout, req.StartDate)
	if err != nil {
		c.Error(shared.ErrInvalidDateLayout)
		return
	}

	endDate, err := time.Parse(shared.TimeLayout, req.StartDate)
	if err != nil {
		c.Error(shared.ErrInvalidDateLayout)
		return
	}

	promo := &model.Promo{
		Name:      req.Name,
		StartDate: startDate,
		EndDate:   endDate,
		Amount:    req.Amount,
	}

	promo, err = h.promoUsecase.CreatePromo(ctx, promo)
	if err != nil {
		c.Error(err)
		return
	}

	isEnded := h.promoUsecase.CheckPromoEndDate(ctx, promo)

	res := dto.PromoResponse{
		ID:        promo.Model.ID,
		Name:      promo.Name,
		StartDate: promo.StartDate.Format(shared.TimeLayout),
		EndDate:   promo.EndDate.Format(shared.TimeLayout),
		Amount:    promo.Amount,
		IsEnded:   isEnded,
	}

	c.JSON(http.StatusCreated, dto.JSONResponse{Data: res})
}

func (h *PromoHandler) GetPromos(c *gin.Context) {
	ctx := c.Request.Context()

	promos, err := h.promoUsecase.FindPromos(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	res := []dto.PromoResponse{}

	for _, promo := range promos {
		isEnded := h.promoUsecase.CheckPromoEndDate(ctx, &promo)

		res = append(res, dto.PromoResponse{
			ID:        promo.Model.ID,
			Name:      promo.Name,
			StartDate: promo.StartDate.Format(shared.TimeLayout),
			EndDate:   promo.EndDate.Format(shared.TimeLayout),
			Amount:    promo.Amount,
			IsEnded:   isEnded,
		})
	}

	c.JSON(http.StatusOK, dto.JSONResponse{Data: res})
}

func (h *PromoHandler) GeneratePromo(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.SendPromoRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	userPromo := &model.UserPromo{
		UserID:  req.UserID,
		PromoID: req.PromoID,
	}

	userPromo, err := h.promoUsecase.GeneratePromo(ctx, userPromo)
	if err != nil {
		c.Error(err)
		return
	}

	res := dto.SendPromoResponse{
		Code:    userPromo.Code,
		IsUsed:  userPromo.IsUsed,
		UserID:  userPromo.UserID,
		PromoID: userPromo.PromoID,
	}

	c.JSON(http.StatusCreated, dto.JSONResponse{Data: res})
}

func (h *PromoHandler) RedeemPromo(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.RedeemPromoRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	userPromo := &model.UserPromo{
		Code:   req.Code,
		UserID: req.UserID,
	}

	userPromo, err := h.promoUsecase.FindPromoByCode(ctx, userPromo)
	if err != nil {
		c.Error(err)
		return
	}

	isEnded := h.promoUsecase.CheckPromoEndDate(ctx, &model.Promo{Model: gorm.Model{ID: userPromo.PromoID}})
	if isEnded {
		c.Error(shared.ErrPromoHasEnded)
		return
	}

	userPromo, err = h.promoUsecase.RedeemPromo(ctx, userPromo)
	if err != nil {
		c.Error(err)
		return
	}

	res := dto.RedeemPromoResponse{
		Code:    userPromo.Code,
		IsUsed:  userPromo.IsUsed,
		UserID:  userPromo.UserID,
		PromoID: userPromo.PromoID,
	}

	c.JSON(http.StatusCreated, dto.JSONResponse{Data: res})
}
