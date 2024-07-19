package services

import (
	"fmt"

	"github.com/TheoremN1/Coins/database/models"
	"gorm.io/gorm"
)

type UsersService struct {
	database     *gorm.DB
	rolesService *RolesService
}

func NewUsersService(database *gorm.DB, rolesService *RolesService) *UsersService {
	return &UsersService{database, rolesService}
}

func (usersService *UsersService) IsUserExist(id int) bool {
	var users []*models.User
	usersService.database.
		Where("id = ?", id).
		Find(&users)
	return len(users) > 0
}

func (usersService *UsersService) IsUsernameExist(username string) bool {
	var users []models.User
	usersService.database.
		Where("username = ?", username).
		Find(&users)
	return len(users) > 0
}

func (usersService *UsersService) GetUser(id int) *models.User {
	if usersService.IsUserExist(id) {
		var user models.User
		usersService.database.
			Where("id = ?", id).
			First(&user)
		return &user
	}
	return nil
}

func (usersService *UsersService) GetAllUsers() []*models.User {
	var users []*models.User
	usersService.database.Model(&models.User{}).Find(&users)
	if len(users) > 0 {
		return users
	}
	return nil
}

func (usersService *UsersService) NewUser(username string, name string, surname string, roleKey string) (bool, int) {
	if !usersService.IsUsernameExist(username) && usersService.rolesService.IsRoleExist(roleKey) {
		user := models.User{
			Username: username,
			Name:     name,
			Surname:  surname,
			Balance:  0,
		}
		usersService.database.Save(&user)
		if usersService.SetRole(user.Id, roleKey) {
			return true, user.Id
		}
		usersService.DeleteUser(user.Id)
	}
	fmt.Printf("UsernameExist: %t\nRoleExist: %t\n", usersService.IsUsernameExist(username), usersService.rolesService.IsRoleExist(roleKey))
	return false, -1
}

func (usersService *UsersService) EditUser(id int, name string, surname string) bool {
	if usersService.IsUserExist(id) {
		user := usersService.GetUser(id)
		isEdited := false
		if user.Name != name {
			user.Name = name
			isEdited = true
		}
		if user.Surname != surname {
			user.Surname = surname
			isEdited = true
		}
		if isEdited {
			usersService.database.Save(&user)
			return isEdited
		}
	}
	return false
}

func (usersService *UsersService) DeleteUser(id int) bool {
	if usersService.IsUserExist(id) {
		user := usersService.GetUser(id)
		usersService.database.Delete(&user)
		return true
	}
	return false
}

func (usersService *UsersService) SetRole(userId int, roleKey string) bool {
	if usersService.IsUserExist(userId) && usersService.rolesService.IsRoleExist(roleKey) {
		user := usersService.GetUser(userId)
		oldRole, index := usersService.rolesService.GetUserRole(userId)
		role := usersService.rolesService.GetRole(roleKey)
		if oldRole != role && role != nil {
			if oldRole != nil {
				role.Users = append(role.Users[:index], role.Users[index+1:]...)
			}
			role.Users = append(role.Users, *user)
			return true
		}
	}
	fmt.Printf("UserExist: %t\nRoleExist: %t\n", usersService.IsUserExist(userId), usersService.rolesService.IsRoleExist(roleKey))
	return false
}
