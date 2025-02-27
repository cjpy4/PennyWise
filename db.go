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

func queryDB(db *sql.DB, query string) []map[string]any {
	rows, err := db.Query(query)
	if err != nil {
		// If the query execution fails, print the error message to standard error and exit.
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	// Ensure that the rows are closed after we finish processing them.
	defer rows.Close()

	// Instantiate JSON rows
	rowJSON := make([]map[string]any, 0)

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("Error getting columns:", err)
		return nil // Return nil to indicate an error
	}

	// Iterate over the rows returned by the query.
	for rows.Next() {
		// Create a slice to hold the values for each column
		values := make([]any, len(columns))
		// Create a map to hold the column name and value pairs
		rowMap := make(map[string]any)

		// Scan the columns into the values slice
		if err := rows.Scan(values...); err != nil {
			// If scanning fails, print an error message and return from the function.
			fmt.Println("Error scanning row:", err)
			return nil
		}

		// Populate the rowMap with column names and their corresponding values
		for i, colName := range columns {
			rowMap[colName] = values[i]
		}

		// Append the rowMap to rowJSON
		rowJSON = append(rowJSON, rowMap)
	}

	// Check if there were any errors encountered during the row iteration.
	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return rowJSON // Return the JSON rows
}

func executeSQL(db *sql.DB, stmt string) {
	result, err := db.Exec(stmt)
	if err != nil {
		// If the query execution fails, print the error message to standard error and exit.
		fmt.Fprintf(os.Stderr, "failed to execute sql: %v\n", err)
		os.Exit(1)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Cannot obtain affected rows", err)
	}
	
	fmt.Println("rows affected: ", rowsAffected)
}
