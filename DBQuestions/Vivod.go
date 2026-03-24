package DBQuestions

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

//func Connect () ( *Host) {
//	ccc := &Host{
//		user: "postgres",
//		password : "Password",
//		host : 1313,
//		nameDB : "BooksDB",
//	}
//
//	return ccc
//}

func VivodAuthors(id int) ([]Author, error) {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())
	if id != 0 {
		aut, err := db.Query(
			context.Background(),
			"SELECT id, name FROM authors WHERE id = $1",
			id)
		if err != nil {
			return nil, err
		}
		defer aut.Close()
		var author []Author

		for aut.Next() { //проходиммся по строкам и собираем их в одно
			var a Author
			err := aut.Scan(&a.ID, &a.Name)
			if err != nil {
				return nil, err
			}
			author = append(author, a)
		}

		if err = aut.Err(); err != nil {
			return nil, err
		}

		return author, nil
	} else {
		aut, err := db.Query(
			context.Background(),
			"SELECT id, name FROM authors ")
		if err != nil {
			return nil, err
		}
		defer aut.Close()
		var author []Author

		for aut.Next() { //проходиммся по строкам и собираем их в одно
			var a Author
			err := aut.Scan(&a.ID, &a.Name)
			if err != nil {
				return nil, err
			}
			author = append(author, a)
		}

		if err = aut.Err(); err != nil {
			return nil, err
		}

		return author, nil
	}

}

func VivodBook(id int) ([]Books, error) {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())
	//подключилтсь
	if id == 0 {
		kn, err := db.Query(
			context.Background(),
			"SELECT id, nazvanie, about, author_id FROM books ",
		)
		if err != nil {
			return nil, err
		}
		defer kn.Close()
		var knigi []Books
		//
		for kn.Next() { //проходиммся по строкам и собираем их в одно
			var b Books
			err := kn.Scan(&b.ID, &b.Nazvaie, &b.About, &b.AuthorID)
			if err != nil {
				return nil, err
			}
			knigi = append(knigi, b)
		}

		if err = kn.Err(); err != nil {
			return nil, err
		}
		return knigi, nil
	} else {
		kn, err := db.Query(
			context.Background(),
			"SELECT id, nazvanie, about, author_id FROM books WHERE author_id = ($1)",
			id)
		if err != nil {
			return nil, err
		}
		defer kn.Close()
		var knigi []Books
		//
		for kn.Next() { //проходиммся по строкам и собираем их в одно
			var b Books
			err := kn.Scan(&b.ID, &b.Nazvaie, &b.About, &b.AuthorID)
			if err != nil {
				return nil, err
			}
			knigi = append(knigi, b)
		}

		if err = kn.Err(); err != nil {
			return nil, err
		}
		return knigi, nil
	}

}
