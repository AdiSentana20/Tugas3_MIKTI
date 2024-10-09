package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
    var err error
    db, err = sql.Open("sqlite3", "./items.db")
    if err != nil {
        log.Fatal(err)
    }

    // Buat tabel jika belum ada
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        nama TEXT,
        harga INTEGER
    )`)
    if err != nil {
        log.Fatal(err)
    }
}