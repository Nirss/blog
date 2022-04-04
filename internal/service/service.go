package service

import (
	"context"

	"github.com/Nirss/blog/internal/domain/post"

	"github.com/Nirss/blog/internal/repository"
)

type BlogService interface {
	GetAllPosts(ctx context.Context) ([]post.Post, error)
	CreatePost(ctx context.Context, post post.Post) error
}

type BlogRepo interface {
	GetAllPosts(ctx context.Context) ([]post.Post, error)
	CreatePost(ctx context.Context, post post.Post) error
}

type blogServiceImpl struct {
	repo repository.BlogRepo
}

func NewBlogService(repo BlogRepo) BlogService {
	return &blogServiceImpl{repo: repo}
}

func (s *blogServiceImpl) GetAllPosts(ctx context.Context) ([]post.Post, error) {
	return s.repo.GetAllPosts(ctx)
}
func (s *blogServiceImpl) CreatePost(ctx context.Context, post post.Post) error {
	return s.repo.CreatePost(ctx, post)
}
