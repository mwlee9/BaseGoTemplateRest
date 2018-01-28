package main

import (
	"net/http"

	"../todo/handlers"

	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
)

// Server
func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handlers
	router := httprouter.New()
	router.GET("/", handlers.Home)
	router.GET("/:id", handlers.GetOneTask)
	router.DELETE("/:id", handlers.DeleteOneTask)
	router.POST("/:id", handlers.NewTask)

	http.ListenAndServe(":8000", router)
}
