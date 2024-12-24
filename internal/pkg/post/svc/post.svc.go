package svc

import (
	"app/internal/pkg/post/ent"
	"context"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type PostService struct {
	db gossiper.Database
}

func NewPostService(db gossiper.Database) *PostService {
	return &PostService{db: db}
}

func (s PostService) GetPosts(ctx context.Context, filter gossiper.Filter[string]) (gossiper.PaginatedResult[*ent.Post], error) {
	//TODO implement me
	panic("implement me")
}

func (s PostService) GetPost(ctx context.Context, id int) (ent.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s PostService) CreatePost(ctx context.Context, post *ent.Post) (ent.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s PostService) DeletePost(ctx context.Context, id int) (ent.Post, error) {
	//TODO implement me
	panic("implement me")
}
