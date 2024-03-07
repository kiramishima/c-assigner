package interfaces

import (
	"context"
	models "kiramishima/credit_assigner/internal/models"
)

// CreditService interface
type CreditService interface {
	Assign(ctx context.Context, investment int32) (*models.Credit, error)
	Stats(ctx context.Context) (*models.Stats, error)
}
