package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type JsonInternalResponse struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type JsonExternalResponse struct {
	JsonInternalResponse `json:"USDBRL"`
}

func main() {
	//mux := http.NewServeMux()

	http.HandleFunc("/cotacao", handlerGetPrice)

	http.ListenAndServe(":8080", nil)

}

func handlerGetPrice(w http.ResponseWriter, r *http.Request) {

	exchance, err := getPrice()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exchance)
}

func getPrice() (*JsonExternalResponse, error) {
	req, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
	}

	defer req.Body.Close()

	resp, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler o Body: %v\n", err)
	}

	var cotacao JsonExternalResponse
	err = json.Unmarshal(resp, &cotacao)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao converter Json: %v\n", err)
	}

	return &cotacao, nil
}
