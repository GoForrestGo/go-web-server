package handler

import (
	"net/http"

	"github.com/MrBomber0x001/sample/pkg/render"
)

func Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "home.page.tmpl")
}

func About(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "about.page.tmpl")
}
