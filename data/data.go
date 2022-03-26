package data

import (
	"database/sql"
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

func (p *Pool) GetS() []string {
	// For a much better performance use:
	// SELECT * FROM table WHERE id IN (SELECT id FROM table ORDER BY RANDOM() LIMIT x)
	// SQL engines first load projected fields of rows to memory then sort them,
	// here we just do a random sort on id field of each row which is in memory because it's indexed,
	// then separate X of them, and find the whole row using these X ids.
	// So this consume less RAM and CPU as table grows!
	rows, _ := p.DB.Query(`SELECT text FROM sentences WHERE id IN (SELECT id FROM sentences ORDER BY RANDOM() LIMIT 10)`)
	var text string
	var quotes []string
	for rows.Next() {
		rows.Scan(&text)
		quotes = append(quotes, text)
	}
	return quotes
}

func (p *Pool) GetZ() []string {
	rows, _ := p.DB.Query(`SELECT text FROM zgxw WHERE id IN (SELECT id FROM zgxw ORDER BY RANDOM() LIMIT 10)`)
	var text string
	var quotes []string
	for rows.Next() {
		rows.Scan(&text)
		quotes = append(quotes, text)
	}
	return quotes
}

func (p *Pool) GetW() []string {
	rows, _ := p.DB.Query(`SELECT text FROM words WHERE id IN (SELECT id FROM words ORDER BY RANDOM() LIMIT 10)`)
	var text string
	var quotes []string
	for rows.Next() {
		rows.Scan(&text)
		quotes = append(quotes, text)
	}
	return quotes
}

func (p *Pool) AddS(text string) {
	stmt, _ := p.DB.Prepare(`INSERT INTO sentences (text) values(?)`)
	stmt.Exec(text)
}

func (p *Pool) SearchS(word string) []string {
	rows, _ := p.DB.Query(`SELECT * FROM setences WHERE text LIKE "%?%"`)
	var text string
	var quotes []string
	for rows.Next() {
		rows.Scan(&text)
		quotes = append(quotes, text)
	}
	return quotes
}

func CreateTableS(db *sql.DB) {
	stmt, _ := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "sentences" (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	text TEXT)`)
	stmt.Exec()
}

func CreateTableZ(db *sql.DB) {
	stmt, _ := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "zgxw" (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	text TEXT)`)
	stmt.Exec()
}
