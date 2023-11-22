package dto

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Birthday string `json:"birthday" binding:"required"`
}

type UserUpdateRequest struct {
	IsVerified bool `json:"is_verified" binding:"required"`
}

type UserResponse struct {
	ID         uint   `json:"id"`
	Email      string `json:"email"`
	Birthday   string `json:"birthday"`
	IsVerified bool   `json:"is_verified"`
	IsBirthday bool   `json:"is_birthday"`
}
