package repository

import (
	"database/sql"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

// InitDB initializes the database connection and schema
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", os.Getenv("DB_PATH"))
	if err != nil {
		log.Println("DB OPEN ERROR:", err)
		return nil, err
	}

	if err := createTablesAndSeed(db); err != nil {
		log.Println("DB INIT ERROR:", err)
		return nil, err
	}

	return db, nil
}

// createTablesAndSeed creates tables and inserts a secure admin user
func createTablesAndSeed(db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("adminbatman@123"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_admin BOOLEAN DEFAULT 0,
		profile_picture TEXT
	);

	CREATE TABLE IF NOT EXISTS jobs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		company TEXT NOT NULL,
		location TEXT NOT NULL,
		salary TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);

	INSERT OR IGNORE INTO users (username, password, email, is_admin)
	VALUES (?, ?, ?, 1);
	`,
		"adminbatman",
		string(hashedPassword),
		"admibatmann@test.com",
	)

	return err
}
