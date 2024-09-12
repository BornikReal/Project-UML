package rating

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{pool: pool}
}

func (r *repository) Create(ctx context.Context, rating Rating) (int64, error) {
	var id int64
	err := pgxscan.Get(ctx, r.pool, &id,
		`INSERT INTO ratings(score, review, film_id, user_id, is_special)
						VALUES ($1, $2, $3, $4, $5)
						RETURNING id`,
		rating.Score, rating.Review, rating.FilmID, rating.UserID, rating.IsSpecial)
	if err != nil {
		return 0, fmt.Errorf("pgxscan.Get: %w", err)
	}
	return id, nil
}

func (r *repository) GetByIDs(ctx context.Context, id []int64) ([]ratingModel, error) {
	var u []ratingModel
	err := pgxscan.Select(ctx, r.pool, &u, "SELECT * FROM ratings WHERE id = ANY($1)", pq.Array(id))
	if err != nil {
		return nil, fmt.Errorf("pgxscan.Select: %w", err)
	}
	return u, nil
}
