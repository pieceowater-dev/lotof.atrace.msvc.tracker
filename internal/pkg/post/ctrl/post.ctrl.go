package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/post/svc"
	"context"
)

type PostController struct {
	postService *svc.PostService
	pb.UnimplementedPostServiceServer
}

func NewPostController(service *svc.PostService) *PostController {
	return &PostController{postService: service}
}

func (p PostController) GetPosts(ctx context.Context, request *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostController) GetPost(ctx context.Context, request *pb.GetPostRequest) (*pb.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostController) CreatePost(ctx context.Context, request *pb.CreatePostRequest) (*pb.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostController) DeletePost(ctx context.Context, request *pb.DeletePostRequest) (*pb.Post, error) {
	//TODO implement me
	panic("implement me")
}
