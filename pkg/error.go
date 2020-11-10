package response

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorHandler(num int, w http.ResponseWriter, r *http.Request) {

	var text string
	switch num {
	case 422:
		text = "Unprocessable entity."
	case 404:
		text = "Not found."
	case 400:
		text = "Bad Request."
	case 500:
		text = "Internal Server Error."
	}
	fmt.Println(text)
	t, e := template.ParseFiles("./templates/error.html")
	if e != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.WriteHeader(num)
	t.Execute(w, err{ErrorNumber: num, ErrorData: text})
}
