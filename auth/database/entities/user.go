package entities

type User struct {
	Id       string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name     string `gorm:"not null"`
	Surname  string `gorm:"not null"`
	Email    string `gorm:"not null;unique" json:"-"`
	Password string `gorm:"not null" json:"-"`
}
