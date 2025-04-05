package endereco

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchFromAPI_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"cep": "01153-000",
			"logradouro": "Rua Teste",
			"bairro": "Centro",
			"localidade": "São Paulo",
			"uf": "SP"
		}`)
	}))
	defer server.Close()

	ch := make(chan Address, 1)
	go FetchFromAPI(server.URL, "MockAPI", ch)

	select {
	case result := <-ch:
		if result.Err != nil {
			t.Fatalf("Erro inesperado: %v", result.Err)
		}
		if result.CEP != "01153-000" {
			t.Errorf("CEP incorreto. Esperado '01153-000', recebido '%s'", result.CEP)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout inesperado na chamada")
	}
}

func TestFetchFromAPI_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
	}))
	defer server.Close()

	ch := make(chan Address, 1)
	go FetchFromAPI(server.URL, "MockAPI", ch)

	select {
	case result := <-ch:
		if result.Err == nil {
			t.Fatal("Esperava erro de timeout, mas recebeu sucesso")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout não tratado corretamente")
	}
}

func TestFetchFromAPI_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"cep": 01153-000`)
	}))
	defer server.Close()

	ch := make(chan Address, 1)
	go FetchFromAPI(server.URL, "MockAPI", ch)

	select {
	case result := <-ch:
		if result.Err == nil {
			t.Fatal("Esperava erro de parsing JSON")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout não tratado corretamente")
	}
}
