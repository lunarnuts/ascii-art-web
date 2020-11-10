package response

import (
	"html/template"
	"net/http"
)

type err struct {
	ErrorNumber int
	ErrorData   string
}

type input struct {
	Banner string `json:"banner"`
	Text   string `json:"text"`
}

func ResponseHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		if r.Method == "GET" {
			var data input
			t, e := template.ParseFiles("./templates/index.html")
			if e != nil {
				ErrorHandler(500, w, r)
				return
			}
			t.Execute(w, data)
		} else if r.Method == "POST" {
			PostHandler(w, r)
		} else {
			ErrorHandler(400, w, r)
		}
	} else if r.URL.Path == "/export" {
		w.Header().Set("Content-Disposition", "attachment; filename=output.txt")
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		http.ServeFile(w, r, "./export/output.txt")
	} else if r.URL.Path == "/error500" {
		ErrorHandler(500, w, r)
	} else if r.URL.Path == "/error400" {
		ErrorHandler(400, w, r)
	} else if r.URL.Path == "/error422" {
		ErrorHandler(422, w, r)
	} else {
		ErrorHandler(404, w, r)
	}

}
