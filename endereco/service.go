package endereco

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Address struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Source     string
	Err        error
}

func FetchFromAPI(url string, source string, ch chan Address) {
	client := http.Client{Timeout: 1 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		ch <- Address{Source: source, Err: err}
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var addr Address
	if err := json.Unmarshal(body, &addr); err != nil {
		ch <- Address{Source: source, Err: err}
		return
	}

	addr.Source = source
	ch <- addr
}
