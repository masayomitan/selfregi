package repository

import (
	"context"
	"selfregi/model"
)

type IVisitorRepository interface {
	CreateVisitor(ctx context.Context) (*model.Visitor, error)
}
