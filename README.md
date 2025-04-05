# Desafio 02 - Go Expert 🚀

[![Go](https://img.shields.io/badge/Go-1.22-blue.svg)](https://golang.org)
[![Build](https://img.shields.io/badge/build-passing-brightgreen.svg)](#)
[![License](https://img.shields.io/badge/license-Educational-informational)](#)

> Projeto desenvolvido como parte do curso **Go Expert** da Full Cycle.  
> O objetivo é realizar requisições simultâneas a duas APIs de CEP e utilizar a resposta mais rápida.

---

## 💡 Descrição

Este projeto realiza consultas simultâneas a duas APIs distintas de CEP:

- [BrasilAPI](https://brasilapi.com.br/)
- [ViaCEP](https://viacep.com.br/)

A aplicação escolhe automaticamente a **resposta mais rápida**, descartando a mais lenta ou com erro.  
Também há um **timeout de 1 segundo**, para garantir a performance da aplicação.

---

## 🛠️ Tecnologias

- [Go (Golang)](https://golang.org/)
- `net/http`
- `goroutines` e `channels`
- `httptest` (para testes unitários)

---

## ▶️ Como rodar

1. Clone o projeto:
   ```bash
   git clone https://github.com/mersonff/desafio-02-go-expert.git
   cd desafio-02-go-expert

2. Instale as dependências:
   ```bash
   go mod tidy

3. Execute o programa:
   ```bash
   go run main.go

4. Digite um CEP válido:
   ```bash
   Digite o CEP: 01153000

5. Exemplo de Saída:
   ```bash
   Digite o CEP: 01153000
   Resultado recebido da ViaCEP:
   CEP: 01153-000
   Logradouro: Rua Exemplo
   Bairro: Centro
   Cidade: São Paulo
   UF: SP

---

## 🧪 Testes

```bash
go test ./... -v
