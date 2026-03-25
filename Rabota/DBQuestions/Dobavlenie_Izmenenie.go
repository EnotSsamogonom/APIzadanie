package DBQuestions

import (
	"APIzadanie/Rabota/DBQuestions/types"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func IzmenenimAbout(new string, id int) error {
	var b types.Host
	host := b.Host
	db, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s", host))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close(context.Background())
	_, err = db.Exec(
		context.Background(),
		"UPDATE books SET about = ($1),WHERE id = ($2) ",
		new,
		id)
	return err
}

// !!!!!!!!!!!!!!!!!!
func DobavAuthor(aaa *types.Author) error {
	var b types.Host
	host := b.Host
	db, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s", host))
	//db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
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

func DobavBook(book *types.Books) error {
	var b types.Host
	host := b.Host
	db, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s", host))
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
