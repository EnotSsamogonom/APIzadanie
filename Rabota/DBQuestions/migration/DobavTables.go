package migration

import (
	"APIzadanie/Rabota/DBQuestions/types"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jackc/pgx/v5"
)

const migrat = "Rabota/DBQuestions/migration"

func Dobav() {
	var dbURL string
	dbURL = fmt.Sprintf("postgres://%s", types.Host)
	db, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Printf("Не удалось подключиться к базе данных (pgx): %v", err)
	}
	defer db.Close(context.Background())

	projectRoot, err := os.Getwd()                      //находим путь в компе к нашему файлу
	migrationsDir := filepath.Join(projectRoot, migrat) //объеденяем с путем до миграции

	goose := "goose"
	cmdArgs := []string{
		"-dir", migrationsDir,
		"postgres",
		dbURL,
		"up",
	}
	cmd := exec.Command(goose, cmdArgs...)

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Критическая ошибка при выполнении миграций 'goose up': %v", err)
	}

	log.Println("Миграции успешно применены (или уже были применены).")
}

//func DobavTableAuthots() {
//	createTableSQLA := `
//	CREATE TABLE IF NOT EXISTS authors (
//		id SERIAL PRIMARY KEY,
//		name VARCHAR(255) NOT NULL
//	);`
//
//	db, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s", types.Host))
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
//		os.Exit(1)
//	}
//	defer db.Close(context.Background())
//
//	// 3. Выполнение запроса
//	_, err = db.Exec(context.Background(), createTableSQLA)
//	if err != nil {
//		log.Fatal("Ошибка создания таблицы:", err)
//	}
//
//	fmt.Println("Таблица 'authors' успешно создана (или уже существует).")
//}
//
//func DobavTableBooks() {
//	createTableSQLB := `
//	CREATE TABLE IF NOT EXISTS books (
//    id SERIAL PRIMARY KEY,
//    nazvanie VARCHAR(255) NOT NULL,
//    about TEXT NOT NULL,
//    author_id INTEGER REFERENCES authors(id) ON DELETE CASCADE -- ON DELETE CASCADE удалит посты автора, если сам автор удален
//);`
//
//	db, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s", types.Host))
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
//		os.Exit(1)
//	}
//	defer db.Close(context.Background())
//
//	_, err = db.Exec(context.Background(), createTableSQLB)
//	if err != nil {
//		log.Fatal("Ошибка создания таблицы:", err)
//	}
//
//	fmt.Println("Таблица 'books' успешно создана (или уже существует).")
//}
