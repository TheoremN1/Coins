package models

type AuthorizationData struct {
	Email    string `form:"email" json:"email" bindind:"required,email"`
	Password string `form:"password" json:"password" bindind:"required,min=8"`
}
