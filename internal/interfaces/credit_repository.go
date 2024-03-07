package interfaces

import (
	"context"
	"kiramishima/credit_assigner/internal/models"
)

// CreditRepository interface
type CreditRepository interface {
	RegisterAssign(ctx context.Context, data *models.Credit) error
	SumUp(ctx context.Context) (*models.Stats, error)
}
