package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/Nirss/blog/internal/domain/post"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
	validPost, err := post.NewPostWithID(
		uuid.New(),
		"Краткая история Dell UNIX",
		"Личные воспоминания одного из разработчиков Dell UNIX об истории создания этой системы",
		[]string{"nix", "История IT", "Старое железо"},
		time.Now())
	require.NoError(t, err)

	tests := []struct {
		name      string
		post      post.Post
		wantError error
	}{
		{
			name:      "success",
			post:      validPost,
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf(
				"postgres://%s:%s@%s:%s/%s", // Нужно заполнить данными
				"",                          // user
				"",                          // password
				"",                          // host
				"",                          // port
				"",                          // db
			)

			db, err := sqlx.Connect("pgx", url)
			require.NoError(t, err)

			blogRepo := NewBlogRepo(db)

			err = blogRepo.CreatePost(context.Background(), tt.post)

			assert.Equal(t, tt.wantError, err)
		})
	}
}

func TestGetPosts(t *testing.T) {
	tests := []struct {
		name      string
		want      []post.Post
		wantError error
	}{
		{
			name:      "success",
			want:      []post.Post{},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf(
				"postgres://%s:%s@%s:%s/%s", // Нужно заполнить данными
				"",                          // user
				"",                          // password
				"",                          // host
				"",                          // port
				"",                          // db
			)

			db, err := sqlx.Connect("pgx", url)
			require.NoError(t, err)

			blogRepo := NewBlogRepo(db)

			_, err = blogRepo.GetAllPosts(context.Background())

			assert.Equal(t, tt.wantError, err)
		})
	}
}
