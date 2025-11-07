package service

import "github.com/Aboagye-Dacosta/shopBackend/internal/database/models"

type Service struct {
	UserService       *UserService
	PermissionService *PermissionService
	RoleService       *RoleService
}

func NewService(m *models.Models) *Service {
	return &Service{
		UserService:       &UserService{m.Users},
		PermissionService: &PermissionService{m.Permissions},
		RoleService:       &RoleService{m.Roles},
	}
}
