package Rabota

import (
	"APIzadanie/DBQuestions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IzmenenieAbout(c *gin.Context) {

	var new string
	var id int
	err := c.ShouldBindJSON(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Ошибка": "Некорректный формат.",
		})
		return
	}
	err = c.ShouldBindJSON(new)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Ошибка": "Некорректный формат.",
		})
		return
	}
	err = DBQuestions.IzmenenimAbout(new, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Ошибка": "Не удалось  изменить данные"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Успешно:": "Данные книги обновлены",
	})
}

func CreateDanniiBooks(bok *gin.Context) {
	var knijka DBQuestions.Books
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
	NewBookDanni, err := DBQuestions.NewBook(nazvanie, about, authorID)
	if err != nil {
		bok.JSON(http.StatusOK, gin.H{"Ошибка": "Не достаточно данных."})
		return
	}
	err = DBQuestions.DobavBook(NewBookDanni)
	if err != nil {
		bok.JSON(http.StatusOK, gin.H{"Ошибка": "Не удалось сохранить данные"})
		return
	}
	bok.JSON(http.StatusOK, gin.H{
		"Успешно:": "Данные книги сохранены",
	})
}

func CreateDanniiAuthor(aut *gin.Context) {
	var imia DBQuestions.Author
	err := aut.ShouldBindJSON(&imia)
	if err != nil {
		aut.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат. Ожидается JSON с полем 'Name'.",
		})
		return
	}

	name := imia.Name
	NewAuthorDanni, err := DBQuestions.NewAuthor(name)

	if err != nil {
		aut.JSON(http.StatusOK, gin.H{
			"Ошибка": "Нет данных автора",
		})
		return
	}
	err = DBQuestions.DobavAuthor(NewAuthorDanni)
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
