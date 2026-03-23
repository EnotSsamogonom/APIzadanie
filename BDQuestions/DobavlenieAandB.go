package BDQuestions

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func DobavAuthor(aaa *Author) error {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	_, err = db.Exec(
		context.Background(),
		"INSERT INTO authors (name) VALUES ($1)",
		aaa.Name)
	return err
}

func DobavBook(book *Books) error {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	_, err = db.Exec(
		context.Background(),
		"INSERT INTO books (nazvanie, about, author_id) VALUES ($1, $2, $3) ",
		book.Nazvaie,
		book.About,
		book.AuthorID)
	return err
}
