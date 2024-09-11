package models

type Users struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	RoleId   int    `json:"roleId" db:"role_id"`
}
