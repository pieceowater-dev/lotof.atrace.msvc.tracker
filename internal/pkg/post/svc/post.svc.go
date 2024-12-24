package svc

import (
	"app/internal/pkg/post/ent"
	"context"
	"fmt"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type PostService struct {
	db gossiper.Database
}

// NewPostService initializes a new PostService.
func NewPostService(db gossiper.Database) *PostService {
	return &PostService{db: db}
}

// GetPosts retrieves paginated posts from the database.
func (s PostService) GetPosts(ctx context.Context, filter gossiper.Filter[string]) (gossiper.PaginatedResult[ent.Post], error) {
	var posts []ent.Post
	var count int64

	query := s.db.GetDB().Model(&ent.Post{}).Preload("Location")

	// Apply search filters
	if filter.Search != "" {
		search := "%" + filter.Search + "%"
		query = query.Where("title LIKE ? OR description LIKE ?", search, search)
	}

	// Count total records
	if err := query.Count(&count).Error; err != nil {
		return gossiper.PaginatedResult[ent.Post]{}, fmt.Errorf("failed to count posts: %w", err)
	}

	// Apply pagination
	query = query.Offset((filter.Pagination.Page - 1) * filter.Pagination.Length).Limit(filter.Pagination.Length)

	// Apply sorting dynamically
	if field := filter.Sort.Field; field != "" && gossiper.IsFieldValid(&ent.Post{}, field) {
		query = query.Order(fmt.Sprintf("%s %s", gossiper.ToSnakeCase(field), filter.Sort.Direction))
	}

	// Fetch data
	if err := query.Find(&posts).Error; err != nil {
		return gossiper.PaginatedResult[ent.Post]{}, fmt.Errorf("failed to fetch posts: %w", err)
	}

	return gossiper.NewPaginatedResult(posts, int(count)), nil
}

// GetPost retrieves a single post by ID.
func (s PostService) GetPost(ctx context.Context, id int) (ent.Post, error) {
	var post ent.Post
	result := s.db.GetDB().Preload("Location").First(&post, "id = ?", id)
	if result.Error != nil {
		return ent.Post{}, result.Error
	}
	return post, nil
}

// CreatePost adds a new post to the database.
func (s PostService) CreatePost(ctx context.Context, post *ent.Post) (ent.Post, error) {
	if err := s.db.GetDB().Create(post).Error; err != nil {
		return ent.Post{}, err
	}
	return *post, nil
}

// DeletePost removes a post from the database by ID.
func (s PostService) DeletePost(ctx context.Context, id int) (ent.Post, error) {
	var post ent.Post
	result := s.db.GetDB().First(&post, "id = ?", id)
	if result.Error != nil {
		return ent.Post{}, result.Error
	}
	if err := s.db.GetDB().Delete(&post).Error; err != nil {
		return ent.Post{}, err
	}
	return post, nil
}
