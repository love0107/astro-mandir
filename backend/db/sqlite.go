package db

import (
	"database/sql"
	"log"

	generated "github.com/love0107/astro-mandir/db/generated"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var Queries *generated.Queries

func InitDB() {
	var err error

	// SQLite auto creates this file if not exists
	DB, err = sql.Open("sqlite3", "./astroMandir.db")
	if err != nil {
		log.Fatal("Database error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connect error:", err)
	}

	// Run schema — create all tables
	runSchema()

	// Initialize SQLC queries
	Queries = generated.New(DB)

	log.Println("Database ready!")
}

func runSchema() {
	schema := `
	CREATE TABLE IF NOT EXISTS panchang (
		date TEXT PRIMARY KEY,
		vrat TEXT,
		tithi TEXT,
		nakshatra TEXT,
		sunrise TEXT,
		sunset TEXT,
		muhurat TEXT,
		festival TEXT
	);

	CREATE TABLE IF NOT EXISTS bhajans (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		youtube_id TEXT NOT NULL,
		festival_type TEXT,
		rashi TEXT,
		scheduled_date TEXT
	);

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		phone TEXT UNIQUE,
		rashi TEXT,
		name TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS kundali_requests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		dob TEXT,
		tob TEXT,
		place TEXT,
		rashi TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(schema)
	if err != nil {
		log.Fatal("Schema error:", err)
	}

	log.Println("Tables ready!")
}
