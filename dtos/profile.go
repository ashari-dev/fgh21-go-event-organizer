package dtos

type FormEditProfile struct {
	Fullname    string `form:"fullname"`
	Username    string `form:"username"`
	Email       string `form:"email" binding:"email"`
	PhoneNumber string `form:"phoneNumber"`
	Gender      int    `form:"gender"`
	Profession  string `form:"profession"`
	Nationality int    `form:"nationality"`
	BirthDate   string `form:"birthDate"`
}
