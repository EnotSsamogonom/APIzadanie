package Rabota

import (
	"APIzadanie/Rabota/DBQuestions"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func YdalenieAuthor(y *gin.Context) {
	idStr := y.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := DBQuestions.DeleteAuthor(id)
	if err != nil {
		y.JSON(http.StatusBadRequest, gin.H{
			"error": "Не удалось удалить автора",
		})
		return
	}
	y.JSON(http.StatusOK, gin.H{
		"Успешно:": "Данные автора и его книги удалены",
	})

}

func YdalenieBook(y *gin.Context) {
	idStr := y.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := DBQuestions.DeliteBook(id)
	if err != nil {
		y.JSON(http.StatusBadRequest, gin.H{
			"error": "Не удалось удалить книгу",
		})
		return
	}
	y.JSON(http.StatusOK, gin.H{
		"Успешно:": "Данные книги удалены",
	})
}
