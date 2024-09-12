package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"project4/internal/film"
	"project4/internal/rating"
	"project4/internal/restrictions"
	"project4/internal/server"
	"project4/internal/user"
	"project4/pkg/logger"
	"sync"
)

func main() {
	logger.InitLogger()
	logger.Info("init service")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
	}
	defer pool.Close()

	userRep := user.NewRepository(pool)
	userService := user.NewService(userRep)

	filmRep := film.NewRepository(pool)
	filmService := film.NewService(filmRep)

	restrictionRep := restrictions.NewRepository(pool)
	restrictionService := restrictions.NewService(restrictionRep)

	ratingRep := rating.NewRepository(pool)
	ratingService := rating.NewService(ratingRep)

	impl := server.NewImplementation(userService, filmService, restrictionService, ratingService)

	initGrpc(impl)
	initHttp(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
