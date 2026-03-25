package Rabota

import (
	DBQuestions2 "APIzadanie/Rabota/DBQuestions"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IzmenenieAbout(c *gin.Context) {

	var new string
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := c.ShouldBindJSON(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Ошибка": "Некорректный формат.",
		})
		return
	}
	err = c.ShouldBindJSON(new) //исправить
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Ошибка": "Некорректный формат.",
		})
		return
	}
	err = DBQuestions2.IzmenenimAbout(new, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Ошибка": "Не удалось  изменить данные"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Успешно:": "Данные книги обновлены",
	})
}
