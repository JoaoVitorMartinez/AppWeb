package models

import (
	"App/data"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SelectTodosProdutos() []Produto {
	db := data.ConectaBanco()
	selectProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &quantidade, &preco)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
	}
	return produtos

}

func NovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := data.ConectaBanco()

	insert, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(nome, descricao, quantidade, preco)

	defer db.Close()
}

func DeletaProduto(idProduto string) {
	db := data.ConectaBanco()

	delete, err := db.Prepare("DELETE FROM produtos WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(idProduto)

	defer db.Close()

}

func EditaProduto(idproduto string) Produto {
	db := data.ConectaBanco()
	produtoRecuperado, err := db.Query("SELECT * FROM produtos WHERE id=$1", idproduto)

	if err != nil {
		panic(err.Error())
	}

	produtoAtualizado := Produto{}

	for produtoRecuperado.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoRecuperado.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}
		produtoAtualizado.Id = id
		produtoAtualizado.Nome = nome
		produtoAtualizado.Descricao = descricao
		produtoAtualizado.Preco = preco
		produtoAtualizado.Quantidade = quantidade

	}
	defer db.Close()
	return produtoAtualizado
}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := data.ConectaBanco()
	atualizarProduto, err := db.Prepare("UPDATE produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}

	atualizarProduto.Exec(nome, descricao, quantidade, preco, id)

	defer db.Close()
}
