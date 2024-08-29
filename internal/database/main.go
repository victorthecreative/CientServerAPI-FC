package database

import (
	"context"
	"errors"
	"github.com/victorthecreative/CientServerAPI-FC/internal/models"
	"gorm.io/gorm"
	"time"
)

func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&models.JsonInternalResponse{})
}

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
