package dto

type PromoRequest struct {
	Name      string `json:"name" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	Amount    int64  `json:"amount" binding:"required"`
}

type SendPromoRequest struct {
	UserID  uint `json:"user_id" binding:"required"`
	PromoID uint `json:"promo_id" binding:"required"`
}

type SendPromoResponse struct {
	Code    string `json:"code"`
	IsUsed  bool   `json:"is_used"`
	UserID  uint   `json:"user_id"`
	PromoID uint   `json:"promo_id"`
}

type RedeemPromoRequest struct {
	Code   string `json:"code" binding:"required"`
	UserID uint   `json:"user_id" binding:"required"`
}

type PromoResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Amount    int64  `json:"amount"`
	IsEnded   bool   `json:"is_ended"`
}

type RedeemPromoResponse struct {
	Code    string `json:"code"`
	IsUsed  bool   `json:"is_used"`
	UserID  uint   `json:"user_id"`
	PromoID uint   `json:"promo_id"`
}
