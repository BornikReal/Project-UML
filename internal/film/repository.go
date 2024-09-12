package film

import (
	"context"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{pool: pool}
}

func (r *repository) Get(ctx context.Context, id int64) (filmModel, error) {
	var u filmModel
	err := pgxscan.Get(ctx, r.pool, &u, "SELECT * FROM films WHERE id = $1", id)
	if err != nil {
		return filmModel{}, fmt.Errorf("pgxscan.Get: %w", err)
	}
	return u, nil
}

func (r *repository) GetAll(ctx context.Context) ([]filmModel, error) {
	var u []filmModel
	err := pgxscan.Select(ctx, r.pool, &u, "SELECT * FROM films")
	if err != nil {
		return nil, fmt.Errorf("pgxscan.Select: %w", err)
	}
	return u, nil
}

func (r *repository) AddRating(ctx context.Context, filmID, ratingID int64) error {
	res, err := r.pool.Exec(ctx, `UPDATE films SET ratings = array_append(ratings, $1) WHERE id = $2`,
		ratingID, filmID)
	if err != nil {
		return fmt.Errorf("pgx.Exec: %w", err)
	}
	if res.RowsAffected() == 0 {
		return errors.New("film not found")
	}
	return nil
}
