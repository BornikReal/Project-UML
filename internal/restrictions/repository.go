package restrictions

import (
	"context"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{pool: pool}
}

func (r *repository) GetObjectRestrictions(ctx context.Context, req GetObjectRestrictions) ([]restrictionModel, error) {
	query := `SELECT * FROM restrictions
			WHERE restriction_on = $1 AND (valid_until IS NULL OR valid_until < $2)`
	args := []interface{}{req.RestrictionOn, time.Now().UTC()}

	if req.ObjectID != nil {
		args = append(args, *req.ObjectID)
		query += fmt.Sprintf(" AND object_id = $%d", len(args))
	}
	if req.Type != nil {
		args = append(args, *req.Type)
		query += fmt.Sprintf(" AND restriction_type = $%d", len(args))
	}

	var restrictions []restrictionModel
	err := pgxscan.Select(ctx, r.pool, &restrictions, query, args...)
	if err != nil {
		return nil, fmt.Errorf("pgxscan.Select: %w", err)
	}
	return restrictions, nil
}

func (r *repository) AddRestriction(ctx context.Context, restriction Restriction) (int64, error) {
	var id int64
	err := pgxscan.Get(ctx, r.pool, &id,
		`INSERT INTO restrictions(valid_until, restriction_type, object_id, restriction_on)
						VALUES ($1, $2, $3, $4)
						RETURNING id`,
		restriction.ValidUntil, restriction.Type, restriction.ObjectID, restriction.RestrictionOn)
	if err != nil {
		return 0, fmt.Errorf("pgxscan.Get: %w", err)
	}
	return id, nil
}

func (r *repository) DeleteRestriction(ctx context.Context, id int64) error {
	res, err := r.pool.Exec(ctx, "DELETE FROM restrictions WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("pgx.Exec: %w", err)
	}
	if res.RowsAffected() == 0 {
		return errors.New("restriction not found")
	}
	return nil
}
