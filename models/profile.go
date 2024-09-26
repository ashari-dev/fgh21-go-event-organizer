package models

type Profile struct {
	Id            int     `json:"id"`
	Picture       *string `json:"picture" db:"picture"`
	FullName      string  `json:"fullName" db:"full_name"`
	BirthDate     *string `json:"birthDate" db:"birth_date"`
	Gender        *int    `json:"gender" db:"gender"`
	PhoneNumber   *string `json:"phone-number" db:"phone_number"`
	Profession    *string `json:"profession" db:"profession"`
	NationalityId *int    `json:"nationality-id" db:"nationality_id"`
	UsersId       int     `json:"usersId" db:"users_id"`
}

type JoinUserProfile struct {
	Id          int     `json:"id"`
	Picture     *string `json:"picture"`
	FullName    string  `json:"fullname" db:"full_name"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phoneNumber" db:"phone_number"`
	Gender      *int    `json:"gender"`
	Profession  *string `json:"profession"`
	Nationality *int    `json:"nationalityId" db:"nationality_id"`
	BirthDate   *string `json:"birthDate" db:"birth_date"`
	RoleId      *int    `json:"roleId" db:"role_id"`
}
