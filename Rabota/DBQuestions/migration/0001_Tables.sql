-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS authors (
 id SERIAL PRIMARY KEY,
 name VARCHAR(255) NOT NULL
    );

CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    nazvanie VARCHAR(255) NOT NULL,
    about TEXT NOT NULL,
    author_id INTEGER REFERENCES authors(id) ON DELETE CASCADE -- ON DELETE CASCADE удалит посты автора, если сам автор удален
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE authors;
DROP TABLE books;
-- +goose StatementEnd