package BDQuestions

import (
	"errors"
)

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type RequestBody struct {
	ID int `json:"user_id"`
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
