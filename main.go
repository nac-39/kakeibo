package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	database "github.com/nac-39/kakeibo/pkg/infrastructure"
)

func run() error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	// list all authors
	queries := database.New(conn)
	authors, err := queries.GetBox(ctx, 0)
	if err != nil {
		return err
	}
	log.Println(authors)
	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
