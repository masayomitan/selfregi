package model

import (
	"context"
	"log"
	"selfregi/ent"
	"selfregi/ent/item"
	"selfregi/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
  entItem *ent.Item
}

type Items struct {
  entItem[] *ent.Items
}

func (i *Item) GetItemDisplayOn(ctx context.Context, itemId int) (*ent.Item, error)  {
	client, _ := database.EntOpen()
	item, err := client.Item.
		Query().
		Where(item.ID(itemId)).
		Where(item.IsDisplay(true)).
		Where(item.DeletedAtIsNil()).
		First(ctx)
		if err != nil {
			log.Fatal(err)
		}
	return item, err
}

func (i *Items) GetItemsDisplayOnFromCategoryId(ctx context.Context, categoryId int) ([]*ent.Item, error)  {
	client, _ := database.EntOpen()
	items, err := client.Item.
		Query().
		Where(item.CategoryID(categoryId)).
		Where(item.IsDisplay(true)).
		Where(item.DeletedAtIsNil()).
		All(ctx)
		if err != nil {
			log.Fatal(err)
		}
	return items, err
}