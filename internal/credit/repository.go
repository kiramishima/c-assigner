package credit

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"kiramishima/credit_assigner/internal/interfaces"
	"kiramishima/credit_assigner/internal/models"
)

// implement credit repository
var _ interfaces.CreditRepository = (*repository)(nil)

// Repository struct
type repository struct {
	db  *sqlx.DB
	log *zap.SugaredLogger
}

// NewCreditRepository Creates a new instance of Repository
func NewCreditRepository(conn *sqlx.DB, logger *zap.SugaredLogger) *repository {
	return &repository{
		db:  conn,
		log: logger,
	}
}

// RegisterAssign save success or fail investments
func (r repository) RegisterAssign(ctx context.Context, data *models.Credit) error {

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		r.log.Error("[Repository] - ", err)
		return errors.New("Error al iniciar la transacción")
	}

	var query = `INSERT INTO credit_assigns(invest, credit_300, credit_500, credit_700, status) 
		VALUES($1, $2, $3, $4, $5)`

	r.log.Info(data)

	_, err = r.db.ExecContext(ctx, query, data.Invest, data.Credit300, data.Credit500, data.Credit700, data.Status)

	if err != nil {
		r.log.Error("[Repository] - ", err)
		_ = tx.Rollback()
		return errors.New("Error al ejecutar el insert")
	}
	// commit
	err = tx.Commit()
	if err != nil {
		// _ = tx.Rollback()
		r.log.Error("[Repository] - ", err)
		return errors.New("Error durante el commit")
	}
	// all good :D
	return nil
}

func (r repository) SumUp(ctx context.Context) (*models.Stats, error) {
	var query = `SELECT total,
		total_sucess, 
		total_fails, 
		avg_total_success_inv,
		avg_total_fail_inv 
	FROM statistics`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.New("Ocurrio un error al preparar la consulta")
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Error(err.Error())
		}
	}(stmt)

	u := &models.Stats{}

	row := stmt.QueryRowContext(ctx)
	err = row.Scan(&u.TotalAssigns, &u.TotalSuccessAssigns, &u.TotalFailAssigns, &u.AVGSuccessAssigns, &u.AVGFailAssigns)
	if errors.Is(err, sql.ErrNoRows) {
		return u, errors.New("No hay registros")
	}

	return u, nil
}
