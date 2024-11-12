package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func databasecon() (string, error) {
	// Maak de database connectie met omgevingsvariabelen
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	userIDStr := os.Getenv("USER_ID")
	userID, err := strconv.Atoi(userIDStr)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Failed to open database connection:", err)
		return "", fmt.Errorf("failed to open database connection: %w", err)
	}
	defer db.Close()

	// Controleer de verbinding
	err = db.Ping()
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return "", fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Maak een nieuwe sessie-ID aan
	sessionID := uuid.New().String()

	// Sla de sessie-ID op in de database
	err = storeSession(db, sessionID, userID)
	if err != nil {
		log.Fatalf("Failed to store session: %v", err)
		return "", fmt.Errorf("failed to store session: %w", err)
	}

	fmt.Printf("Session stored successfully with ID: %s\n", sessionID)
	log.Println("Database connection successful")

	// Return de sessie-ID
	return sessionID, nil
}

func storeSession(db *sql.DB, sessionID string, userID int) error {
	query := "INSERT INTO sessions (session_id, user_id, created_at) VALUES (?, ?, NOW())"
	_, err := db.Exec(query, sessionID, userID)
	return err
}
