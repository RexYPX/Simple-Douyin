package service

import (
	"Simple-Douyin/cmd/relation/dal/db"
	"Simple-Douyin/kitex_gen/relation"
	"context"
)

type RelationFollowCountService struct {
	ctx context.Context
}

// NewRelationFollowCountService new RelationFollowCountService
func NewRelationFollowCountService(ctx context.Context) *RelationFollowCountService {
	return &RelationFollowCountService{ctx: ctx}
}

// RelationFollowCount return number of users who followed each other
func (s *RelationFollowCountService) RelationFollowCount(req *relation.RelationFollowCountRequest) (int64, error) {
	var resp int64

	userID := req.UserId

	followCount, err := db.QueryFollowCount(s.ctx, userID)
	if err != nil {
		return resp, err
	}

	resp = followCount

	return resp, nil
}
