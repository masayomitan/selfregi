package database

import (
	"context"
	"fmt"
	"log"
	"selfregi/ent"
	"selfregi/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
)

type EntClient struct {
	*ent.Client
}

func Migrate(client *ent.Client) {
		ctx := context.Background()
    err := client.Schema.Create(
        ctx,
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true),
    )
    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    } else {
			fmt.Println("migrate ok")
		}
}