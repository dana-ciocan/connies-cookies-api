package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Cookie represents the cookie structure for database operations
type Cookie struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Description string  `json:"description"`
	Price  float64 `json:"price"`
}

// InitDB initializes the SQLite database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./cookies.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected to SQLite database")

	// Create the cookies table if it doesn't exist
	createTable()
	
	// Insert sample data if the table is empty
	insertSampleData()
}

// createTable creates the cookies table
func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS cookies (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		price REAL NOT NULL
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

// insertSampleData adds sample cookies if the table is empty
func insertSampleData() {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM cookies").Scan(&count)
	if err != nil {
		log.Fatal("Failed to count cookies:", err)
	}

	if count == 0 {
		sampleCookies := []Cookie{
			{ID: "1", Name: "Chocolate chip cookie", Description: "A delicious, crunchy chocolate chip cookie.", Price: 2.00},
			{ID: "2", Name: "Oatmeal raisin cookie", Description: "A chewy oatmeal raisin cookie with a hint of cinnamon.", Price: 1.75},
			{ID: "3", Name: "Peanut butter cookie", Description: "A soft peanut butter cookie with a rich flavour.", Price: 2.25},
		}

		for _, cookie := range sampleCookies {
			err := CreateCookie(cookie)
			if err != nil {
				log.Printf("Failed to insert sample cookie %s: %v", cookie.ID, err)
			}
		}
		log.Println("Sample data inserted")
	}
}

// GetAllCookies retrieves all cookies from the database
func GetAllCookies() ([]Cookie, error) {
	rows, err := DB.Query("SELECT id, name, description, price FROM cookies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cookies []Cookie
	for rows.Next() {
		var cookie Cookie
		if err := rows.Scan(&cookie.ID, &cookie.Name, &cookie.Description, &cookie.Price); err != nil {
			return nil, err
		}
		cookies = append(cookies, cookie)
	}

	return cookies, nil
}

// GetCookieByID retrieves a specific cookie by ID
func GetCookieByID(id string) (*Cookie, error) {
	var cookie Cookie
	err := DB.QueryRow("SELECT id, name, description, price FROM cookies WHERE id = ?", id).
		Scan(&cookie.ID, &cookie.Name, &cookie.Description, &cookie.Price)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Cookie not found
		}
		return nil, err
	}

	return &cookie, nil
}

// CreateCookie adds a new cookie to the database
func CreateCookie(cookie Cookie) error {
	query := "INSERT INTO cookies (id, name, description, price) VALUES (?, ?, ?, ?)"
	_, err := DB.Exec(query, cookie.ID, cookie.Name, cookie.Description, cookie.Price)
	return err
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}