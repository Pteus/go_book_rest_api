package models

import (
	"github.com/pteus/book_rest_api/db"
)

type Book struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"       binding:"required"`
	Description string `json:"description"`
}

func (b *Book) Save() error {
	query := `
	INSERT INTO books(title, description)
	VALUES (?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(b.Title, b.Description)
	if err != nil {
		return err
	}

	b.ID, err = result.LastInsertId()

	return err
}

func GetBookById(id int64) (*Book, error) {
	query := "SELECT id, title, description FROM books WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Description)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func GetAllBooks() ([]Book, error) {
	query := "SELECT id, title, description FROM books"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Description)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
