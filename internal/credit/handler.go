package credit

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/unrolled/render"
	"go.uber.org/zap"
	impl "kiramishima/credit_assigner/internal/interfaces"
	"kiramishima/credit_assigner/internal/models"
	httpUtils "kiramishima/credit_assigner/internal/pkg/utils"
	"net/http"
)

var _ impl.CreditHandlers = (*handler)(nil)

// NewCreditHandlers creates a instance of credit handlers
func NewCreditHandlers(r *chi.Mux, logger *zap.SugaredLogger, s impl.CreditService, render *render.Render, validate *validator.Validate) {
	handler := &handler{
		logger:   logger,
		service:  s,
		response: render,
		validate: validate,
	}

	r.Route("/v1/credits", func(r chi.Router) {
		r.Post("/credit-assignment", handler.CreateAssignHandler)
		r.Get("/statistics", handler.GetStatistics)
	})
}

type handler struct {
	logger   *zap.SugaredLogger
	service  impl.CreditService
	response *render.Render
	validate *validator.Validate
}

func (h handler) CreateAssignHandler(w http.ResponseWriter, req *http.Request) {
	var form = &models.CreditPostFormRequest{}

	err := httpUtils.ReadJSON(w, req, &form)

	if err != nil {
		h.logger.Error(err.Error())
		_ = h.response.JSON(w, http.StatusBadRequest, models.ErrorResponse{ErrorMessage: "the request body is invalid or malformed"})
		return
	}

	h.logger.Info(form)
	// Validate data
	err = form.Validate(h.validate)
	if err != nil {
		h.logger.Error(err.Error())
		_ = h.response.JSON(w, http.StatusBadRequest, models.ErrorResponse{ErrorMessage: err.Error()})
		return
	}
	ctx := req.Context()
	credit, err := h.service.Assign(ctx, form.Investment)
	if err != nil {
		h.logger.Error(err.Error())
		_ = h.response.JSON(w, http.StatusBadRequest, models.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	// response
	if err := h.response.JSON(w, http.StatusOK, credit); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, models.ErrorResponse{ErrorMessage: "Internal Server Error"})
		return
	}
}

// GetStatistics stats handler
func (h handler) GetStatistics(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("Statistics")
	ctx := req.Context()

	resp, err := h.service.Stats(ctx)
	if err != nil {
		h.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			_ = h.response.JSON(w, http.StatusGatewayTimeout, models.ErrorResponse{ErrorMessage: "Request Timeout"})
		default:
			if errors.Is(err, errors.New("No hay registros")) {
				_ = h.response.JSON(w, http.StatusOK, models.Stats{})
			} else {
				_ = h.response.JSON(w, http.StatusInternalServerError, models.ErrorResponse{ErrorMessage: "Internal Server Error"})
			}
		}
		return
	}

	if err := h.response.JSON(w, http.StatusOK, resp); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, models.ErrorResponse{ErrorMessage: "Internal Server Error"})
		return
	}
}
