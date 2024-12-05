package db

import( "database/sql"
_ "github.com/lib/pq"

)

// FUNÇÃO PARA CONECTAR COM O BANCO DE DADOS
// retorna ponteiro para banco de dados
func ConectaComBancoDeDados() *sql.DB {
	conexao := "user= postgres dbname= gopher_loja password=kyra host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

