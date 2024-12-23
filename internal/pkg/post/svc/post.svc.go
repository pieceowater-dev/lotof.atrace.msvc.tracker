package svc

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type PostService struct {
	db gossiper.Database
}

func NewPostService(db gossiper.Database) *PostService {
	return &PostService{db: db}
}
