package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/post/svc"
)

type PostController struct {
	postService *svc.PostService
	pb.UnimplementedPostServiceServer
}

func NewPostController(service *svc.PostService) *PostController {
	return &PostController{postService: service}
}
