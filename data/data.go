package data

import (
	"database/sql"
	"fmt"
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

func (p *Pool) Add(table, item string) {
	q := fmt.Sprintf("INSERT INTO %s (text) values(?)", table)
	stmt, _ := p.DB.Prepare(q)
	stmt.Exec(item)
}

func (p *Pool) Search(table, word string) []string {
	q := fmt.Sprintf("SELECT text FROM %s WHERE text LIKE ?", table)
	s := "%" + word + "%"
	rows, _ := p.DB.Query(q, s)
	var text string
	var quotes []string
	for rows.Next() {
		rows.Scan(&text)
		quotes = append(quotes, text)
	}
	return quotes
}

func (p *Pool) Random(table string) []string {
	// For a much better performance use:
	// SELECT * FROM table WHERE id IN (SELECT id FROM table ORDER BY RANDOM() LIMIT x)
	// SQL engines first load projected fields of rows to memory then sort them,
	// here we just do a random sort on id field of each row which is in memory because it's indexed,
	// then separate X of them, and find the whole row using these X ids.
	// So this consume less RAM and CPU as table grows!
	q := fmt.Sprintf("SELECT text FROM %s WHERE id IN (SELECT id FROM %s ORDER BY RANDOM() LIMIT 10)", table, table)
	rows, _ := p.DB.Query(q)
	var text string
	var quotes []string
	for rows.Next() {
		rows.Scan(&text)
		quotes = append(quotes, text)
	}
	return quotes
}

func (p *Pool) CreateTable(table string) {
	s := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, text TEXT)", table)
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}
