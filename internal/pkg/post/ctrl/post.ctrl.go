package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/post/ent"
	"app/internal/pkg/post/svc"
	"context"
	"github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type PostController struct {
	postService *svc.PostService
	pb.UnimplementedPostServiceServer
}

// NewPostController initializes a new PostController.
func NewPostController(service *svc.PostService) *PostController {
	return &PostController{postService: service}
}

// GetPosts retrieves paginated posts based on filters.
func (p PostController) GetPosts(ctx context.Context, request *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	filter := gossiper.NewFilter[string](
		request.GetSearch(),
		gossiper.NewSort[string](
			request.GetSort().GetField(),
			gossiper.SortDirection(request.GetSort().GetDirection()),
		),
		gossiper.NewPagination(
			int(request.GetPagination().GetPage()),
			int(request.GetPagination().GetLength()),
		),
	)

	paginatedResult, err := p.postService.GetPosts(ctx, filter)
	if err != nil {
		return nil, err
	}

	var posts []*pb.Post
	for _, post := range paginatedResult.Rows {
		posts = append(posts, &pb.Post{
			Id:    uint64(post.ID),
			Title: post.Title,
			Description: func() string {
				if post.Description != nil {
					return *post.Description
				}
				return ""
			}(),
			Location: func() *pb.PostLocation {
				if post.Location != nil {
					return &pb.PostLocation{
						Comment: func() string {
							if post.Location.Comment != nil {
								return *post.Location.Comment
							}
							return ""
						}(),
						Country:   post.Location.Country,
						City:      post.Location.City,
						Address:   post.Location.Address,
						Latitude:  float32(post.Location.Latitude),
						Longitude: float32(post.Location.Longitude),
					}
				}
				return nil
			}(),
		})
	}

	return &pb.GetPostsResponse{
		Posts: posts,
		PaginationInfo: &pb.PaginationInfo{
			Count: int32(paginatedResult.Info.Count),
		},
	}, nil
}

// GetPost retrieves a single post by ID.
func (p PostController) GetPost(ctx context.Context, request *pb.GetPostRequest) (*pb.Post, error) {
	post, err := p.postService.GetPost(ctx, int(request.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Post{
		Id:    uint64(post.ID),
		Title: post.Title,
		Description: func() string {
			if post.Description != nil {
				return *post.Description
			}
			return ""
		}(),
		Location: func() *pb.PostLocation {
			if post.Location != nil {
				return &pb.PostLocation{
					Comment: func() string {
						if post.Location.Comment != nil {
							return *post.Location.Comment
						}
						return ""
					}(),
					Country:   post.Location.Country,
					City:      post.Location.City,
					Address:   post.Location.Address,
					Latitude:  float32(post.Location.Latitude),
					Longitude: float32(post.Location.Longitude),
				}
			}
			return nil
		}(),
	}, nil
}

// CreatePost creates a new post.
func (p PostController) CreatePost(ctx context.Context, request *pb.CreatePostRequest) (*pb.Post, error) {
	post := &ent.Post{
		Title:       request.Title,
		Description: &request.Description,
		Phrase:      request.Phrase,
		Location: &ent.PostLocation{
			Comment:   &request.Location.Comment,
			Country:   request.Location.Country,
			City:      request.Location.City,
			Address:   request.Location.Address,
			Latitude:  float64(request.Location.Latitude),
			Longitude: float64(request.Location.Longitude),
		},
	}

	createdPost, err := p.postService.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return &pb.Post{
		Id:          uint64(createdPost.ID),
		Title:       createdPost.Title,
		Description: *createdPost.Description,
		Location: &pb.PostLocation{
			Comment:   *createdPost.Location.Comment,
			Country:   createdPost.Location.Country,
			City:      createdPost.Location.City,
			Address:   createdPost.Location.Address,
			Latitude:  float32(createdPost.Location.Latitude),
			Longitude: float32(createdPost.Location.Longitude),
		},
	}, nil
}

// DeletePost deletes a post by ID.
func (p PostController) DeletePost(ctx context.Context, request *pb.DeletePostRequest) (*pb.Post, error) {
	deletedPost, err := p.postService.DeletePost(ctx, int(request.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Post{
		Id:    uint64(deletedPost.ID),
		Title: deletedPost.Title,
		Description: func() string {
			if deletedPost.Description != nil {
				return *deletedPost.Description
			}
			return ""
		}(),
		Location: func() *pb.PostLocation {
			if deletedPost.Location != nil {
				return &pb.PostLocation{
					Comment: func() string {
						if deletedPost.Location.Comment != nil {
							return *deletedPost.Location.Comment
						}
						return ""
					}(),
					Country:   deletedPost.Location.Country,
					City:      deletedPost.Location.City,
					Address:   deletedPost.Location.Address,
					Latitude:  float32(deletedPost.Location.Latitude),
					Longitude: float32(deletedPost.Location.Longitude),
				}
			}
			return nil
		}(),
	}, nil
}
