
package model

import (
	"context"
	"fmt"
	"selfregi/ent"
	"selfregi/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
  entAdmin *ent.Admin
}

type SAdmin struct {
  *ent.Admin
}

func (a *Admin) CreateAdmin (ctx context.Context) (*Admin, error)  {
	db, _ := database.EntOpen()
	fmt.Println("model come in")
	admin, err := db.Admin.Create().
	SetName("no name").
	SetPassword("aaa").
	Save(ctx)
	a.entAdmin = admin

	return a, err
}

func (a *Admin) UpdateAdmin (ctx context.Context) (*Admin, error)  {
	db, _ := database.EntOpen()
	var id int = a.entAdmin.ID
	fmt.Println("model come in")
	admin, err := db.Admin.UpdateOneID(id).
	SetName(a.entAdmin.Name).
	SetPassword(*a.entAdmin.Password).
	Save(ctx)
	a.entAdmin = admin

	return a, err
}
