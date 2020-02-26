package repo

import (
	"github.com/jmoiron/sqlx"
	//github.com/lib/pq nao eh usado diretamente pela aplicacao
	_ "github.com/nakagami/firebirdsql"
)

//Db armazena a conexao com banco de dados
var Db *sqlx.DB

//AbrirConexaoDB abre a conexao com o banco de bados
func AbrirConexaoDB() (err error) {
	err = nil
	Db, err = sqlx.Open("firebirdsql", "SYSDBA:masterkey@127.0.0.1/home/thiago/Documentos/golang/thiago-michelini/golang-rest-api/golang.fdb")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
