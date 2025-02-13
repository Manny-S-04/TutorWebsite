package main

import (
	"database/sql"
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"tutor/models"

	"github.com/labstack/echo/v4"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type DB struct {
	db *sql.DB
}

func ConnectDatabase(e *echo.Echo) DB {
	var DbUrl = os.Getenv("DB_URL")
	var DbAuthToken = os.Getenv("DB_AUTH")

	url := fmt.Sprintf("%s?authToken=%s", DbUrl, DbAuthToken)

	db, err := sql.Open("libsql", url)
	if err != nil {
		panic(err)
	}

	return DB{db: db}
}

func (Db *DB) GetReviews() ([]models.Review, error) {
	stmt := `SELECT name, body, stars FROM Reviews WHERE status = 'approved'`
	rows, err := Db.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review

	for rows.Next() {
		var review models.Review

		if err := rows.Scan(&review.Name, &review.Body, &review.Stars); err != nil {
			return nil, err
		}

		reviews = append(reviews, review)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return reviews, nil
}

func (Db *DB) CreatePendingReview(r models.Review) error {
	stmt := `INSERT INTO Reviews (name, body, stars) VALUES (?,?,?)`
	_, err := Db.db.Exec(stmt, r.Name, r.Body, r.Stars)
	if err != nil {
		return err
	}

    EmailService("New Pending Review", fmt.Sprintf("Approve: %s", r.Name))

	return nil
}

// for admin

func (Db *DB) GetAllReviews() ([]models.Review, error) {
	stmt := `SELECT name, body, stars FROM Reviews`
	rows, err := Db.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review

	for rows.Next() {
		var review models.Review

		if err := rows.Scan(&review.Name, &review.Body, &review.Stars); err != nil {
			return nil, err
		}

		reviews = append(reviews, review)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return reviews, nil
}

func (Db *DB) ApproveReview(r models.Review) error {
	stmt := `UPDATE Reviews SET status = 'approved' WHERE name = ? AND body = ?`
	_, err := Db.db.Exec(stmt, r.Name, r.Body)
	if err != nil {
		return err
	}

	return nil
}

func EmailService(subject string, message string) {
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASS")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	raw := `Subject: {subject}
    Content-Type: text/plain; charset="UTF-8"

    {message}
`
	raw = strings.Replace(raw, "{subject}", subject, -1)
	raw = strings.Replace(raw, "{message}", message, -1)

	body := []byte(raw)

	auth := smtp.PlainAuth("", email, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{email}, body)
	if err != nil {
		panic(err)
	}
}
