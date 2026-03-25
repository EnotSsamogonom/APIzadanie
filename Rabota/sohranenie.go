package Rabota

import (
	DBQuestions2 "APIzadanie/Rabota/DBQuestions"
	"APIzadanie/Rabota/DBQuestions/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDanniiBooks(bok *gin.Context) {
	var knijka types.Books
	err := bok.ShouldBindJSON(&knijka)
	if err != nil {
		bok.JSON(http.StatusBadRequest, gin.H{
			"Ошибка": "Некорректный формат. Ожидается JSON с полями 'nazvanie', 'about', 'authorID'.",
		})
		return
	}
	nazvanie := knijka.Nazvaie
	about := knijka.About
	authorID := knijka.AuthorID
	NewBookDanni, err := DBQuestions2.NewBook(nazvanie, about, authorID)
	if err != nil {
		bok.JSON(http.StatusOK, gin.H{"Ошибка": "Не достаточно данных."})
		return
	}
	err = DBQuestions2.DobavBook(NewBookDanni)
	if err != nil {
		bok.JSON(http.StatusOK, gin.H{"Ошибка": "Не удалось сохранить данные"})
		return
	}
	bok.JSON(http.StatusOK, gin.H{
		"Успешно:": "Данные книги сохранены",
	})
}

func CreateDanniiAuthor(aut *gin.Context) {
	var imia types.Author
	err := aut.ShouldBindJSON(&imia)
	if err != nil {
		aut.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат. Ожидается JSON с полем 'Name'.",
		})
		return
	}

	name := imia.Name
	NewAuthorDanni, err := DBQuestions2.NewAuthor(name)

	if err != nil {
		aut.JSON(http.StatusOK, gin.H{
			"Ошибка": "Нет данных автора",
		})
		return
	}
	err = DBQuestions2.DobavAuthor(NewAuthorDanni)
	if err != nil {
		aut.JSON(http.StatusOK, gin.H{
			"Ошибка": "Не удалось сохранить данные",
		})
		return
	}
	aut.JSON(http.StatusOK, gin.H{
		"Успешно:": "Данные автора сохранены",
	})

}
