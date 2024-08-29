//

package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/victorthecreative/CientServerAPI-FC/internal/database"
	"github.com/victorthecreative/CientServerAPI-FC/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

// NewServer inicializa um novo servidor HTTP escutando na porta 8080.
// Define a rota "/cotacao" que responde com a cotação atual do dólar e a salva no banco de dados.
func NewServer() {
	http.HandleFunc("/cotacao", handlerGetPrice)
	http.ListenAndServe(":8080", nil)
}

// handlerGetPrice é o handler para a rota "/cotacao".
// Ele faz uma chamada para uma API externa para obter a cotação do dólar, retorna a cotação ao cliente,
// e insere a cotação no banco de dados.
func handlerGetPrice(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	exchance, err := getPrice(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx = context.WithValue(r.Context(), models.ExchangeRateKey, exchance)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exchance)

	dsn := "root:root@tcp(localhost:3306)/RateExchange?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.CreateTable(db)
	database.NewExchange(ctx, db)
}

// getPrice realiza uma requisição para uma API externa que fornece a cotação do dólar (USD-BRL).
// Ele retorna a cotação no formato de uma estrutura JSON ou um erro em caso de falha.
func getPrice(ctx context.Context) (*models.JsonExternalResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da resposta: %w", err)
	}

	var cotacao models.JsonExternalResponse
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter JSON: %w", err)
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &cotacao, nil
	}
}
