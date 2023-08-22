package main

import (
	"bytes"
	"context"
	"database/sql"
	"embed"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"

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

func (a *App) Lookup(userinput string) string {
	//Initialize Lookup
	var answer string
	allRecords := ReadData() //reads all records from embedded QandA.csv

	//Ensures correct user input is a single word
	words := strings.Fields(userinput) //uses ' ' as default delimiter
	wordCount := len(words)
	if wordCount == 0 {
		answer = "No input given. Please input a word."
	}
	if wordCount > 1 {
		answer = "Too many words given. Please input one word."
	}
	if wordCount == 1 {

		userinput = ProcessInput(userinput) //Choose representation of input data

		//Creates SQL database representation of corpus from allRecords that could be modified to include tf-idf
		db, _ := sql.Open("sqlite", "file::memory:?cache=shared")                                       //temp
		stmt, _ := db.Prepare(`create table if not exists qat(id integer, question text, answer text)`) //headers
		stmt.Exec()
		for id := 0; id < len(allRecords); id++ { //populating database
			record := allRecords[id]
			question := record[0] //Use this field to create 2nd SQL database with ID for a tf-idf representation
			answer := record[1]
			stmt, _ = db.Prepare("insert into qat(id, question, answer) values(?, ?, ?)")
			stmt.Exec(strconv.Itoa(id), question, answer)
		}

		//Querying SQL database for a match
		err := db.QueryRow("SELECT answer FROM qat WHERE question = ?", userinput).Scan(&answer) //performs actual search
		if err != nil {
			answer = fmt.Sprintf("Query failed as '%s' is not in database", userinput)
		}
	}
	return answer
}

// Reads CSV file from embedded binary
func ReadData() [][]string {
	csvInput, _ := database.ReadFile("QandA.csv")
	csvReader := csv.NewReader(bytes.NewReader(csvInput))
	csvReader.FieldsPerRecord = -1
	allRecords, _ := csvReader.ReadAll()
	return allRecords
}

// could be tokenized/lemmatized here to calculate similarity with a tf-idf representation
func ProcessInput(userinput string) string {
	userinput = "'" + userinput + "'"
	return userinput
}
