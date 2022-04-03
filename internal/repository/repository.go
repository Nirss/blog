package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"

	"github.com/Nirss/blog/internal/domain/post"
)

type BlogRepo interface {
	GetAllPosts(ctx context.Context) ([]post.Post, error)
	CreatePost(ctx context.Context, post post.Post) error
}

type blogRepoImpl struct {
	db *sqlx.DB
}

func newBlogRepoImpl(db *sqlx.DB) BlogRepo {
	database := &blogRepoImpl{
		db: db,
	}

	return database
}

func NewBlogRepo(db *sqlx.DB) BlogRepo {
	return newBlogRepoImpl(db)
}

type PostDTO struct {
	Id          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Text        string    `db:"text"`
	Tags        string    `db:"tags"`
	CreatedDate time.Time `db:"created_at"`
}

func (r *blogRepoImpl) GetAllPosts(ctx context.Context) ([]post.Post, error) {
	sql := `SELECT * FROM posts`

	var result []PostDTO
	err := sqlx.SelectContext(ctx, r.db, &result, sql)
	if err != nil {
		return nil, err
	}

	posts, err := transformPosts(result)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *blogRepoImpl) CreatePost(ctx context.Context, post post.Post) error {
	sql := `INSERT INTO posts(id, title, text, tags, created_at)
					values($1, $2, $3, $4, $5);`

	if _, err := r.db.ExecContext(ctx, sql, post.GetId(), post.GetTitle(), post.GetText(), post.GetTags(), post.GetCreatedDate()); err != nil {
		return err
	}

	return nil
}
