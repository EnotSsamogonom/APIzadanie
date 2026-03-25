package types

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

const Host = "postgres:Password@localhost:1313/API"
