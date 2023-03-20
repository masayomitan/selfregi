package model

import (
		"context"
		"log"
		"selfregi/ent"
		"selfregi/entity"
		"selfregi/ent/item"
		"selfregi/infrastructure/database"
		_ "github.com/go-sql-driver/mysql"
		// "fmt"
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

func (i *Item) CreateItem(ctx context.Context, item *entity.ItemEntity) (*entity.ItemEntity, error) {
		// db, _ := database.EntOpen()
		// fmt.Println(item)
		// newItem, err := db.Item.Create().
		// SetName(item.Name).
		// SetPrice(item.Price).
		// SetCategoryID(item.CategoryID).
		// Save(ctx)
		// i.entItem = newItem
		// if err != nil {
		// 		log.Fatal(err)
		// }
		return nil, nil
}