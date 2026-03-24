package Rabota

import (
	"APIzadanie/DBQuestions"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func VivodVsexAvtorov(aut *gin.Context) {
	id := 0
	authors, err := DBQuestions.VivodAuthors(id)
	if err != nil {
		log.Fatalf("Не удалось получить авторов: %v", err)
	}
	if len(authors) == 0 {
		fmt.Println("Список авторов пуст.")
		return
	}
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
	idStr := aut.Param("id")
	id, _ := strconv.Atoi(idStr) //строку в число переделали

	authors, err := DBQuestions.VivodAuthors(id)
	if err != nil {
		log.Fatalf("Не удалось получить авторов: %v", err)
	}
	if len(authors) == 0 {
		fmt.Println("Список авторов пуст.")
		return
	}
	aut.JSON(http.StatusOK, gin.H{
		"Список авторов:": "",
	})
	for _, author := range authors {
		aut.JSON(http.StatusOK, gin.H{
			"--": author,
		})

	}
}

func VivodVsexKnig(knig *gin.Context) {
	id := 0
	knigki, err := DBQuestions.VivodBook(id)
	if err != nil {
		log.Fatalf("Не удалось получить авторов: %v", err)
	}
	if len(knigki) == 0 {
		fmt.Println("Список книг пуст.")
		return
	}
	knig.JSON(http.StatusOK, gin.H{
		"Список книг:": "",
	})
	for _, knigki := range knigki {
		knig.JSON(http.StatusOK, gin.H{
			"--": knigki,
		})

	}
}

func VivodKnigi(knig *gin.Context) {

	idStr := knig.Param("id")
	id, _ := strconv.Atoi(idStr)
	knigki, err := DBQuestions.VivodBook(id)
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
