package main

import (
	"bytes"
	"context"
	"database/sql"
	"embed"
	"encoding/csv"
	"fmt"

	"strconv"

	_ "github.com/glebarez/go-sqlite"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

//go:embed QandA.csv
var database embed.FS

// Greet returns a greeting for the given name
func (a *App) Lookup(userinput string) string {
	allRecords := readCSV()

	db, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		return fmt.Sprintf("Error with SQL database creation from CSV: %s", err)
	}
	defer db.Close()

	//Read in CSV to database for queries
	stmt, err1 := db.Prepare(`create table if not exists qat(id integer, question text, answer text)`)
	if err1 != nil {
		return fmt.Sprintf("dbPrepare create table failed: %s", err1)
	}

	if _, err = stmt.Exec(); err != nil {
		return fmt.Sprintf("stmt.Exec create table failed: %s", err)
	}
	for id := 0; id < len(allRecords); id++ {
		record := allRecords[id]
		question := record[0]
		answer := record[1]
		stmt, err = db.Prepare("insert into qat(id, question, answer) values(?, ?, ?)")
		if err != nil {
			return fmt.Sprintf("db.Prepare insert statement failed: %s", err)
		}
		_, err = stmt.Exec(strconv.Itoa(id), question, answer)
		if err != nil {
			return fmt.Sprintf("db.Exec populate table failed: %s", err)
		}
	}

	//The Query to the database to return an answer
	var answer string
	userinput = "'" + userinput + "'"
	err = db.QueryRow("SELECT answer FROM qat WHERE question = ?", userinput).Scan(&answer)
	if err != nil {
		return fmt.Sprintf("Query failed as '%s' is not in database", userinput)
	}
	return answer
}

// Reads CSV file
func readCSV() [][]string {
	csvInput, _ := database.ReadFile("QandA.csv")
	csvReader := csv.NewReader(bytes.NewReader(csvInput))
	csvReader.FieldsPerRecord = -1
	allRecords, _ := csvReader.ReadAll()
	return allRecords
}
