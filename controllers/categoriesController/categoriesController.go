package categoriescontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/Muhless/go-crud/entities"
	categoriesmodel "github.com/Muhless/go-crud/models/categoriesModel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categoriesmodel.GetALL()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var categories entities.Category

		categories.Name = r.FormValue("name")
		categories.CreatedAt = time.Now()
		categories.UpdatedAt = time.Now()

		if ok := categoriesmodel.Create(categories); !ok {
			temp, _ := template.ParseFiles("views/category/create.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			panic(nil)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category := categoriesmodel.Detail(id)
		data := map[string]any{
			"category": category,
		}

		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var category entities.Category

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()
		if ok := categoriesmodel.Update(id, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

}
