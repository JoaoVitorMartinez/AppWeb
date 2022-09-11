package controllers

import (
	"App/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.SelectTodosProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)

}

func Novo(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Novo", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))

		if err != nil {
			log.Println(err)
		}

		models.NovoProduto(nome, descricao, preco, quantidade)

	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)

}

func Editar(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idProduto)
	temp.ExecuteTemplate(w, "Editar", produto)
}

func Atualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))

		if err != nil {
			err.Error()

		}
		models.AtualizaProduto(id, nome, descricao, preco, quantidade)
		http.Redirect(w, r, "/", 301)
	}

}
