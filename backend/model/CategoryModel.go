package model

import (
	"context"
	"log"
	"selfregi/ent"
	"selfregi/ent/category"
	"selfregi/infrastructure/database"

	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
  entCategory *ent.Category
}

type Categories struct {
  *ent.Categories
}


func (c *Categories) GetCategories (ctx context.Context) ([]*ent.Category, error)  {
	client, _ := database.EntOpen()
	Categories, err := client.Category.
		Query().
		// Select("name", "is_display").
		Where(category.DeletedAtIsNil()).
		All(ctx)
		if err != nil {
			log.Fatal(err)
		}
	return Categories, err
}

func (c *Categories) GetCategoriesDisplayOn (ctx context.Context) ([]*ent.Category, error)  {
	client, _ := database.EntOpen()
	Categories, err := client.Category.
		Query().
		Select("name", "is_display").
		Where(category.IsDisplay(true)).
		Where(category.DeletedAtIsNil()).
		All(ctx)
		if err != nil {
			log.Fatal(err)
		}
	return Categories, err
}