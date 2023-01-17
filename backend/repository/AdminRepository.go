package repository

import (
	"context"
	"selfregi/model"
)

type IAdminRepository interface {
	CreateAdmin(ctx context.Context) (*model.Admin, error)
	// UpdateAdmin(ctx context.Context, admin *model.Admin) (*model.Admin, error)
}
