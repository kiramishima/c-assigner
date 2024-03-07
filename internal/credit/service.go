package credit

import (
	"context"
	"errors"
	"go.uber.org/zap"
	impl "kiramishima/credit_assigner/internal/interfaces"
	"kiramishima/credit_assigner/internal/models"
	"time"
)

var _ impl.CreditService = (*service)(nil)

type service struct {
	logger         *zap.SugaredLogger
	repository     impl.CreditRepository
	contextTimeOut time.Duration
}

// NewCreditService creates a new credit service
func NewCreditService(repo impl.CreditRepository, logger *zap.SugaredLogger, timeout time.Duration) *service {
	return &service{
		logger:         logger,
		repository:     repo,
		contextTimeOut: timeout,
	}
}

// Assign service method
func (c service) Assign(ctx context.Context, investment int32) (*models.Credit, error) {
	assigner := models.Assigner{}
	var credit *models.Credit
	// calculate
	inv300, inv500, inv700, err := assigner.Assign(investment)
	cxt, cancel := context.WithTimeout(ctx, c.contextTimeOut)
	defer cancel()

	if err != nil {
		// create the bad record in db
		credit = &models.Credit{Invest: investment, Status: 0}
		err2 := c.repository.RegisterAssign(cxt, credit)
		if err2 != nil {
			// repository error
			return nil, err2
		}
		// assigner error
		return nil, err
	}
	//
	credit = &models.Credit{Invest: investment, Credit300: inv300, Credit500: inv500, Credit700: inv700, Status: 1}
	err = c.repository.RegisterAssign(cxt, credit)
	if err != nil {
		return nil, err
	}

	return credit, nil
}

// Stats service method
func (c service) Stats(ctx context.Context) (*models.Stats, error) {
	cxt, cancel := context.WithTimeout(ctx, c.contextTimeOut)
	defer cancel()

	stats, err := c.repository.SumUp(cxt)
	if err != nil {
		c.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			return nil, errors.New("Request Timeout")
		default:
			return stats, nil
		}
	}

	return stats, nil
}
