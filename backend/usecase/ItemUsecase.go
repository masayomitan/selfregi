package usecase

import (
    "log"
		"selfregi/entity"
		"selfregi/model"
		"context"
		"selfregi/repository"
		_ "github.com/go-sql-driver/mysql"
)


func ItemCreate (ctx context.Context, item *entity.ItemEntity) (*model.Item, error) {

		var i repository.IItemRepository = &model.Item{}
		item, err := i.CreateItem(ctx, item)

		if err != nil {
				log.Fatalf("failed create visitor: %v", err)
		}

		return nil, nil
}
