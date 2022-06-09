package main

import (
	"fmt"
	"net/http"

	handler "github.com/MrBomber0x001/sample/pkg/handlers"
)

const portnumber = ":8080"

func main() {

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	fmt.Println(fmt.Sprintf("Firing on port %s", portnumber))

	_ = http.ListenAndServe(portnumber, nil)

}
