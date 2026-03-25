package DBQuestions

import (
	"APIzadanie/Rabota/DBQuestions/types"
	"errors"
)

// "postgres://postgres:Password@localhost:1313/BooksDB"
func Connect() *types.Host {
	ccc := &types.Host{
		Host: "postgres:Password@localhost:1313/BooksDB",
	}

	return ccc
}

func NewBook(nazvanie string, about string, author_id int) (*types.Books, error) {

	if nazvanie == ("") {
		return nil, errors.New("НЕТ НАЗВАНИЯ")
	}

	if author_id == 0 {
		return nil, errors.New("НЕ УКАЗАН ID АВТОРА")
	}
	Vremennii := &types.Books{
		Nazvaie:  nazvanie,
		About:    about,
		AuthorID: author_id,
	}
	return Vremennii, nil
}

func NewAuthor(name string) (*types.Author, error) {

	if name == ("") {
		return nil, errors.New("НЕТ Автора")
	}

	Vremennie := &types.Author{
		Name: name,
	}
	return Vremennie, nil
}
