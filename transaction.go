package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

)


type Transaction struct {
	ID string
	TransactionDate time.Time
	Description string
	Amount float64
	Balance float64	
	Category string
	FinancialEntityName string
}

// Iterate over the returned rows, scans each row into a User struct,
// append the user to a slice, and prints the user's details.
func queryTransactions(db *sql.DB) []Transaction {
	rows, err := db.Query("SELECT * FROM transactions")
	if err != nil {
		// If the query execution fails, print the error message to standard error and exit.
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	// Ensure that the rows are closed after we finish processing them.
	defer rows.Close()

	// Slice to hold all the User objects fetched from the database.
	var transactions []Transaction

	// Iterate over the rows returned by the query.
	for rows.Next() {
		var t Transaction
		// Scan the columns from the current row into the user struct fields.
		if err := rows.Scan(&t.ID, &t.TransactionDate, &t.Description, &t.Amount, &t.Balance, &t.Category, &t.FinancialEntityName); err != nil {
			// If scanning fails, print an error message and return from the function.
			fmt.Println("Error scanning row:", err)
			return nil
		}

		// Add the user to the slice.
		transactions = append(transactions, t)
		// Print the user's ID and Name.
		fmt.Println(t)
	}

	// Check if there were any errors encountered during the row iteration.
	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return transactions
}

func genTransactionHtml(db *sql.DB) string {
	prefix := `<div id="transactions-table" class="overflow-x-auto rounded-box border border-base-content/5 bg-base-100">
  <table class="table">
    <!-- head -->
    <thead>
      <tr>
        <th>Transaction Date</th>
        <th>Description</th>
        <th>Amount</th>
        <th>Balance</th>
        <th>Type</th>
        <th>Bank Name</th>
      </tr>
    </thead>
	<tbody>`

	rows := ""

	suffix := `</tbody>
  </table>
</div>`
	transactions := queryTransactions(db)

	for _,t := range transactions {
		rowHTML := `<tr>`
		rowHTML += "<td>" + t.TransactionDate.Format("2006-01-02") + "</td>"
		rowHTML += "<td>" + t.Description + "</td>"
		rowHTML += "<td>" + strconv.FormatFloat(t.Amount, 'f', 2, 64) + "</td>"
		rowHTML += "<td>" + strconv.FormatFloat(t.Balance, 'f', 2, 64) + "</td>"
		rowHTML += "<td>" + t.Category + "</td>"
		rowHTML += "<td>" + t.FinancialEntityName + "</td>"

		rows += rowHTML
	}

	return prefix + rows + suffix
	
}