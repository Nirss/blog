package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/Nirss/blog/internal/repository"
	"github.com/Nirss/blog/internal/service"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"

	httpapi "github.com/Nirss/blog/internal/api/http"
	"golang.org/x/sync/errgroup"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("loading .env file error: %s", err)
	}

	postgresConn := connectToPostgres()
	defer postgresConn.Close()

	blogRepo := repository.NewBlogRepo(postgresConn)
	blogService := service.NewBlogService(blogRepo)

	group, _ := errgroup.WithContext(context.Background())

	router := httpapi.MakeHandlers(blogService)

	group.Go(func() error {
		return http.ListenAndServe(os.Getenv("HTTP_SERVER_URL"), router)
	})

	if err := group.Wait(); err != nil {
		log.Fatalf("http server error: %s", err)
	}
}

func connectToPostgres() *sqlx.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DATABASE"))

	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		panic(fmt.Sprintf("connect to postgres error: %s", err))
	}

	return db
}
