// TODO import Echo framework and configure HTTP Server
// TODO configure routes #8

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	utils "pennywise/utilities"
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
	e.POST("/upload-csv", uploadCSVHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}


func uploadCSVHandler(c echo.Context) error { // Updated to use echo.Context
    if c.Request().Method != http.MethodPost {
        return c.String(http.StatusMethodNotAllowed, "Invalid request method") // Updated to use echo.Context
    }

    var body map[string]interface{}
    err := json.NewDecoder(c.Request().Body).Decode(&body) // Updated to use echo.Context
    if err != nil {
        log.Printf("Error decoding request body: %v", err)
        return c.String(http.StatusBadRequest, "Error parsing request body") // Updated to use echo.Context
    }

    csvString, ok := body["data"].(string)
    if !ok {
        return c.String(http.StatusBadRequest, "No CSV data found") // Updated to use echo.Context
    }

    log.Printf("Data preview (first 200 chars): %s", csvString[:200])
    
    // Process the CSV string (you will need to implement parseCSV and createCSVTable)
    csvData, err := ParseCSV(csvString)
    if err != nil {
        log.Printf("Error Parsing CSV: %v", err)
        errorHTML := fmt.Sprintf(`
        <div class="book-detail">
            <h3>Data Preview</h3>
            <h4>Error Parsing CSV: %v</h4>
        </div>`, err)
        return c.HTML(http.StatusInternalServerError, errorHTML) // Updated to use echo.Context
    }

    log.Printf("CSV successfully processed, length: %d", len(csvData))
    updateHTML, err := utils.CreateCSVTable(csvData, 50)
	if err != nil {
		fmt.Println("Could not create CSV table.", err)
	}
    return c.String(http.StatusOK, updateHTML) // Updated to use echo.Context
}