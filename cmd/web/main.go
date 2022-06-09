package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MrBomber0x001/sample/pkg/config"
	handler "github.com/MrBomber0x001/sample/pkg/handlers"
	"github.com/MrBomber0x001/sample/pkg/render"
)

const portnumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	fmt.Println(fmt.Sprintf("Firing on port %s", portnumber))

	_ = http.ListenAndServe(portnumber, nil)

}
