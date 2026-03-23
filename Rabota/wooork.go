package Rabota

import (
	"ZadanieTest/BDQuestions"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VivodVsexAvtorov(aut *gin.Context) {
	authors, err := BDQuestions.VivodAuthors()
	if err != nil {
		log.Fatalf("Не удалось получить авторов: %v", err)
	}
	if len(authors) == 0 {
		fmt.Println("Список авторов пуст.")
		return
	}
	//
	aut.JSON(http.StatusOK, gin.H{
		"Список авторов:": "",
	})
	for _, author := range authors {
		aut.JSON(http.StatusOK, gin.H{
			"--": author,
		})

	}
}

func VivodAvtora(aut *gin.Context) {

}

func VivodVsexKnig(aut *gin.Context) {

}

func VivodKnigi(knig *gin.Context) {

	fmt.Print("Введите id автора для просмотра книг: ")
	var id BDQuestions.RequestBody
	err := knig.ShouldBindJSON(&id)
	if err != nil {
		knig.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат. Ожидается JSON с полем 'ID'.",
		})
		return
	}

	knigki, err := BDQuestions.VivodBook(id.ID)
	if err != nil {
		knig.JSON(http.StatusOK, gin.H{
			"Данный автор:": "НЕ НАЙДЕН",
		})
		return
	}
	if len(knigki) == 0 {
		knig.JSON(http.StatusOK, gin.H{
			"Список книг:": "ПУСТО",
		})
		return
	}
	knig.JSON(http.StatusOK, gin.H{
		"Список авторов:": "",
	})
	for _, knigi := range knigki {
		knig.JSON(http.StatusOK, gin.H{
			"книга": knigi,
		})
	}

}

func CreateDanniiAuthor(aut *gin.Context) {
	var imia BDQuestions.Author
	err := aut.ShouldBindJSON(&imia)
	if err != nil {
		aut.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат. Ожидается JSON с полем 'Name'.",
		})
		return
	}

	name := imia.Name
	NewAuthorDanni, err := BDQuestions.NewAuthor(name)

	if err != nil {
		aut.JSON(http.StatusOK, gin.H{
			"Ошибка": "Нет данных автора",
		})
		return
	}
	err = BDQuestions.DobavAuthor(NewAuthorDanni)
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

func CreateDanniiBooks(bok *gin.Context) {
	var knijka BDQuestions.Books
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
	NewBookDanni, err := BDQuestions.NewBook(nazvanie, about, authorID)
	if err != nil {
		bok.JSON(http.StatusOK, gin.H{"Ошибка": "Не достаточно данных."})
		return
	}
	err = BDQuestions.DobavBook(NewBookDanni)
	if err != nil {
		bok.JSON(http.StatusOK, gin.H{"Ошибка": "Не удалось сохранить данные"})
		return
	}
	bok.JSON(http.StatusOK, gin.H{
		"Успешно:": "Данные книги сохранены",
	})
}

func YdalenieAuthor(y *gin.Context) {
	var id BDQuestions.RequestBody
	err := y.ShouldBindJSON(&id)
	if err != nil {
		y.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат. Ожидается JSON с полем 'ID'.",
		})
		return
	}
	//err = Sohranenie.DeleteAuthor(id.ID)
	//if err != nil {
	//
	//	fmt.Println("Не удалось удалить Автора")
	//}
	y.JSON(http.StatusOK, gin.H{
		"Успешно:": "Данные автора удалены",
	})

}

func YdalenieBook(y *gin.Context) {
	var id BDQuestions.RequestBody
	err := y.ShouldBindJSON(&id)
	if err != nil {
		y.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный формат. Ожидается JSON с полем 'ID'.",
		})
		return
	}

	//err = Sohranenie.DeliteBook(id.ID)
	//if err != nil {
	//	fmt.Println("Не удалось удалить книгу")
	//}
	y.JSON(http.StatusOK, gin.H{
		"Успешно:": "Данные книги удалены",
	})
}
