package migration

import (
	"APIzadanie/Rabota/DBQuestions/types"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func DobavTableAuthots() {
	createTableSQLA := `
	CREATE TABLE IF NOT EXISTS authors (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	);`

	var b types.Host
	host := b.Host
	db, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s", host))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	// 3. Выполнение запроса
	_, err = db.Exec(context.Background(), createTableSQLA)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}

	fmt.Println("Таблица 'authors' успешно создана (или уже существует).")
}

func DobavTableBooks() {
	createTableSQLB := `
	CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    nazvanie VARCHAR(255) NOT NULL,
    about TEXT NOT NULL,
    author_id INTEGER REFERENCES authors(id) ON DELETE CASCADE -- ON DELETE CASCADE удалит посты автора, если сам автор удален
);`
	var b types.Host
	host := b.Host
	db, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s", host))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	_, err = db.Exec(context.Background(), createTableSQLB)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}

	fmt.Println("Таблица 'books' успешно создана (или уже существует).")
}
