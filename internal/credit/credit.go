package credit

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/unrolled/render"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"kiramishima/credit_assigner/internal/models"
	"time"
)

// Module credit
var Module = fx.Module("credit",
	fx.Invoke(func(conn *sqlx.DB, logger *zap.SugaredLogger, cfg *models.Configuration, r *chi.Mux, render *render.Render, validate *validator.Validate) error {
		// loads repository
		var repo = NewCreditRepository(conn, logger)
		logger.Info(repo)
		// loads service
		var svc = NewCreditService(repo, logger, time.Duration(cfg.ContextTimeout)*time.Second)
		// loads handlers
		NewCreditHandlers(r, logger, svc, render, validate)
		return nil
	}),
)
