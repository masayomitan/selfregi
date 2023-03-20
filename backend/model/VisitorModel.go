package model

import (
	"context"
	// "fmt"
	"selfregi/ent"
	"selfregi/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)

type Visitor struct {
  entVisitor *ent.Visitor
}

func (v *Visitor) CreateVisitor (ctx context.Context) (*Visitor, error)  {
	db, _ := database.EntOpen()
	visitor, err := db.Visitor.Create().
	SetName("no name").
	SetSex(0).
	Save(ctx)
	v.entVisitor = visitor

	return v, err
}
