package models

type ChangePasswordRequest struct {
	OldPassword             string `json:"old_password" binding:"required"`
	NewPassword             string `json:"new_password" binding:"required"`
	NewPasswordConfirmation string `json:"new_password_confirmation" binding:"required"`
}

type JWTData struct {
	Phone string
	UserID string
}