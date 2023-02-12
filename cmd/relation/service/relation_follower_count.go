package service

import (
	"Simple-Douyin/cmd/relation/dal/db"
	"Simple-Douyin/kitex_gen/relation"
	"context"
)

type RelationFollowerCountService struct {
	ctx context.Context
}

// NewRelationFollowerCountService new RelationFollowerCountService
func NewRelationFollowerCountService(ctx context.Context) *RelationFollowerCountService {
	return &RelationFollowerCountService{ctx: ctx}
}

// RelationFollowerCount return number of users who are followed
func (s *RelationFollowerCountService) RelationFollowerCount(req *relation.RelationFollowerCountRequest) (int64, error) {
	var resp int64

	userID := req.UserId

	followerCount, err := db.QueryFollowerCount(s.ctx, userID)
	if err != nil {
		return resp, err
	}

	resp = followerCount

	return resp, nil
}
