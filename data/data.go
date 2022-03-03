package data

import (
	"database/sql"
	"strings"
)

// If not struct, *sql.DB can't be accessed by AddQuote func
type Pool struct {
	DB *sql.DB
}

func NewPool(db *sql.DB) *Pool {
	return &Pool{
		DB: db,
	}
}

func (p *Pool) SearchS() []Quote {
	rows, _ := p.DB.Query(`
			SELECT * FROM S WHERE TEXT LIKE "%aware%"
	`)

	var id int
	var text string
	quotes := []Quote{}

	for rows.Next() {
		rows.Scan(&id, &text)
		quotes = append(quotes, Quote{ID: id, Text: text})
	}
	return quotes
}

func (p *Pool) GetS() []Quote {
	// For a much better performance use:
	// SELECT * FROM table WHERE id IN (SELECT id FROM table ORDER BY RANDOM() LIMIT x)
	// SQL engines first load projected fields of rows to memory then sort them,
	// here we just do a random sort on id field of each row which is in memory because it's indexed,
	// then separate X of them, and find the whole row using these X ids.
	// So this consume less RAM and CPU as table grows!
	rows, _ := p.DB.Query(`
			SELECT * FROM S WHERE ID IN (SELECT ID FROM S ORDER BY RANDOM() LIMIT 10)
	`)
	var id int
	var text string
	quotes := []Quote{}
	for rows.Next() {
		rows.Scan(&id, &text)
		text = strings.ReplaceAll(strings.Trim(text, "'"), "''", "'")
		quotes = append(quotes, Quote{ID: id, Text: text})
	}
	return quotes
}

func (p *Pool) AddS(text string) {
	stmt, _ := p.DB.Prepare(`
		INSERT INTO S (TEXT) values(?)
	`)

	stmt.Exec(text)
}

func CreateTableS(db *sql.DB) {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS S (
			"ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"TEXT" TEXT
		);
	`)
	stmt.Exec()
}

func (p *Pool) GetZGXW() []Quote {
	rows, _ := p.DB.Query(`
			SELECT * FROM ZGXW WHERE ID IN (SELECT ID FROM ZGXW ORDER BY RANDOM() LIMIT 10)
	`)
	var id int
	var text string
	quotes := []Quote{}
	for rows.Next() {
		rows.Scan(&id, &text)
		quotes = append(quotes, Quote{ID: id, Text: text})
	}
	return quotes
}

func CreateTableZGXW(db *sql.DB) {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS ZGXW (
			"ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"TEXT" TEXT
		);
	`)
	stmt.Exec()
}
