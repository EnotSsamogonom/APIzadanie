package DBQuestions

import (
	"errors"
)

type Host struct {
	user     string
	password string
	host     int
	nameDB   string
}

// "postgres://postgres:Password@localhost:1313/BooksDB"
func Connect() *Host {
	ccc := &Host{
		user:     "postgres",
		password: "Password",
		host:     1313,
		nameDB:   "BooksDB",
	}

	return ccc
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Books struct {
	ID       int    `json:"id"`
	Nazvaie  string `json:"nazvanie"`
	About    string `json:"about"`
	AuthorID int    `json:"author_id"` // Связь с автором
}

func NewBook(nazvanie string, about string, author_id int) (*Books, error) {

	if nazvanie == ("") {
		return nil, errors.New("НЕТ НАЗВАНИЯ")
	}

	if author_id == 0 {
		return nil, errors.New("НЕ УКАЗАН ID АВТОРА")
	}
	Vremennii := &Books{
		Nazvaie:  nazvanie,
		About:    about,
		AuthorID: author_id,
	}
	return Vremennii, nil
}

func NewAuthor(name string) (*Author, error) {

	if name == ("") {
		return nil, errors.New("НЕТ Автора")
	}

	Vremennie := &Author{
		Name: name,
	}
	return Vremennie, nil
}
