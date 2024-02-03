package homecontroller

import (
	"net/http"
	"html/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, nil)
}
