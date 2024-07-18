package services

import (
	"github.com/TheoremN1/Coins/database/models"
	"gorm.io/gorm"
)

type RolesService struct {
	database *gorm.DB
}

func NewRolesService(database *gorm.DB) *RolesService {
	return &RolesService{database}
}

func (rolesService *RolesService) IsRoleExist(key string) bool {
	var roles []*models.Role
	rolesService.database.
		Where("key = ?", key).
		Find(&roles)
	return len(roles) > 0
}

func (rolesService *RolesService) GetUserRole(userId int) (*models.Role, int) {
	roles := rolesService.GetAllRoles()
	if roles != nil {
		for _, role := range *roles {
			for i, user := range role.Users {
				if user.Id == userId {
					return role, i
				}
			}
		}
	}
	return nil, -1
}

func (rolesService *RolesService) GetAllRoles() *[]*models.Role {
	var roles []*models.Role
	rolesService.database.Model(&models.Role{}).Find(&roles)
	if len(roles) > 0 {
		return &roles
	}
	return nil
}

func (rolesService *RolesService) GetRole(key string) *models.Role {
	if rolesService.IsRoleExist(key) {
		var role models.Role
		rolesService.database.
			Where("key = ?", key).
			First(&role)
		return &role
	}
	return nil
}
