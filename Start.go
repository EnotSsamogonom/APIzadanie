package main

import (
	"ZadanieTest/DBQuestions/tables"
	"ZadanieTest/Rabota"

	"github.com/gin-gonic/gin"
)

func main() {
	tables.DobavTableAuthots()
	tables.DobavTableBooks()
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/authory", Rabota.VivodVsexAvtorov) // GET /api/authory
	api.GET("/authory/:id", Rabota.VivodAvtora)
	api.GET("/knigi", Rabota.VivodVsexKnig)
	api.GET("/knigi/:id", Rabota.VivodKnigi)
	api.POST("/dobavAutor", Rabota.CreateDanniiAuthor)
	api.POST("/dobavBook", Rabota.CreateDanniiBooks)
	api.DELETE("/deliteAutor", Rabota.YdalenieAuthor)
	api.DELETE("/deliteBook", Rabota.YdalenieBook)
	router.Run(":8081")

}
