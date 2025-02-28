// TODO import Echo framework and configure HTTP Server
// TODO configure routes #8

package main

import (
	"fmt"
	"os"
	"github.com/labstack/echo/v4"
)

// Handler
// func transactions(c echo.Context) error {
// 	return c.File("public/transactions.html")
// }

func main() {

	connector, db, dir, err := connectToDB()
	if err != nil {
		fmt.Printf("Error connecting to DB: %v\n", err)
		return
	}
	// Clean up resources when main returns.
	defer os.RemoveAll(dir)
	defer connector.Close()
	defer db.Close()

	// createUser(db, "Lando Coderissian")
	// createUser(db, "CJ Pyethonian")
	queryUsers(db)

	// Echo Server and Routes
	// New instance of Echo
	e := echo.New()
	// Routes
	e.Static("/", "public")
	e.File("/transaction", genTransactionHtml(db))
	e.File("/overview", "public/overview-card.html")
	e.File("/budgets", "public/budgets-table.html")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
