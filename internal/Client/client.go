package main

import (
	"context"
	"encoding/json"
	"github.com/victorthecreative/CientServerAPI-FC/internal/models"
	"io"
	"net/http"
	"os"
	"text/template"
	"time"
)

type Bid struct {
	Bid string `json:"bid"`
}

func main() {
	Client()
}

func Client() {
	var cotacao models.JsonExternalResponse

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		panic(err)
	}

	bid := cotacao.JsonInternalResponse.Bid

	writeCotacaoTXT(bid)

}

func writeCotacaoTXT(bid string) {

	dolar := Bid{Bid: bid}

	templateText := "Dólar: {{.Bid}}"

	temp := template.Must(template.New("cotacao").Parse(templateText))

	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = temp.Execute(file, dolar)
	if err != nil {
		panic(err)
	}
}
