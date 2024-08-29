package database

import (
	"context"
	"errors"
	"github.com/victorthecreative/CientServerAPI-FC/internal/models"
	"gorm.io/gorm"
	"time"
)

// CreateTable realiza a migração automática do modelo JsonInternalResponse no banco de dados,
// criando a tabela correspondente, se ainda não existir.
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&models.JsonInternalResponse{})
}

// NewExchange insere uma nova cotação de câmbio no banco de dados.
// A cotação é extraída do contexto, onde foi armazenada anteriormente.
// Retorna um erro se a cotação não puder ser recuperada ou se houver um problema ao inseri-la no banco.
func NewExchange(ctx context.Context, db *gorm.DB) error {

	exchange, ok := ctx.Value(models.ExchangeRateKey).(*models.JsonExternalResponse)
	if !ok || exchange == nil {
		return errors.New("failed to retrieve exchange rate from context")
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	if err := db.WithContext(ctx).Create(&exchange.JsonInternalResponse).Error; err != nil {
		return err
	}

	return nil
}
