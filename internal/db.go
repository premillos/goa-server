package internal

import (
	"context"
	"fmt"

	"com.goa/ent"
	"com.goa/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func InitDb(ctx context.Context) (clearFunc func(), err error) {

	client, err := ent.Open(config.C.Db.Type, config.C.Db.Dsn)

	if err != nil {
		fmt.Printf("failed opening connection to %s: %v", config.C.Db.Dsn, err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		fmt.Printf("failed creating schema resources: %v", err)
	}

	return func() {
		client.Close()
	}, err
}
