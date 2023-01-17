package usecase

import (
	"context"
	"fmt"
	"log"
	"selfregi/model"
	"selfregi/repository"
	_ "github.com/go-sql-driver/mysql"
)

func Create (ctx context.Context) (*model.Visitor, error) {
	var v repository.IVisitorRepository = &model.Visitor{}
	visitor, err := v.CreateVisitor(ctx)
	fmt.Println(v)
	if err != nil {
		log.Fatalf("failed create visitor: %v", err)
	}

	return visitor, err
}