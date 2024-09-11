package dtos

type Register struct {
	FullName        string `form:"fullName"`
	Email           string `form:"email" binding:"email,required"`
	Password        string `form:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirmPassword" binding:"eqfield=Password"`
}

type Login struct {
	Email    string `form:"email" binding:"email"`
	Password string `form:"password"`
}

type Token struct {
	Token string `json:"token"`
}
