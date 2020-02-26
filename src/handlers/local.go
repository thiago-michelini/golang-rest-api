package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"../model"
	"../repo"
)

const sqlConsulta = "select country, city, telcode " +
	"from local " +
	"where telcode = ?"

const sqlInsert = "insert into requests " +
	"values(?)"

//Local - manipular do Local
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Não foi enviado um número válido, verifique!", http.StatusBadRequest)
		fmt.Println("Não foi enviado um número válido --> ", err.Error())
		return
	}

	local, err = executarConsultaDB(codigoTelefone, w)
	if err != nil {
		return
	}

	err = construirJSONeAtribuirNaResposta(local, w)
	if err != nil {
		return
	}

	_ = inserirRequestNoDB()
}

func inserirRequestNoDB() (err error) {
	err = nil
	//resultado, err := repo.Db.Exec(sqlInsert, time.Now().Format("02/01/2006 15:04:05"))
	resultado, err := repo.Db.Exec(sqlInsert, time.Now())
	if err != nil {
		fmt.Println("Erro ao incluir request no DB, ", sqlInsert, " - ", err.Error())
		return
	}
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		fmt.Println("Erro ao pegao o nr de linhas afetadas pelo sql --> ", sqlInsert, " - ", err.Error())
		return
	}
	fmt.Println("Sucesso! ", linhasAfetadas, " linha(s) afetada(s)")
	return
}

func construirJSONeAtribuirNaResposta(local model.Local, w http.ResponseWriter) (err error) {
	err = nil
	localJSON, err := json.Marshal(local)
	if err != nil {
		http.Error(w, "Erro ao fzr parse da struct para JSON!", http.StatusInternalServerError)
		fmt.Println("Erro ao fzr parse da struct para JSON --> ", err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, string(localJSON))
	return
}

func executarConsultaDB(codigoTelefone int, w http.ResponseWriter) (local model.Local, err error) {
	err = nil
	linha, err := repo.Db.Queryx(sqlConsulta, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possivel pesquisar esse numero!", http.StatusInternalServerError)
		fmt.Println("Não foi possivel executar a query, sql: ", sqlConsulta, "; erro: ", err.Error())
		return
	}
	for linha.Next() {
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "Erro ao fzr o bind do banco para a struct!", http.StatusInternalServerError)
			fmt.Println("Erro ao fzr o bind do banco para a struct --> ", err.Error())
			return
		}
	}
	return
}
