package handle

import (
	"net/http"
	"SHR/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.HtmlRenderer(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.HtmlRenderer(w, "about.page.html")
}
