package routes

import (
	"App/controllers"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/novo", controllers.Novo)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/editar", controllers.Editar)
	http.HandleFunc("/atualizar", controllers.Atualizar)

}
