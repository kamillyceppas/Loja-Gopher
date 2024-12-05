package controllers

import (
	"html/template"
	"kamilly/src/loja/models"
	"log"
	"net/http"
	"strconv"
)

// VARIAVEL PARA VISUALIZAR TODOS OS TEMPLATES HTML
var temp = template.Must(template.ParseGlob("templates/*.html"))

// FUNÇÃO PARA ABRIR A PÁGINA INICIAL DA LOJA
func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}

// FUNÇÃO PARA ABRIR A PÁGINA DE CADASTRAR PRODUTO DA LOJA
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

//FUNÇÃO PARA PEGAR OS DADOS DO FORMULÁRIO

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		//CONVERTER STRINGS PARA OS VALORES CERTOS
		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}
		models.CriarNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}

// FUNÇÃO PARA DELETAR PRODUTO

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)

}

// FUNÇÃO PARA EDITAR PRODUTO

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

// FUNÇÃO PARA ATUALIZAR PRODUTO

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidoParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do id para int:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço para float:", err)
		}
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade para int:", err)
		}
		models.AtualizaProduto(idConvertidoParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
		http.Redirect(w, r, "/", 301)

	}
}
