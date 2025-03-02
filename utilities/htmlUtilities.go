package utilities

import (
    "fmt"
    "html/template"
    "strings"
)

// createCSVTable generates an HTML table from the provided JSON-like data.
func CreateCSVTable(jsonData []map[string]any, previewCount int) (string, error) {
    if len(jsonData) == 0 {
        return "", fmt.Errorf("no data provided")
    }

    var tablePrefix = "<table>"
    var tableSuffix = "</table>"
    var htmlRows strings.Builder
    headers := make([]string, 0)

    // Extract headers
    for key := range jsonData[0] {
        headers = append(headers, key)
    }

    // Create header row
    htmlRows.WriteString("<tr>")
    for _, header := range headers {
        htmlRows.WriteString(fmt.Sprintf("<td><b>%s</b></td>", header))
    }
    htmlRows.WriteString("</tr>")

    // Create data rows
    for i, row := range jsonData {
        if i >= previewCount {
            break
        }
        htmlRows.WriteString("<tr>")
        for _, header := range headers {
            htmlRows.WriteString(fmt.Sprintf("<td>%v</td>", row[header]))
        }
        htmlRows.WriteString("</tr>")
    }

    // Create final HTML
    htmlTemplate := `
    <div class="card-body">
        <h3>Data Preview</h3>
        {{.TablePrefix}}
            {{.HtmlRows}}
        {{.TableSuffix}}
    </div>
    `
    t, err := template.New("html").Parse(htmlTemplate)
    if err != nil {
        return "", err
    }

    var result strings.Builder
    err = t.Execute(&result, struct {
        TablePrefix string
        HtmlRows    string
        TableSuffix string
    }{
        TablePrefix: tablePrefix,
        HtmlRows:    htmlRows.String(),
        TableSuffix: tableSuffix,
    })
    if err != nil {
        return "", err
    }

    return result.String(), nil
}