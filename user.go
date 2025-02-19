package main

import (
	"database/sql"
	"fmt"
	"os"
)

type User struct {
	ID   int
	Name string
}

// Iterate over the returned rows, scans each row into a User struct,
// append the user to a slice, and prints the user's details.
func queryUsers(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		// If the query execution fails, print the error message to standard error and exit.
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	// Ensure that the rows are closed after we finish processing them.
	defer rows.Close()

	// Slice to hold all the User objects fetched from the database.
	var users []User

	// Iterate over the rows returned by the query.
	for rows.Next() {
		var user User
		// Scan the columns from the current row into the user struct fields.
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			// If scanning fails, print an error message and return from the function.
			fmt.Println("Error scanning row:", err)
			return
		}

		// Add the user to the slice.
		users = append(users, user)
		// Print the user's ID and Name.
		fmt.Println(user.ID, user.Name)
	}

	// Check if there were any errors encountered during the row iteration.
	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}
}

func createUser(db *sql.DB, name string) error {
	_, err := db.Exec("INSERT INTO users (name) VALUES (?)", name)
	return err
}