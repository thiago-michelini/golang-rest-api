package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../model"
	"../repo"
)

//Local - manipular do Local
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Não foi enviado um número válido, verifique!", http.StatusBadRequest)
		fmt.Println("Não foi enviado um número válido --> ", err.Error())
		return
	}

	sql := "select country, city, telcode from local where telcode = ?"
	linha, err := repo.Db.Queryx(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possivel pesquisar esse numero!", http.StatusInternalServerError)
		fmt.Println("Não foi possivel executar a query, sql: ", sql, "; erro: ", err.Error())
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

	localJSON, err := json.Marshal(local)
	if err != nil {
		http.Error(w, "Erro ao fzr parse da struct para JSON!", http.StatusInternalServerError)
		fmt.Println("Erro ao fzr parse da struct para JSON --> ", err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, string(localJSON))
}
