package core

import (
	"context"
	"database/sql"
	"embed"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

//go:embed schema.sql
var schemaFile embed.FS

type Database struct {
	UserStore UserStore
	DB        *sql.DB
}

func MustGetDatabase(ctx context.Context, config *Config /*const*/) *Database {
	db, err := sql.Open("pgx", config.DatabaseUrl)
	if err != nil {
		panic(WrappedError(err, "failed to open database"))
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}

	if err := initializeSchema(ctx, db); err != nil {
		panic(WrappedError(err, "failed to initialize schema"))
	}

	return &Database{
		DB:        db,
		UserStore: NewUserStore(db),
	}
}

func initializeSchema(ctx context.Context, db *sql.DB) error {
	b, err := schemaFile.ReadFile("schema.sql")
	if err != nil {
		return WrappedError(err, "failed to load schema file")
	}
	if _, err := db.ExecContext(ctx, string(b)); err != nil {
		return WrappedError(err, "failed to execute sql schema")
	}
	return nil
}
