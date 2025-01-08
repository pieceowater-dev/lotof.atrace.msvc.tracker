package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

type contextKey string

type UserMetadata struct {
	UserID    string
	Namespace string
}

func GetUserMetadata(ctx context.Context) (*UserMetadata, bool) {
	userMetadata, ok := ctx.Value(contextKey("userMetadata")).(*UserMetadata)
	return userMetadata, ok
}

type GrpcMetadata struct {
	contextUserMetadataKey contextKey
}

func NewMetadataMiddleware() *GrpcMetadata {
	return &GrpcMetadata{
		contextUserMetadataKey: "userMetadata",
	}
}

func (md *GrpcMetadata) MiddlewareMethod() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		userMetadata, err := md.getUserMetadataFromContext(ctx)
		if err != nil {
			log.Printf("Failed to extract metadata: %v", err)
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}
		ctx = md.withUserMetadata(ctx, userMetadata)
		return handler(ctx, req)
	}
}

func (md *GrpcMetadata) getUserMetadataFromContext(ctx context.Context) (*UserMetadata, error) {
	incoming, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("no metadata found in context")
	}
	userIDs := incoming.Get("userid")
	namespaces := incoming.Get("namespace")
	if len(userIDs) == 0 || len(namespaces) == 0 {
		return nil, fmt.Errorf("missing metadata keys: userid or namespace")
	}
	return &UserMetadata{
		UserID:    userIDs[0],
		Namespace: namespaces[0],
	}, nil
}

func (md *GrpcMetadata) withUserMetadata(ctx context.Context, userMetadata *UserMetadata) context.Context {
	return context.WithValue(ctx, md.contextUserMetadataKey, userMetadata)
}
