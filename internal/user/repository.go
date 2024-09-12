package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

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

func (r *repository) Create(ctx context.Context, req CreateUserRequest) (int64, error) {
	var id int64
	err := pgxscan.Get(ctx, r.pool, &id, `INSERT INTO users(role, username, profile_description, avatar, email, password)
						VALUES ($1, $2, $3, $4, $5, $6)
						RETURNING id`, req.Role, req.Username, req.ProfileDescription, req.Avatar, req.Email, req.Password)
	if err != nil {
		return 0, fmt.Errorf("pgxscan.Get: %w", err)
	}
	return id, nil
}

func (r *repository) Get(ctx context.Context, id int64) (userModel, error) {
	var u userModel
	err := pgxscan.Get(ctx, r.pool, &u, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return userModel{}, fmt.Errorf("pgxscan.Get: %w", err)
	}
	return u, nil
}

func (r *repository) Update(ctx context.Context, req UpdateUserRequest) error {
	var setQuery strings.Builder
	var args []interface{}

	if req.Role != nil {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, *req.Role)
		setQuery.WriteString(fmt.Sprintf(" role = $%d", len(args)))
	}

	if req.Username != nil {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, *req.Username)
		setQuery.WriteString(fmt.Sprintf(" username = $%d", len(args)))
	}

	if req.ProfileDescription != nil {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, *req.ProfileDescription)
		setQuery.WriteString(fmt.Sprintf(" profile_description = $%d", len(args)))
	}

	if req.Avatar != nil {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, *req.Avatar)
		setQuery.WriteString(fmt.Sprintf(" avatar = $%d", len(args)))
	}

	if req.Email != nil {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, *req.Email)
		setQuery.WriteString(fmt.Sprintf(" email = $%d", len(args)))
	}

	if req.Password != nil {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, *req.Password)
		setQuery.WriteString(fmt.Sprintf(" password = $%d", len(args)))
	}

	if len(req.AddRatings) != 0 {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, pq.Array(req.AddRatings))
		setQuery.WriteString(fmt.Sprintf(" ratings = ratings || $%d", len(args)))
	}

	if len(req.AddPosts) != 0 {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, pq.Array(req.AddPosts))
		setQuery.WriteString(fmt.Sprintf(" posts = posts || $%d", len(args)))
	}

	if len(req.AddComments) != 0 {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, pq.Array(req.AddComments))
		setQuery.WriteString(fmt.Sprintf(" comments = comments || $%d", len(args)))
	}

	if len(req.AddPrivateMessages) != 0 {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, pq.Array(req.AddPrivateMessages))
		setQuery.WriteString(fmt.Sprintf(" private_messages = private_messages || $%d", len(args)))
	}

	if len(req.AddBlackList) != 0 {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, pq.Array(req.AddBlackList))
		setQuery.WriteString(fmt.Sprintf(" black_list = black_list || $%d", len(args)))
	}

	if len(req.AddRestrictions) != 0 {
		if len(args) != 0 {
			setQuery.WriteByte(',')
		}
		args = append(args, pq.Array(req.AddRestrictions))
		setQuery.WriteString(fmt.Sprintf(" restrictions = restrictions || $%d", len(args)))
	}
	if len(args) == 0 {
		return nil
	}

	args = append(args, req.ID)

	res, err := r.pool.Exec(ctx, fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", setQuery.String(), len(args)), args...)
	if err != nil {
		return fmt.Errorf("pool.Exec: %w", err)
	}
	if res.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *repository) Login(ctx context.Context, username, password string) (authDataModel, error) {
	var authData authDataModel
	err := pgxscan.Get(ctx, r.pool, &authData, "SELECT id, role FROM users WHERE username = $1 AND password = $2", username, password)
	if err != nil {
		return authDataModel{}, fmt.Errorf("pgxscan.Get: %w", err)
	}
	return authData, nil
}
