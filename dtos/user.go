package dtos

type FormUser struct {
	Email    string  `form:"email"`
	Password string `form:"password"`
	Username *string  `form:"username"`
}

type FormChangePassword struct {
	OldPassword     string `form:"oldPassword"`
	NewPassword     string `form:"newPassword"`
	ConfirmPassword string `form:"confirmPassword" binding:"eqfield=NewPassword"`
}
