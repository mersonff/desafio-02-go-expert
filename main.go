package main

import (
	"fmt"
	"github.com/mersonff/desafio-02-go-expert/endereco"
	"strings"
	"time"
)

func main() {
	var cep string
	fmt.Print("Digite o CEP: ")
	fmt.Scanln(&cep)

	cep = strings.ReplaceAll(cep, "-", "")
	cep = strings.TrimSpace(cep)

	ch := make(chan endereco.Address, 2)

	urlBrasilAPI := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	urlViaCEP := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	go endereco.FetchFromAPI(urlBrasilAPI, "BrasilAPI", ch)
	go endereco.FetchFromAPI(urlViaCEP, "ViaCEP", ch)

	select {
	case result := <-ch:
		if result.Err != nil {
			fmt.Printf("Erro na requisição da %s: %v\n", result.Source, result.Err)
		} else {
			fmt.Printf("Resultado recebido da %s:\n", result.Source)
			fmt.Printf("CEP: %s\nLogradouro: %s\nBairro: %s\nCidade: %s\nUF: %s\n",
				result.CEP, result.Logradouro, result.Bairro, result.Localidade, result.UF)
		}
	case <-time.After(1 * time.Second):
		fmt.Println("Erro: timeout. Nenhuma API respondeu em 1 segundo.")
	}
}
