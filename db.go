package main

import (
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func NewDB(path string) (*DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}

var ErrDuplicateUser error = errors.New("Duplicate User")

func (db *DB) Init() error {
	_, err := db.db.Exec("create table users (id integer primary key autoincrement, name varchar unique, hits integer default 0)")
	if err != nil {
		log.Println("[WARN] db error:", err)
	}
	return err
}

func (db *DB) AddUser(name string) error {
	_, err := db.db.Exec("insert into users (name) values (?)", name)
	if err != nil {
		log.Println("[WARN] db error:", err)
		if strings.Contains(err.Error(), "UNIQUE") {
			return ErrDuplicateUser
		}
	}
	return err
}

func (db *DB) GetUser(id int) (string, error) {
	tx, err := db.db.Begin()
	if err != nil {
		log.Println("[WARN] db error:", err)
		return "", err
	}
	row := tx.QueryRow("select name, hits from users where id = ?", id)

	var name string
	var hits int
	err = row.Scan(&name, &hits)
	if err != nil {
		tx.Rollback()
		log.Println("[WARN] db error:", err)
		return "", err
	}

	_, err = tx.Exec("update users set hits = ? where id = ?", hits+1, id)
	if err != nil {
		tx.Rollback()
		log.Println("[WARN] db error:", err)
		return "", err
	}

	tx.Commit()
	return name, nil
}

func (db *DB) RandomUser() (string, error) {
	i := rand.Int()
	row := db.db.QueryRow("select max(id) from users")
	var max int
	err := row.Scan(&max)
	if err != nil {
		log.Println("[WARN] db error:", err)
		return "", err
	}
	return db.GetUser((i % max) + 1)
}

func (db *DB) Close() error {
	return db.db.Close()
}
