package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"../models"

	"github.com/julienschmidt/httprouter"
)

// Types - Remember, names must be capital to be exported for the json package to use.

type animal struct {
	ID      string
	Name    string
	Species string
}

var animals []animal

// Home ...
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rows := models.GetAllTasks()

	tempAnimal := animal{ID: "", Name: "", Species: ""}

	for rows.Next() {
		rows.Scan(&tempAnimal.ID, &tempAnimal.Name, &tempAnimal.Species)
		fmt.Println(tempAnimal.ID, tempAnimal.Name, tempAnimal.Species)

	}

	t, err := template.ParseFiles("views/home.html")
	checkErr(err)

	t.Execute(w, "Home")

}

// GetOneTask ...
func GetOneTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	params := ps.ByName("id")

	tempAnimal := animal{ID: "", Name: "", Species: ""}

	rows := models.GetOneTask(params)
	for rows.Next() {
		rows.Scan(&tempAnimal.ID, &tempAnimal.Name, &tempAnimal.Species)
		fmt.Println(tempAnimal.ID, tempAnimal.Name, tempAnimal.Species)

	}

}

// DeleteTask ...
func DeleteOneTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	models.DeleteOneTask(ps.ByName("id"))

}

// NewTask ...
func NewTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
