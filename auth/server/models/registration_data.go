package models

type RegistrationData struct {
	Name     string `form:"name" json:"name,omitempty"`
	Surname  string `form:"surname" json:"surname,omitempty"`
	Email    string `form:"email" json:"email,omitempty" bindind:"required"`
	Password string `form:"password" json:"password,omitempty" bindind:"required"`
}
