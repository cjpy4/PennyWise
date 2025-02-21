package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"github.com/joho/godotenv"
	"github.com/tursodatabase/go-libsql"
)
 func connectToDB() (*libsql.Connector, *sql.DB, string, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, nil, "", fmt.Errorf("error loading .env file: %w", err)
	}

	// Database credentials are in our private chat
	dbName := "local.db"
	primaryUrl := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		return nil, nil, "", fmt.Errorf("error creating temporary directory: %w", err)
	}

	dbPath := filepath.Join(dir, dbName)

	syncInterval := time.Minute

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl,
		libsql.WithAuthToken(authToken),
		libsql.WithSyncInterval(syncInterval),
	)

	
	if err != nil {
		return nil, nil, dir, fmt.Errorf("error creating connector: %w", err)
	}

	db := sql.OpenDB(connector)

	return connector, db, dir, nil

 }
