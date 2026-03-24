package DBQuestions

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func DeleteAuthor(idAuthor int) error {

	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	_, err = db.Exec(
		context.Background(),
		"DELETE FROM authors WHERE id = ($1)",
		idAuthor)
	return err
}

func DeliteBook(idBook int) error {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	_, err = db.Exec(
		context.Background(),
		"DELETE FROM books WHERE id = ($1)",
		idBook)
	return err
}
