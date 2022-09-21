package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func HtmlRenderer(w http.ResponseWriter, tmpl string) {
	temp, _ := template.ParseFiles("./template/" + tmpl)
	err := temp.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", tmpl)
		return
	}
}
