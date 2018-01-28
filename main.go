package main

import (
	"net/http"

	"./handlers"
	"./models"

	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
)

// Server
func main() {

	models.CreateTable()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handlers
	router := httprouter.New()
	router.GET("/", handlers.Home)
	router.GET("/item/:id/", handlers.GetOneTask) //REMEMBER: a req without an ID still routes to this URL, passing in "favicon.ico" as the param. can't be /:id, must be diff branch url/item/:id
	router.DELETE("/item/:id", handlers.DeleteOneTask)
	router.POST("/", handlers.NewTask)

	http.ListenAndServe(":8000", router)
}
