package goodday

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

func (p *Pool) CreateTableSection() {
	s := `CREATE TABLE IF NOT EXISTS sections (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_TIMESTAMP,
		updated_at DEFAULT CURRENT_TIMESTAMP,
		name TEXT NOT NULL,
		description TEXT,
		start_date TEXT,
		due_date TEXT,
		type TEXT NOT NULL CHECK (type="work" or type="personal"),
		status TEXT NOT NULL CHECK (status="on_track" or status="on_hold" or status="complete") DEFAULT "on_track"
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}

func (p *Pool) CreateTableTask() {
	s := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_TIMESTAMP,
		updated_at DEFAULT CURRENT_TIMESTAMP,
		name TEXT NOT NULL,
		description TEXT,
		start_date TEXT,
		due_date TEXT,
		dependency_id INTEGER, 
		type TEXT NOT NULL CHECK (type="epic" or type="feature" or type="bug" or type="action"),
		status TEXT NOT NULL CHECK (status="todo" or status="doing" or status="done" or status="pending"),
		priority TEXT NOT NULL CHECK (priority="P1" or priority="P2" or priority="P3"),
		section_id INTEGER NOT NULL,
		FOREIGN KEY(dependency_id) REFERENCES tasks(id) ON DELETE SET NULL,
		FOREIGN KEY(section_id) REFERENCES sections(id) ON DELETE CASCADE 
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}

func (p *Pool) CreateTableNotice() {
	s := `CREATE TABLE IF NOT EXISTS notices (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DEFAULT CURRENT_TIMESTAMP,
		updated_at DEFAULT CURRENT_TIMESTAMP,
		name TEXT NOT NULL,
		description TEXT,
		start_date DEFAULT CURRENT_DATE,
		due_date TEXT,
		type TEXT NOT NULL CHECK (type="task" or type="note" or type="quote")
	)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec()
}

func (p *Pool) InsertSection(name, t string) {
	s := `INSERT INTO sections (name, type) values(?, ?)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec(name, t)
}

func (p *Pool) InsertNotice(name, t string) {
	s := `INSERT INTO notices (name, type) values(?, ?)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec(name, t)
}

func (p *Pool) InsertTask(name, t, status, priority string, section int) {
	s := `INSERT INTO tasks (name, type, status, priority, section_id) values(?, ?, ?, ?, ?)`
	stmt, _ := p.DB.Prepare(s)
	stmt.Exec(name, t, status, priority, section)
}

type Task struct {
	Name     string
	Kind     string
	Status   string
	Priority string
}

type Section struct {
	Name  string
	Kind  string
	Tasks []Task
}

func (p *Pool) GetSections() []Section {
	q := `SELECT id, name, type FROM sections`
	rows, _ := p.DB.Query(q)
	var ID int
	var Name, Type string
	var sections []Section

	for rows.Next() {
		rows.Scan(&ID, &Name, &Type)
		tasks := p.GetTasks(ID)
		section := Section{Name, Type, tasks}
		sections = append(sections, section)
	}
	return sections
}

func (p *Pool) GetTasks(id int) []Task {
	q := fmt.Sprintf("SELECT name, type, status, priority FROM tasks WHERE section_id = %d", id)

	rows, _ := p.DB.Query(q)
	var n string
	var t string
	var s string
	var pr string
	var tasks []Task
	for rows.Next() {
		rows.Scan(&n, &t, &s, &pr)
		task := Task{
			Name:     n,
			Kind:     t,
			Status:   s,
			Priority: pr,
		}
		tasks = append(tasks, task)
	}
	return tasks
}
