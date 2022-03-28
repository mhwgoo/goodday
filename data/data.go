package data

import (
	"database/sql"
	"fmt"
)

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

func (p *Pool) CreateTableTask() {
	s := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_DATE,
		updated_at DEFAULT CURRENT_DATE,
		name TEXT NOT NULL,
		description TEXT,
		start_date TEXT,
		due_date TEXT,
		dependency INTEGER, 
		type_id INTEGER NOT NULL,
		status_id INTEGER NOT NULL,
		priority_id INTEGER NOT NULL,
		section_id INTEGER NOT NULL,
		FOREIGN KEY(type_id) REFERENCES task_type(id) ON DELETE SET NULL,
		FOREIGN KEY(status_id) REFERENCES task_status(id) ON DELETE SET NULL,
		FOREIGN KEY(priority_id) REFERENCES task_priority(id) ON DELETE SET NULL,
		FOREIGN KEY(section_id) REFERENCES sections(id) ON DELETE CASCADE,
		FOREIGN KEY(dependency) REFERENCES tasks(id) ON DELETE SET NULL
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}

func (p *Pool) CreateTableTaskType() {
	s := `CREATE TABLE IF NOT EXISTS task_type (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_DATE,
		updated_at DEFAULT CURRENT_DATE,
		name TEXT NOT NULL,
		description TEXT
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}

func (p *Pool) CreateTableTaskStatus() {
	s := `CREATE TABLE IF NOT EXISTS task_status (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_DATE,
		updated_at DEFAULT CURRENT_DATE,
		name TEXT NOT NULL,
		description TEXT
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}

func (p *Pool) CreateTableTaskPriority() {
	s := `CREATE TABLE IF NOT EXISTS task_priority (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_DATE,
		updated_at DEFAULT CURRENT_DATE,
		name TEXT NOT NULL,
		description TEXT
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}

func (p *Pool) CreateTableSection() {
	s := `CREATE TABLE IF NOT EXISTS sections (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_DATE,
		updated_at DEFAULT CURRENT_DATE,
		name TEXT NOT NULL,
		description TEXT,
		start_date TEXT,
		due_date TEXT,
		type_id INTEGER NOT NULL,
		status_id INTEGER NOT NULL,
		FOREIGN KEY(type_id) REFERENCES section_type(id) ON DELETE SET NULL,
		FOREIGN KEY(status_id) REFERENCES section_status(id) ON DELETE SET NULL
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}

func (p *Pool) CreateTableSectionType() {
	s := `CREATE TABLE IF NOT EXISTS section_type (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_DATE,
		updated_at DEFAULT CURRENT_DATE,
		name TEXT NOT NULL,
		description TEXT
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}

func (p *Pool) CreateTableSectionStatus() {
	s := `CREATE TABLE IF NOT EXISTS section_status (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_DATE,
		updated_at DEFAULT CURRENT_DATE,
		name TEXT NOT NULL,
		description TEXT
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}

func (p *Pool) CreateTableNotice() {
	s := `CREATE TABLE IF NOT EXISTS notices (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_DATE,
		updated_at DEFAULT CURRENT_DATE,
		name TEXT NOT NULL,
		description TEXT
		start_date DEFAULT CURRENT_DATE,
		due_date TEXT NOT NULL,
		type_id INTEGER NOT NULL,
		FOREIGN KEY(type_id) REFERENCES notice_type(id) ON DELETE SET NULL
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}

func (p *Pool) CreateTableNoticeType() {
	s := `CREATE TABLE IF NOT EXISTS notice_type (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_DATE,
		updated_at DEFAULT CURRENT_DATE,
		name TEXT NOT NULL,
		description TEXT
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}
