package service

import (
	"context"
	"time"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"github.com/Aboagye-Dacosta/shopBackend/internal/errors"
)

type PermissionService struct {
	permissions *models.PermissionModel
}

func (s *PermissionService) GetPermissions(ctx context.Context) ([]*models.Permission, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var permissions []*models.Permission

	if err := s.permissions.DB.WithContext(ctx).Find(&permissions).Error; err != nil {
		return nil, errors.FromDb(entities.PERMISSIONS, err)
	}

	return permissions, nil
}
