package main

import (
	"APIzadanie/DBQuestions"
	"APIzadanie/DBQuestions/tables"
	"APIzadanie/Rabota"

	"github.com/gin-gonic/gin"
)

func main() {
	DBQuestions.Connect()
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
	api.PATCH("/knigi/:id", Rabota.IzmenenieAbout)
	api.DELETE("/deliteAutor/:id", Rabota.YdalenieAuthor)
	api.DELETE("/deliteBook/:id", Rabota.YdalenieBook)
	router.Run(":8081")

}
