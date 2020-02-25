package main

import (
	"fmt"
	"net/http"

	"../repo"
)

const porta = ":8080"

func init() {
	fmt.Println("Iniciando a aplicação!")
}

func main() {
	err := repo.AbrirConexaoDB()
	if err != nil {
		fmt.Println("Erro ao iniciar conexao com banco de dados! --> ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Teste lalala")
	})

	fmt.Printf("Iniciado servidor na porta %s...\n", porta)
	err = http.ListenAndServe(porta, nil)
	if err != nil {
		fmt.Println("Houve erro --> , " + err.Error())
	}
}
