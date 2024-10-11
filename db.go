package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Credential struct {
	Username string
	Password string
}

type EmailServer struct {
	SMTPHost string
	SMTPPort string
}

type Sqlite struct {
	db *sql.DB
}

type Things3 struct {
	Email string
}

func (s *Sqlite) New() error {
	db, err := sql.Open("sqlite3", "./db/things.db")
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Sqlite) CreateTable() error {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS credentials (
		username TEXT PRIMARY KEY,
		password TEXT
	);

	CREATE TABLE IF NOT EXISTS email_server (
		smtp_host TEXT PRIMARY KEY,
		smtp_port TEXT
	);

	CREATE TABLE IF NOT EXISTS things3 (
		email TEXT PRIMARY KEY
	);
	`
	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) InsertCredentials(c *Credential) error {
	sqlStmt := `
	INSERT INTO credentials (username, password) VALUES (?, ?);
	`
	_, err := s.db.Exec(sqlStmt, c.Username, c.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sqlite) GetLastCredential() (*Credential, error) {
	sqlStmt := `
	SELECT * FROM credentials ORDER BY rowid DESC LIMIT 1;
	`
	row := s.db.QueryRow(sqlStmt)
	c := &Credential{}
	err := row.Scan(&c.Username, &c.Password)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Sqlite) InsertEmailServer(es *EmailServer) error {
	sqlStmt := `
	INSERT INTO email_server (smtp_host, smtp_port) VALUES (?, ?);
	`
	_, err := s.db.Exec(sqlStmt, es.SMTPHost, es.SMTPPort)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sqlite) GetLastEmailServer() (*EmailServer, error) {
	sqlStmt := `
	SELECT * FROM email_server ORDER BY rowid DESC LIMIT 1;
	`
	row := s.db.QueryRow(sqlStmt)
	es := &EmailServer{}
	err := row.Scan(&es.SMTPHost, &es.SMTPPort)
	if err != nil {
		return nil, err
	}

	return es, nil
}

func (s *Sqlite) InsertThings3(t *Things3) error {
	sqlStmt := `
	INSERT INTO things3 (email) VALUES (?);
	`
	_, err := s.db.Exec(sqlStmt, t.Email)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sqlite) GetLastThings3() (*Things3, error) {
	sqlStmt := `
	SELECT * FROM things3 ORDER BY rowid DESC LIMIT 1;
	`
	row := s.db.QueryRow(sqlStmt)
	t := &Things3{}
	err := row.Scan(&t.Email)
	if err != nil {
		return nil, err
	}

	return t, nil
}
