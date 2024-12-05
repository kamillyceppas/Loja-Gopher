package models

import "kamilly/src/loja/db"

// CRIANDO ESTRUTURA DE PRODUTO / STRUCT
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// FUNÇÃO PARA CRIAÇÃO DE PRODUTO

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	// Executa a query para buscar todos os produtos
	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	// Estrutura para armazenar os produtos
	p := Produto{}
	produtos := []Produto{}

	// Itera sobre cada linha retornada pela query
	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		// Lê os dados da linha e armazena nas variáveis
		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		// Atribui os valores ao struct Produto
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		// Adiciona o produto ao slice produtos
		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

// ENVIAR PARA CRIAR NOVO PRODUTO NO BANCO DE DADOS
// função que recebe os dados coletados no formulário de new.html e envia para o banco de dados inserir

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	// PREPARANDO O BANCO DE DADOS E VERIFICANDO SE O INSERT ESTÁ CORRETO
	insereDadosNoBanco, err := db.Prepare("INSERT INTO produtos(nome,descricao,preco,quantidade) VALUES($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	//EXECUTANDO A INSERÇÃO DOS DADOS
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()

}

// FUNÇÃO PARA DELETAR PRODUTO
func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()
	// PREPARANDO O SCRIPT SQL
	deletarOProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error()) // exibir erro caso haja
	}
	// EXECUTANDO A EXCLUSÃO DO PRODUTO
	deletarOProduto.Exec(id)
	defer db.Close()
}

// FUNÇÃO PARA EDITAR PRODUTO
func EditaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	produtoDoBanco, err := db.Query("SELECT * FROM produtos WHERE id =$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}
	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade

	}
	defer db.Close()
	return produtoParaAtualizar
}

// FUNÇÃO PARA ATUALIZAR PRODUTO
func AtualizaProduto(id string, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	atualizaProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
