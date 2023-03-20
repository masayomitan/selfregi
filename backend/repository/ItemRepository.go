package repository

import (
		"context"
		"selfregi/entity"
		// "selfregi/model"
)

type IItemRepository interface {
    CreateItem(ctx context.Context, item *entity.ItemEntity) (*entity.ItemEntity, error)
}