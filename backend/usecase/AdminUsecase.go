package usecase

import (
	"context"
	"fmt"
	"log"
	"selfregi/model"
	"selfregi/repository"
	_ "github.com/go-sql-driver/mysql"
)

func AdminRegist (ctx context.Context, admin model.Admin) (*model.Admin, error) {
	var v repository.IAdminRepository = &model.Admin{}
	newAdmin, err := v.CreateAdmin(ctx)
	fmt.Println("admin")
	fmt.Println(admin)
	if err != nil {
		log.Fatalf("failed create visitor: %v", err)
	}

	return newAdmin, err
}

func AdminUpdate (ctx context.Context, admin model.Admin) (*model.Admin, error) {
	var v  = &model.Admin{}
	newAdmin, err := v.UpdateAdmin(ctx)
	fmt.Println("admin")
	fmt.Println(admin)
	if err != nil {
		log.Fatalf("failed create visitor: %v", err)
	}

	return newAdmin, err
}