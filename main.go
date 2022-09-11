package main

import (
	"App/routes"
	"net/http"
)

func main() {
	routes.Router()
	http.ListenAndServe(":8000", nil)

}
